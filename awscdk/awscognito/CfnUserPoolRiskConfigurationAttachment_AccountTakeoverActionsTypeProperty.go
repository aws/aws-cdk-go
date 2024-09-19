package awscognito


// A list of account-takeover actions for each level of risk that Amazon Cognito might assess with advanced security features.
//
// This data type is a request parameter of [SetRiskConfiguration](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetRiskConfiguration.html) and a response parameter of [DescribeRiskConfiguration](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeRiskConfiguration.html) .
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
	// The action that you assign to a high-risk assessment by advanced security features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-highaction
	//
	HighAction interface{} `field:"optional" json:"highAction" yaml:"highAction"`
	// The action that you assign to a low-risk assessment by advanced security features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-lowaction
	//
	LowAction interface{} `field:"optional" json:"lowAction" yaml:"lowAction"`
	// The action that you assign to a medium-risk assessment by advanced security features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype-mediumaction
	//
	MediumAction interface{} `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

