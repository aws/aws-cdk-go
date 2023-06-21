package awskendraranking

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExecutionPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExecutionPlanProps := &CfnExecutionPlanProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CapacityUnits: &CapacityUnitsConfigurationProperty{
//   		RescoreCapacityUnits: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnExecutionPlanProps struct {
	// A name for the rescore execution plan.
	Name *string `field:"required" json:"name" yaml:"name"`
	// You can set additional capacity units to meet the needs of your rescore execution plan.
	//
	// You are given a single capacity unit by default. If you want to use the default capacity, you don't set additional capacity units. For more information on the default capacity and additional capacity units, see [Adjusting capacity](https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html) .
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description for the rescore execution plan.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of key-value pairs that identify or categorize your rescore execution plan.
	//
	// You can also use tags to help control access to the rescore execution plan. Tag keys and values can consist of Unicode letters, digits, white space. They can also consist of underscore, period, colon, equal, plus, and asperand.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

