package awscognito


// Settings for compromised-credentials actions and authentication-event sources with advanced security features in full-function `ENFORCED` mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compromisedCredentialsRiskConfigurationTypeProperty := &CompromisedCredentialsRiskConfigurationTypeProperty{
//   	Actions: &CompromisedCredentialsActionsTypeProperty{
//   		EventAction: jsii.String("eventAction"),
//   	},
//
//   	// the properties below are optional
//   	EventFilter: []*string{
//   		jsii.String("eventFilter"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationTypeProperty struct {
	// Settings for the actions that you want your user pool to take when Amazon Cognito detects compromised credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Settings for the sign-in activity where you want to configure compromised-credentials actions.
	//
	// Defaults to all events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-eventfilter
	//
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

