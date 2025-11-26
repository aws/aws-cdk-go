package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecr@ECRImageAction event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRImageAction := #error#.NewECRImageAction()
//
// Experimental.
type RepositoryEvents_ECRImageAction interface {
}

// The jsii proxy struct for RepositoryEvents_ECRImageAction
type jsiiProxy_RepositoryEvents_ECRImageAction struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_ECRImageAction() RepositoryEvents_ECRImageAction {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_ECRImageAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_ECRImageAction_Override(r RepositoryEvents_ECRImageAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageAction",
		nil, // no parameters
		r,
	)
}

