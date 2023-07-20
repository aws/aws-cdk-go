package awscognito


// The compromised credentials risk configuration type.
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
	// The compromised credentials risk configuration actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Perform the action for these events.
	//
	// The default is to perform all events if no event filter is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-eventfilter
	//
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

