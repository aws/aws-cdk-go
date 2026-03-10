package previewawsmanagedblockchainevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainInvitationStatusChange",
		reflect.TypeOf((*ManagedBlockchainInvitationStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ManagedBlockchainInvitationStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainInvitationStatusChange.ManagedBlockchainInvitationStatusChangeProps",
		reflect.TypeOf((*ManagedBlockchainInvitationStatusChange_ManagedBlockchainInvitationStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainProposalStatusChange",
		reflect.TypeOf((*ManagedBlockchainProposalStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ManagedBlockchainProposalStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainProposalStatusChange.ManagedBlockchainProposalStatusChangeProps",
		reflect.TypeOf((*ManagedBlockchainProposalStatusChange_ManagedBlockchainProposalStatusChangeProps)(nil)).Elem(),
	)
}
