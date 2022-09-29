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
//   cfnRulesetProps := &cfnRulesetProps{
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			checkExpression: jsii.String("checkExpression"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnSelectors: []interface{}{
//   				&columnSelectorProperty{
//   					name: jsii.String("name"),
//   					regex: jsii.String("regex"),
//   				},
//   			},
//   			disabled: jsii.Boolean(false),
//   			substitutionMap: []interface{}{
//   				&substitutionValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   			threshold: &thresholdProperty{
//   				value: jsii.Number(123),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   				unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

