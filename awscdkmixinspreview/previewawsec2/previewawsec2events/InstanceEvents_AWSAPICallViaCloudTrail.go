package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ec2@AWSAPICallViaCloudTrail event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for InstanceEvents_AWSAPICallViaCloudTrail
type jsiiProxy_InstanceEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_AWSAPICallViaCloudTrail() InstanceEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_AWSAPICallViaCloudTrail_Override(i InstanceEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		i,
	)
}

