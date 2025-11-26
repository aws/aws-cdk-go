package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecr@ECRReplicationAction event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRReplicationAction := #error#.NewECRReplicationAction()
//
// Experimental.
type RepositoryEvents_ECRReplicationAction interface {
}

// The jsii proxy struct for RepositoryEvents_ECRReplicationAction
type jsiiProxy_RepositoryEvents_ECRReplicationAction struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_ECRReplicationAction() RepositoryEvents_ECRReplicationAction {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_ECRReplicationAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReplicationAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_ECRReplicationAction_Override(r RepositoryEvents_ECRReplicationAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReplicationAction",
		nil, // no parameters
		r,
	)
}

