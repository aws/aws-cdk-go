package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@AWSAPICallViaCloudTrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := awscdkmixinspreview.Events.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for AWSAPICallViaCloudTrail
type jsiiProxy_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewAWSAPICallViaCloudTrail() AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAWSAPICallViaCloudTrail_Override(a AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.AWSAPICallViaCloudTrail",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AWS API Call via CloudTrail.
// Experimental.
func AWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPattern(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.AWSAPICallViaCloudTrail",
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

