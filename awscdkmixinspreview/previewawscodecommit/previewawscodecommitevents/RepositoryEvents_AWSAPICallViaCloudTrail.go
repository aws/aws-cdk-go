package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.codecommit@AWSAPICallViaCloudTrail event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for RepositoryEvents_AWSAPICallViaCloudTrail
type jsiiProxy_RepositoryEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_AWSAPICallViaCloudTrail() RepositoryEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_AWSAPICallViaCloudTrail_Override(r RepositoryEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		r,
	)
}

