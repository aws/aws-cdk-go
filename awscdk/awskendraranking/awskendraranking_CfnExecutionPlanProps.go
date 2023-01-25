package awskendraranking

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnExecutionPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExecutionPlanProps := &cfnExecutionPlanProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	capacityUnits: &capacityUnitsConfigurationProperty{
//   		rescoreCapacityUnits: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnExecutionPlanProps struct {
	// `AWS::KendraRanking::ExecutionPlan.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::KendraRanking::ExecutionPlan.CapacityUnits`.
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// `AWS::KendraRanking::ExecutionPlan.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::KendraRanking::ExecutionPlan.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

