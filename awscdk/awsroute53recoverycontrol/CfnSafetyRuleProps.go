package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSafetyRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSafetyRuleProps := &CfnSafetyRuleProps{
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   	Name: jsii.String("name"),
//   	RuleConfig: &RuleConfigProperty{
//   		Inverted: jsii.Boolean(false),
//   		Threshold: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//
//   	// the properties below are optional
//   	AssertionRule: &AssertionRuleProperty{
//   		AssertedControls: []*string{
//   			jsii.String("assertedControls"),
//   		},
//   		WaitPeriodMs: jsii.Number(123),
//   	},
//   	GatingRule: &GatingRuleProperty{
//   		GatingControls: []*string{
//   			jsii.String("gatingControls"),
//   		},
//   		TargetControls: []*string{
//   			jsii.String("targetControls"),
//   		},
//   		WaitPeriodMs: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html
//
type CfnSafetyRuleProps struct {
	// The Amazon Resource Name (ARN) of the control panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-controlpanelarn
	//
	ControlPanelArn *string `field:"required" json:"controlPanelArn" yaml:"controlPanelArn"`
	// The name of the assertion rule.
	//
	// The name must be unique within a control panel. You can use any non-white space character in the name except the following: & > < ' (single quote) " (double quote) ; (semicolon)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The criteria that you set for specific assertion controls (routing controls) that designate how many control states must be `ON` as the result of a transaction.
	//
	// For example, if you have three assertion controls, you might specify `ATLEAST 2` for your rule configuration. This means that at least two assertion controls must be `ON` , so that at least two AWS Regions have traffic flowing to them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-ruleconfig
	//
	RuleConfig interface{} `field:"required" json:"ruleConfig" yaml:"ruleConfig"`
	// An assertion rule enforces that, when you change a routing control state, that the criteria that you set in the rule configuration is met.
	//
	// Otherwise, the change to the routing control is not accepted. For example, the criteria might be that at least one routing control state is `On` after the transaction so that traffic continues to flow to at least one cell for the application. This ensures that you avoid a fail-open scenario.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-assertionrule
	//
	AssertionRule interface{} `field:"optional" json:"assertionRule" yaml:"assertionRule"`
	// A gating rule verifies that a gating routing control or set of gating routing controls, evaluates as true, based on a rule configuration that you specify, which allows a set of routing control state changes to complete.
	//
	// For example, if you specify one gating routing control and you set the `Type` in the rule configuration to `OR` , that indicates that you must set the gating routing control to `On` for the rule to evaluate as true; that is, for the gating control switch to be On. When you do that, then you can update the routing control states for the target routing controls that you specify in the gating rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-gatingrule
	//
	GatingRule interface{} `field:"optional" json:"gatingRule" yaml:"gatingRule"`
	// The tags associated with the safety rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html#cfn-route53recoverycontrol-safetyrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

