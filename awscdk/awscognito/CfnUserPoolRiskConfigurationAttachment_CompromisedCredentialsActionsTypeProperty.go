package awscognito


// Settings for user pool actions when Amazon Cognito detects compromised credentials with advanced security features in full-function `ENFORCED` mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compromisedCredentialsActionsTypeProperty := &CompromisedCredentialsActionsTypeProperty{
//   	EventAction: jsii.String("eventAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsactionstype.html
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsTypeProperty struct {
	// The action that Amazon Cognito takes when it detects compromised credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsactionstype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsactionstype-eventaction
	//
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
}

