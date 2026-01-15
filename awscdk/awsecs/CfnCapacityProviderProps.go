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
//   	ClusterName: jsii.String("clusterName"),
//   	ManagedInstancesProvider: &ManagedInstancesProviderProperty{
//   		InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   		InstanceLaunchTemplate: &InstanceLaunchTemplateProperty{
//   			Ec2InstanceProfileArn: jsii.String("ec2InstanceProfileArn"),
//   			NetworkConfiguration: &ManagedInstancesNetworkConfigurationProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CapacityOptionType: jsii.String("capacityOptionType"),
//   			InstanceRequirements: &InstanceRequirementsRequestProperty{
//   				MemoryMiB: &MemoryMiBRequestProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeRequestProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				AcceleratorCount: &AcceleratorCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				AcceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				BareMetal: jsii.String("bareMetal"),
//   				BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				BurstablePerformance: jsii.String("burstablePerformance"),
//   				CpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				InstanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				LocalStorage: jsii.String("localStorage"),
//   				LocalStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   				MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				RequireHibernateSupport: jsii.Boolean(false),
//   				SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			Monitoring: jsii.String("monitoring"),
//   			StorageConfiguration: &ManagedInstancesStorageConfigurationProperty{
//   				StorageSizeGiB: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		InfrastructureOptimization: &InfrastructureOptimizationProperty{
//   			ScaleInAfter: jsii.Number(123),
//   		},
//   		PropagateTags: jsii.String("propagateTags"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
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
	AutoScalingGroupProvider interface{} `field:"optional" json:"autoScalingGroupProvider" yaml:"autoScalingGroupProvider"`
	// The cluster that this capacity provider is associated with.
	//
	// Managed instances capacity providers are cluster-scoped, meaning they can only be used within their associated cluster.
	//
	// This is required for Managed instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#cfn-ecs-capacityprovider-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The configuration for the Amazon ECS Managed Instances provider.
	//
	// This includes the infrastructure role, the launch template configuration, and tag propagation settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#cfn-ecs-capacityprovider-managedinstancesprovider
	//
	ManagedInstancesProvider interface{} `field:"optional" json:"managedInstancesProvider" yaml:"managedInstancesProvider"`
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

