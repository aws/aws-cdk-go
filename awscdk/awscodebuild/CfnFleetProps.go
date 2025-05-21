package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &CfnFleetProps{
//   	BaseCapacity: jsii.Number(123),
//   	ComputeConfiguration: &ComputeConfigurationProperty{
//   		Disk: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		MachineType: jsii.String("machineType"),
//   		Memory: jsii.Number(123),
//   		VCpu: jsii.Number(123),
//   	},
//   	ComputeType: jsii.String("computeType"),
//   	EnvironmentType: jsii.String("environmentType"),
//   	FleetProxyConfiguration: &ProxyConfigurationProperty{
//   		DefaultBehavior: jsii.String("defaultBehavior"),
//   		OrderedProxyRules: []interface{}{
//   			&FleetProxyRuleProperty{
//   				Effect: jsii.String("effect"),
//   				Entities: []*string{
//   					jsii.String("entities"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	FleetServiceRole: jsii.String("fleetServiceRole"),
//   	FleetVpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	ImageId: jsii.String("imageId"),
//   	Name: jsii.String("name"),
//   	OverflowBehavior: jsii.String("overflowBehavior"),
//   	ScalingConfiguration: &ScalingConfigurationInputProperty{
//   		MaxCapacity: jsii.Number(123),
//   		ScalingType: jsii.String("scalingType"),
//   		TargetTrackingScalingConfigs: []interface{}{
//   			&TargetTrackingScalingConfigurationProperty{
//   				MetricType: jsii.String("metricType"),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html
//
type CfnFleetProps struct {
	// The initial number of machines allocated to the compute ﬂeet, which deﬁnes the number of builds that can run in parallel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-basecapacity
	//
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// The compute configuration of the compute fleet.
	//
	// This is only required if `computeType` is set to `ATTRIBUTE_BASED_COMPUTE` or `CUSTOM_INSTANCE_TYPE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-computeconfiguration
	//
	ComputeConfiguration interface{} `field:"optional" json:"computeConfiguration" yaml:"computeConfiguration"`
	// Information about the compute resources the compute fleet uses. Available values include:.
	//
	// - `ATTRIBUTE_BASED_COMPUTE` : Specify the amount of vCPUs, memory, disk space, and the type of machine.
	//
	// > If you use `ATTRIBUTE_BASED_COMPUTE` , you must define your attributes by using `computeConfiguration` . AWS CodeBuild will select the cheapest instance that satisfies your specified attributes. For more information, see [Reserved capacity environment types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.types) in the *AWS CodeBuild User Guide* .
	// - `BUILD_GENERAL1_SMALL` : Use up to 4 GiB memory and 2 vCPUs for builds.
	// - `BUILD_GENERAL1_MEDIUM` : Use up to 8 GiB memory and 4 vCPUs for builds.
	// - `BUILD_GENERAL1_LARGE` : Use up to 16 GiB memory and 8 vCPUs for builds, depending on your environment type.
	// - `BUILD_GENERAL1_XLARGE` : Use up to 72 GiB memory and 36 vCPUs for builds, depending on your environment type.
	// - `BUILD_GENERAL1_2XLARGE` : Use up to 144 GiB memory, 72 vCPUs, and 824 GB of SSD storage for builds. This compute type supports Docker images up to 100 GB uncompressed.
	// - `BUILD_LAMBDA_1GB` : Use up to 1 GiB memory for builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER` .
	// - `BUILD_LAMBDA_2GB` : Use up to 2 GiB memory for builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER` .
	// - `BUILD_LAMBDA_4GB` : Use up to 4 GiB memory for builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER` .
	// - `BUILD_LAMBDA_8GB` : Use up to 8 GiB memory for builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER` .
	// - `BUILD_LAMBDA_10GB` : Use up to 10 GiB memory for builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER` .
	//
	// If you use `BUILD_GENERAL1_SMALL` :
	//
	// - For environment type `LINUX_CONTAINER` , you can use up to 4 GiB memory and 2 vCPUs for builds.
	// - For environment type `LINUX_GPU_CONTAINER` , you can use up to 16 GiB memory, 4 vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.
	// - For environment type `ARM_CONTAINER` , you can use up to 4 GiB memory and 2 vCPUs on ARM-based processors for builds.
	//
	// If you use `BUILD_GENERAL1_LARGE` :
	//
	// - For environment type `LINUX_CONTAINER` , you can use up to 16 GiB memory and 8 vCPUs for builds.
	// - For environment type `LINUX_GPU_CONTAINER` , you can use up to 255 GiB memory, 32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.
	// - For environment type `ARM_CONTAINER` , you can use up to 16 GiB memory and 8 vCPUs on ARM-based processors for builds.
	//
	// For more information, see [On-demand environment types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types) in the *AWS CodeBuild User Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The environment type of the compute fleet.
	//
	// - The environment type `ARM_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), EU (Frankfurt), and South America (São Paulo).
	// - The environment type `ARM_EC2` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo), and Asia Pacific (Mumbai).
	// - The environment type `LINUX_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo), and Asia Pacific (Mumbai).
	// - The environment type `LINUX_EC2` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo), and Asia Pacific (Mumbai).
	// - The environment type `LINUX_GPU_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), and Asia Pacific (Sydney).
	// - The environment type `MAC_ARM` is available only in regions US East (Ohio), US East (N. Virginia), US West (Oregon), Europe (Frankfurt), and Asia Pacific (Sydney).
	// - The environment type `WINDOWS_EC2` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo), and Asia Pacific (Mumbai).
	// - The environment type `WINDOWS_SERVER_2019_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific (Sydney), Asia Pacific (Tokyo), Asia Pacific (Mumbai) and EU (Ireland).
	// - The environment type `WINDOWS_SERVER_2022_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Sydney), Asia Pacific (Singapore), Asia Pacific (Tokyo), South America (São Paulo) and Asia Pacific (Mumbai).
	//
	// For more information, see [Build environment compute types](https://docs.aws.amazon.com//codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild user guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-environmenttype
	//
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
	// Information about the proxy configurations that apply network access control to your reserved capacity instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetproxyconfiguration
	//
	FleetProxyConfiguration interface{} `field:"optional" json:"fleetProxyConfiguration" yaml:"fleetProxyConfiguration"`
	// The service role associated with the compute fleet.
	//
	// For more information, see [Allow a user to add a permission policy for a fleet service role](https://docs.aws.amazon.com/codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.html#customer-managed-policies-example-permission-policy-fleet-service-role.html) in the *AWS CodeBuild User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetservicerole
	//
	FleetServiceRole *string `field:"optional" json:"fleetServiceRole" yaml:"fleetServiceRole"`
	// Information about the VPC configuration that AWS CodeBuild accesses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetvpcconfig
	//
	FleetVpcConfig interface{} `field:"optional" json:"fleetVpcConfig" yaml:"fleetVpcConfig"`
	// The Amazon Machine Image (AMI) of the compute fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-imageid
	//
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The name of the compute fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The compute fleet overflow behavior.
	//
	// - For overflow behavior `QUEUE` , your overflow builds need to wait on the existing fleet instance to become available.
	// - For overflow behavior `ON_DEMAND` , your overflow builds run on CodeBuild on-demand.
	//
	// > If you choose to set your overflow behavior to on-demand while creating a VPC-connected fleet, make sure that you add the required VPC permissions to your project service role. For more information, see [Example policy statement to allow CodeBuild access to AWS services required to create a VPC network interface](https://docs.aws.amazon.com/codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.html#customer-managed-policies-example-create-vpc-network-interface) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-overflowbehavior
	//
	OverflowBehavior *string `field:"optional" json:"overflowBehavior" yaml:"overflowBehavior"`
	// The scaling configuration of the compute fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-scalingconfiguration
	//
	ScalingConfiguration interface{} `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// A list of tag key and value pairs associated with this compute fleet.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild compute fleet tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

