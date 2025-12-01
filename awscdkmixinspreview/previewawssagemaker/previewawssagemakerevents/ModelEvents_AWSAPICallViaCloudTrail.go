package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.sagemaker@AWSAPICallViaCloudTrail event types for Model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for ModelEvents_AWSAPICallViaCloudTrail
type jsiiProxy_ModelEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewModelEvents_AWSAPICallViaCloudTrail() ModelEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_ModelEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewModelEvents_AWSAPICallViaCloudTrail_Override(m ModelEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		m,
	)
}

