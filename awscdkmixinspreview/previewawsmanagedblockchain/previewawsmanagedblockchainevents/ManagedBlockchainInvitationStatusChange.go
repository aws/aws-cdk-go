package previewawsmanagedblockchainevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.managedblockchain@ManagedBlockchainInvitationStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedBlockchainInvitationStatusChange := awscdkmixinspreview.Events.NewManagedBlockchainInvitationStatusChange()
//
// Experimental.
type ManagedBlockchainInvitationStatusChange interface {
}

// The jsii proxy struct for ManagedBlockchainInvitationStatusChange
type jsiiProxy_ManagedBlockchainInvitationStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewManagedBlockchainInvitationStatusChange() ManagedBlockchainInvitationStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ManagedBlockchainInvitationStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainInvitationStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewManagedBlockchainInvitationStatusChange_Override(m ManagedBlockchainInvitationStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainInvitationStatusChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for Managed Blockchain Invitation Status Change.
// Experimental.
func ManagedBlockchainInvitationStatusChange_ManagedBlockchainInvitationStatusChangePattern(options *ManagedBlockchainInvitationStatusChange_ManagedBlockchainInvitationStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateManagedBlockchainInvitationStatusChange_ManagedBlockchainInvitationStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainInvitationStatusChange",
		"managedBlockchainInvitationStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

