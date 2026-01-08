package previewawscasesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCaseRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var emptyValue interface{}
//
//   cfnCaseRuleMixinProps := &CfnCaseRuleMixinProps{
//   	Description: jsii.String("description"),
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Rule: &CaseRuleDetailsProperty{
//   		Hidden: &HiddenCaseRuleProperty{
//   			Conditions: []interface{}{
//   				&BooleanConditionProperty{
//   					EqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   					NotEqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultValue: jsii.Boolean(false),
//   		},
//   		Required: &RequiredCaseRuleProperty{
//   			Conditions: []interface{}{
//   				&BooleanConditionProperty{
//   					EqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   					NotEqualTo: &BooleanOperandsProperty{
//   						OperandOne: &OperandOneProperty{
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						OperandTwo: &OperandTwoProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   						Result: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultValue: jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html
//
type CfnCaseRuleMixinProps struct {
	// Description of a case rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html#cfn-cases-caserule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Unique identifier of a Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html#cfn-cases-caserule-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Name of the case rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html#cfn-cases-caserule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents what rule type should take place, under what conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html#cfn-cases-caserule-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-caserule.html#cfn-cases-caserule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

