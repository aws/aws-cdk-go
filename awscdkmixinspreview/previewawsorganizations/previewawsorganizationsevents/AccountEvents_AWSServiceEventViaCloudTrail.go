package previewawsorganizationsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.organizations@AWSServiceEventViaCloudTrail event types for Account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSServiceEventViaCloudTrail := #error#.NewAWSServiceEventViaCloudTrail()
//
// Experimental.
type AccountEvents_AWSServiceEventViaCloudTrail interface {
}

// The jsii proxy struct for AccountEvents_AWSServiceEventViaCloudTrail
type jsiiProxy_AccountEvents_AWSServiceEventViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewAccountEvents_AWSServiceEventViaCloudTrail() AccountEvents_AWSServiceEventViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_AccountEvents_AWSServiceEventViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAccountEvents_AWSServiceEventViaCloudTrail_Override(a AccountEvents_AWSServiceEventViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail",
		nil, // no parameters
		a,
	)
}

