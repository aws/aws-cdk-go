package previewawsmanagedblockchainevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.managedblockchain@ManagedBlockchainProposalStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedBlockchainProposalStatusChange := awscdkmixinspreview.Events.NewManagedBlockchainProposalStatusChange()
//
// Experimental.
type ManagedBlockchainProposalStatusChange interface {
}

// The jsii proxy struct for ManagedBlockchainProposalStatusChange
type jsiiProxy_ManagedBlockchainProposalStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewManagedBlockchainProposalStatusChange() ManagedBlockchainProposalStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ManagedBlockchainProposalStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainProposalStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewManagedBlockchainProposalStatusChange_Override(m ManagedBlockchainProposalStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainProposalStatusChange",
		nil, // no parameters
		m,
	)
}

// EventBridge event pattern for Managed Blockchain Proposal Status Change.
// Experimental.
func ManagedBlockchainProposalStatusChange_EventPattern(options *ManagedBlockchainProposalStatusChange_ManagedBlockchainProposalStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateManagedBlockchainProposalStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_managedblockchain.events.ManagedBlockchainProposalStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

