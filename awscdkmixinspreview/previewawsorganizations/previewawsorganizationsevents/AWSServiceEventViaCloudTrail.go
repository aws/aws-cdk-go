package previewawsorganizationsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.organizations@AWSServiceEventViaCloudTrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSServiceEventViaCloudTrail := awscdkmixinspreview.Events.NewAWSServiceEventViaCloudTrail()
//
// Experimental.
type AWSServiceEventViaCloudTrail interface {
}

// The jsii proxy struct for AWSServiceEventViaCloudTrail
type jsiiProxy_AWSServiceEventViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewAWSServiceEventViaCloudTrail() AWSServiceEventViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_AWSServiceEventViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAWSServiceEventViaCloudTrail_Override(a AWSServiceEventViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AWS Service Event via CloudTrail.
// Experimental.
func AWSServiceEventViaCloudTrail_EventPattern(options *AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAWSServiceEventViaCloudTrail_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

