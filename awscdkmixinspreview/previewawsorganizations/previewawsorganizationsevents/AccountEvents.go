package previewawsorganizationsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsorganizations"
)

// EventBridge event patterns for Account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accountRef IAccountRef
//
//   accountEvents := awscdkmixinspreview.Events.AccountEvents_FromAccount(accountRef)
//
// Experimental.
type AccountEvents interface {
	// EventBridge event pattern for Account AWS Service Event via CloudTrail.
	// Experimental.
	AwsServiceEventViaCloudTrailPattern(options *AccountEvents_AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps) *awsevents.EventPattern
}

// The jsii proxy struct for AccountEvents
type jsiiProxy_AccountEvents struct {
	_ byte // padding
}

// Create AccountEvents from a Account reference.
// Experimental.
func AccountEvents_FromAccount(accountRef interfacesawsorganizations.IAccountRef) AccountEvents {
	_init_.Initialize()

	if err := validateAccountEvents_FromAccountParameters(accountRef); err != nil {
		panic(err)
	}
	var returns AccountEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents",
		"fromAccount",
		[]interface{}{accountRef},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountEvents) AwsServiceEventViaCloudTrailPattern(options *AccountEvents_AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps) *awsevents.EventPattern {
	if err := a.validateAwsServiceEventViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		a,
		"awsServiceEventViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

