package awscognito


// The automated response to a risk level for adaptive authentication in full-function, or `ENFORCED` , mode.
//
// You can assign an action to each risk level that threat protection evaluates.
//
// This data type is a request parameter of `API_SetRiskConfiguration` and a response parameter of `API_DescribeRiskConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountTakeoverActionTypeProperty := &AccountTakeoverActionTypeProperty{
//   	EventAction: jsii.String("eventAction"),
//   	Notify: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionTypeProperty struct {
	// The action to take for the attempted account takeover action for the associated risk level.
	//
	// Valid values are as follows:
	//
	// - `BLOCK` : Block the request.
	// - `MFA_IF_CONFIGURED` : Present an MFA challenge if possible. MFA is possible if the user pool has active MFA methods that the user can set up. For example, if the user pool only supports SMS message MFA but the user doesn't have a phone number attribute, MFA setup isn't possible. If MFA setup isn't possible, allow the request.
	// - `MFA_REQUIRED` : Present an MFA challenge if possible. Block the request if a user hasn't set up MFA. To sign in with required MFA, users must have an email address or phone number attribute, or a registered TOTP factor.
	// - `NO_ACTION` : Take no action. Permit sign-in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-eventaction
	//
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Determines whether Amazon Cognito sends a user a notification message when your user pools assesses a user's session at the associated risk level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-notify
	//
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

