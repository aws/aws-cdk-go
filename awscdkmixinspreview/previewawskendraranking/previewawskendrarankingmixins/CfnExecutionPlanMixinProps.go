package previewawskendrarankingmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnExecutionPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExecutionPlanMixinProps := &CfnExecutionPlanMixinProps{
//   	CapacityUnits: &CapacityUnitsConfigurationProperty{
//   		RescoreCapacityUnits: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendraranking-executionplan.html
//
type CfnExecutionPlanMixinProps struct {
	// You can set additional capacity units to meet the needs of your rescore execution plan.
	//
	// You are given a single capacity unit by default. If you want to use the default capacity, you don't set additional capacity units. For more information on the default capacity and additional capacity units, see [Adjusting capacity](https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendraranking-executionplan.html#cfn-kendraranking-executionplan-capacityunits
	//
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description for the rescore execution plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendraranking-executionplan.html#cfn-kendraranking-executionplan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the rescore execution plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendraranking-executionplan.html#cfn-kendraranking-executionplan-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that identify or categorize your rescore execution plan.
	//
	// You can also use tags to help control access to the rescore execution plan. Tag keys and values can consist of Unicode letters, digits, white space. They can also consist of underscore, period, colon, equal, plus, and asperand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendraranking-executionplan.html#cfn-kendraranking-executionplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

