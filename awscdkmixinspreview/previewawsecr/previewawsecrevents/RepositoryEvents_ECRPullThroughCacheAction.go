package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecr@ECRPullThroughCacheAction event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRPullThroughCacheAction := #error#.NewECRPullThroughCacheAction()
//
// Experimental.
type RepositoryEvents_ECRPullThroughCacheAction interface {
}

// The jsii proxy struct for RepositoryEvents_ECRPullThroughCacheAction
type jsiiProxy_RepositoryEvents_ECRPullThroughCacheAction struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_ECRPullThroughCacheAction() RepositoryEvents_ECRPullThroughCacheAction {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_ECRPullThroughCacheAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRPullThroughCacheAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_ECRPullThroughCacheAction_Override(r RepositoryEvents_ECRPullThroughCacheAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRPullThroughCacheAction",
		nil, // no parameters
		r,
	)
}

