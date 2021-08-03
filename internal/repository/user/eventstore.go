package user

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(UserV1AddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(UserV1RegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1SignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(UserV1PasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1ProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(UserV1AddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(UserV1MFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(UserLockedType, UserLockedEventMapper).
		RegisterFilterEventMapper(UserUnlockedType, UserUnlockedEventMapper).
		RegisterFilterEventMapper(UserDeactivatedType, UserDeactivatedEventMapper).
		RegisterFilterEventMapper(UserReactivatedType, UserReactivatedEventMapper).
		RegisterFilterEventMapper(UserRemovedType, UserRemovedEventMapper).
		RegisterFilterEventMapper(UserTokenAddedType, UserTokenAddedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedType, DomainClaimedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedSentType, DomainClaimedSentEventMapper).
		RegisterFilterEventMapper(UserUserNameChangedType, UsernameChangedEventMapper).
		RegisterFilterEventMapper(HumanAddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(HumanRegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanSignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(HumanPasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPAddedType, HumanExternalIDPAddedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPRemovedType, HumanExternalIDPRemovedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPCascadeRemovedType, HumanExternalIDPCascadeRemovedEventMapper).
		RegisterFilterEventMapper(HumanExternalLoginCheckSucceededType, HumanExternalIDPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanEmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(HumanPhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(HumanProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(HumanAvatarAddedType, HumanAvatarAddedEventMapper).
		RegisterFilterEventMapper(HumanAvatarRemovedType, HumanAvatarRemovedEventMapper).
		RegisterFilterEventMapper(HumanAddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(HumanMFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenAddedType, HumanU2FAddedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenVerifiedType, HumanU2FVerifiedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenSignCountChangedType, HumanU2FSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenRemovedType, HumanU2FRemovedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenBeginLoginType, HumanU2FBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckSucceededType, HumanU2FCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckFailedType, HumanU2FCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenAddedType, HumanPasswordlessAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenVerifiedType, HumanPasswordlessVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenSignCountChangedType, HumanPasswordlessSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenRemovedType, HumanPasswordlessRemovedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenBeginLoginType, HumanPasswordlessBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckSucceededType, HumanPasswordlessCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckFailedType, HumanPasswordlessCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeAddedType, HumanPasswordlessInitCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeRequestedType, HumanPasswordlessInitCodeRequestedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeSentType, HumanPasswordlessInitCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeCheckFailedType, HumanPasswordlessInitCodeCodeCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeCheckSucceededType, HumanPasswordlessInitCodeCodeCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenAddedType, HumanRefreshTokenAddedEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenRenewedType, HumanRefreshTokenRenewedEventEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenRemovedType, HumanRefreshTokenRemovedEventEventMapper).
		RegisterFilterEventMapper(MachineAddedEventType, MachineAddedEventMapper).
		RegisterFilterEventMapper(MachineChangedEventType, MachineChangedEventMapper).
		RegisterFilterEventMapper(MachineKeyAddedEventType, MachineKeyAddedEventMapper).
		RegisterFilterEventMapper(MachineKeyRemovedEventType, MachineKeyRemovedEventMapper)
}
