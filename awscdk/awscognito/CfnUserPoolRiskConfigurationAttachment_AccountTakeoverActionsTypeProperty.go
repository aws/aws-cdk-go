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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html
//
type CfnUserPoolRiskConfigurationAttachment_AccountTakeoverActionsTypeProperty struct {
	// Action to take for a high risk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-highaction
	//
	HighAction interface{} `field:"optional" json:"highAction" yaml:"highAction"`
	// Action to take for a low risk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-lowaction
	//
	LowAction interface{} `field:"optional" json:"lowAction" yaml:"lowAction"`
	// Action to take for a medium risk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-mediumaction
	//
	MediumAction interface{} `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

