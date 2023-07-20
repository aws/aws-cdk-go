package awscognito


// Account takeover action type.
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
	// The action to take in response to the account takeover action. Valid values are as follows:.
	//
	// - `BLOCK` Choosing this action will block the request.
	// - `MFA_IF_CONFIGURED` Present an MFA challenge if user has configured it, else allow the request.
	// - `MFA_REQUIRED` Present an MFA challenge if user has configured it, else block the request.
	// - `NO_ACTION` Allow the user to sign in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-eventaction
	//
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Flag specifying whether to send a notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype-notify
	//
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

