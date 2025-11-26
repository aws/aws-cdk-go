package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecr@ECRReferrerAction event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRReferrerAction := #error#.NewECRReferrerAction()
//
// Experimental.
type RepositoryEvents_ECRReferrerAction interface {
}

// The jsii proxy struct for RepositoryEvents_ECRReferrerAction
type jsiiProxy_RepositoryEvents_ECRReferrerAction struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_ECRReferrerAction() RepositoryEvents_ECRReferrerAction {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_ECRReferrerAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReferrerAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_ECRReferrerAction_Override(r RepositoryEvents_ECRReferrerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReferrerAction",
		nil, // no parameters
		r,
	)
}

