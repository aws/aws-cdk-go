package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRuleset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRulesetProps := &CfnRulesetProps{
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			CheckExpression: jsii.String("checkExpression"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ColumnSelectors: []interface{}{
//   				&ColumnSelectorProperty{
//   					Name: jsii.String("name"),
//   					Regex: jsii.String("regex"),
//   				},
//   			},
//   			Disabled: jsii.Boolean(false),
//   			SubstitutionMap: []interface{}{
//   				&SubstitutionValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   			Threshold: &ThresholdProperty{
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   				Unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRulesetProps struct {
	// The name of the ruleset.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains metadata about the ruleset.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is associated with.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The description of the ruleset.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

