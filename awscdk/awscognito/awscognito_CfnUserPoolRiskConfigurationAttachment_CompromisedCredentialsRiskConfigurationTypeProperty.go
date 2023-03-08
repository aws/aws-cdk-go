package awscognito


// The compromised credentials risk configuration type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compromisedCredentialsRiskConfigurationTypeProperty := &compromisedCredentialsRiskConfigurationTypeProperty{
//   	actions: &compromisedCredentialsActionsTypeProperty{
//   		eventAction: jsii.String("eventAction"),
//   	},
//
//   	// the properties below are optional
//   	eventFilter: []*string{
//   		jsii.String("eventFilter"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationTypeProperty struct {
	// The compromised credentials risk configuration actions.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Perform the action for these events.
	//
	// The default is to perform all events if no event filter is specified.
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

