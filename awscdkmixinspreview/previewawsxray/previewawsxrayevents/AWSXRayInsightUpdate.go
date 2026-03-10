package previewawsxrayevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.xray@AWSXRayInsightUpdate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSXRayInsightUpdate := awscdkmixinspreview.Events.NewAWSXRayInsightUpdate()
//
// Experimental.
type AWSXRayInsightUpdate interface {
}

// The jsii proxy struct for AWSXRayInsightUpdate
type jsiiProxy_AWSXRayInsightUpdate struct {
	_ byte // padding
}

// Experimental.
func NewAWSXRayInsightUpdate() AWSXRayInsightUpdate {
	_init_.Initialize()

	j := jsiiProxy_AWSXRayInsightUpdate{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAWSXRayInsightUpdate_Override(a AWSXRayInsightUpdate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AWS X-Ray Insight Update.
// Experimental.
func AWSXRayInsightUpdate_AwsXRayInsightUpdatePattern(options *AWSXRayInsightUpdate_AWSXRayInsightUpdateProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAWSXRayInsightUpdate_AwsXRayInsightUpdatePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate",
		"awsXRayInsightUpdatePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

