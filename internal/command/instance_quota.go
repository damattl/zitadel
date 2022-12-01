package command

import (
	"context"
	"time"

	"github.com/zitadel/zitadel/internal/api/authz"

	"github.com/zitadel/zitadel/internal/repository/quota"

	"github.com/zitadel/zitadel/internal/command/preparation"
	"github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/instance"

	"github.com/zitadel/zitadel/internal/domain"
)

type QuotaUnit string

const (
	QuotaRequestsAllAuthenticated QuotaUnit = "requests.all.authenticated"
	QuotaActionsAllRunsSeconds              = "actions.all.runs.seconds"
)

func (q *QuotaUnit) Enum() quota.Unit {
	switch *q {
	case QuotaRequestsAllAuthenticated:
		return quota.RequestsAllAuthenticated
	case QuotaActionsAllRunsSeconds:
		return quota.ActionsAllRunsSeconds
	default:
		return quota.Unimplemented
	}
}

func (c *Commands) AddInstanceQuota(
	ctx context.Context,
	quota *Quota,
) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.AddInstanceQuotaCommand(instanceAgg, quota))
	if err != nil {
		return nil, err
	}
	events, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	wm := &eventstore.WriteModel{
		AggregateID:   authz.GetInstance(ctx).InstanceID(),
		ResourceOwner: authz.GetInstance(ctx).InstanceID(),
	}
	err = AppendAndReduce(wm, events...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(wm), nil
}

func (c *Commands) RemoveInstanceQuota(ctx context.Context, unit quota.Unit) (*domain.ObjectDetails, error) {
	// TODO: Implement
	return nil, errors.ThrowUnimplemented(nil, "INSTA-h12vl", "*Commands.RemoveInstanceQuota is unimplemented")
}

type QuotaNotification struct {
	Percent uint32
	Repeat  bool
	CallURL string
}

type QuotaNotifications []*QuotaNotification

func (q *QuotaNotifications) toAddedEventNotifications(genID func() string) []*quota.AddedEventNotification {
	if q == nil {
		return nil
	}

	notifications := make([]*quota.AddedEventNotification, len(*q))
	for idx, notification := range *q {

		notifications[idx] = &quota.AddedEventNotification{
			ID:      genID(),
			Percent: notification.Percent,
			Repeat:  notification.Repeat,
			CallURL: notification.CallURL,
		}
	}

	return notifications
}

type Quota struct {
	Unit          QuotaUnit
	From          string
	Interval      time.Duration
	Amount        uint64
	Limit         bool
	Notifications QuotaNotifications
}

func (c *Commands) AddInstanceQuotaCommand(
	a *instance.Aggregate,
	q *Quota,
) preparation.Validation {
	return func() (preparation.CreateCommands, error) {

		unit := q.Unit.Enum()
		if unit == quota.Unimplemented {
			return nil, errors.ThrowInvalidArgument(nil, "INSTA-SDSfs", "Errors.Invalid.Argument") // TODO: Better error message?
		}

		from, err := time.Parse("2006-01-02 15:04:05", q.From)
		if err != nil {
			return nil, errors.ThrowInvalidArgument(err, "INSTA-H2Poe", "Errors.Invalid.Argument") // TODO: Better error message?
		}

		// TODO: More validations without side effects

		return func(ctx context.Context, filter preparation.FilterToQueryReducer) (cmd []eventstore.Command, err error) {
				// TODO: Validations with side effects
				genID := func() string {
					id, genErr := c.idGenerator.Next()
					if genErr != nil {
						err = genErr
					}
					return id
				}

				return []eventstore.Command{instance.NewQuotaAddedEvent(
					ctx,
					&a.Aggregate,
					unit,
					from,
					q.Interval,
					q.Amount,
					q.Limit,
					q.Notifications.toAddedEventNotifications(genID),
				)}, err
			},
			nil
	}
}