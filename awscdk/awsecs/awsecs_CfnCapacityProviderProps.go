package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCapacityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityProviderProps := &cfnCapacityProviderProps{
//   	autoScalingGroupProvider: &autoScalingGroupProviderProperty{
//   		autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   		// the properties below are optional
//   		managedScaling: &managedScalingProperty{
//   			instanceWarmupPeriod: jsii.Number(123),
//   			maximumScalingStepSize: jsii.Number(123),
//   			minimumScalingStepSize: jsii.Number(123),
//   			status: jsii.String("status"),
//   			targetCapacity: jsii.Number(123),
//   		},
//   		managedTerminationProtection: jsii.String("managedTerminationProtection"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCapacityProviderProps struct {
	// The Auto Scaling group settings for the capacity provider.
	AutoScalingGroupProvider interface{} `field:"required" json:"autoScalingGroupProvider" yaml:"autoScalingGroupProvider"`
	// The name of the capacity provider.
	//
	// If a name is specified, it cannot start with `aws` , `ecs` , or `fargate` . If no name is specified, a default name in the `CFNStackName-CFNResourceName-RandomString` format is used.
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
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

