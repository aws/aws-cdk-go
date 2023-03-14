package awscognito


// Account takeover action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountTakeoverActionTypeProperty := &accountTakeoverActionTypeProperty{
//   	eventAction: jsii.String("eventAction"),
//   	notify: jsii.Boolean(false),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionTypeProperty struct {
	// The action to take in response to the account takeover action. Valid values are as follows:.
	//
	// - `BLOCK` Choosing this action will block the request.
	// - `MFA_IF_CONFIGURED` Present an MFA challenge if user has configured it, else allow the request.
	// - `MFA_REQUIRED` Present an MFA challenge if user has configured it, else block the request.
	// - `NO_ACTION` Allow the user to sign in.
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Flag specifying whether to send a notification.
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

