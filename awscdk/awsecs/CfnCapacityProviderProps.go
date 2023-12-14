package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapacityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProviderProps := &CfnCapacityProviderProps{
//   	AutoScalingGroupProvider: &AutoScalingGroupProviderProperty{
//   		AutoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   		// the properties below are optional
//   		ManagedDraining: jsii.String("managedDraining"),
//   		ManagedScaling: &ManagedScalingProperty{
//   			InstanceWarmupPeriod: jsii.Number(123),
//   			MaximumScalingStepSize: jsii.Number(123),
//   			MinimumScalingStepSize: jsii.Number(123),
//   			Status: jsii.String("status"),
//   			TargetCapacity: jsii.Number(123),
//   		},
//   		ManagedTerminationProtection: jsii.String("managedTerminationProtection"),
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html
//
type CfnCapacityProviderProps struct {
	// The Auto Scaling group settings for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider
	//
	AutoScalingGroupProvider interface{} `field:"required" json:"autoScalingGroupProvider" yaml:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#cfn-ecs-capacityprovider-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The metadata that you apply to the capacity provider to help you categorize and organize it.
	//
	// Each tag consists of a key and an optional value. You define both.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#cfn-ecs-capacityprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

