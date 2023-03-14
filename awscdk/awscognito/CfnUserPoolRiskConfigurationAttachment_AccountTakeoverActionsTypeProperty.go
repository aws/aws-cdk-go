package awscognito


// Account takeover actions type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountTakeoverActionsTypeProperty := &AccountTakeoverActionsTypeProperty{
//   	HighAction: &AccountTakeoverActionTypeProperty{
//   		EventAction: jsii.String("eventAction"),
//   		Notify: jsii.Boolean(false),
//   	},
//   	LowAction: &AccountTakeoverActionTypeProperty{
//   		EventAction: jsii.String("eventAction"),
//   		Notify: jsii.Boolean(false),
//   	},
//   	MediumAction: &AccountTakeoverActionTypeProperty{
//   		EventAction: jsii.String("eventAction"),
//   		Notify: jsii.Boolean(false),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionsTypeProperty struct {
	// Action to take for a high risk.
	HighAction interface{} `field:"optional" json:"highAction" yaml:"highAction"`
	// Action to take for a low risk.
	LowAction interface{} `field:"optional" json:"lowAction" yaml:"lowAction"`
	// Action to take for a medium risk.
	MediumAction interface{} `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

