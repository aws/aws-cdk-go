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
//   	ComputeType: jsii.String("computeType"),
//   	EnvironmentType: jsii.String("environmentType"),
//   	Name: jsii.String("name"),
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
	// Information about the compute resources the compute fleet uses. Available values include:.
	//
	// - `BUILD_GENERAL1_SMALL` : Use up to 3 GB memory and 2 vCPUs for builds.
	// - `BUILD_GENERAL1_MEDIUM` : Use up to 7 GB memory and 4 vCPUs for builds.
	// - `BUILD_GENERAL1_LARGE` : Use up to 16 GB memory and 8 vCPUs for builds, depending on your environment type.
	// - `BUILD_GENERAL1_XLARGE` : Use up to 70 GB memory and 36 vCPUs for builds, depending on your environment type.
	// - `BUILD_GENERAL1_2XLARGE` : Use up to 145 GB memory, 72 vCPUs, and 824 GB of SSD storage for builds. This compute type supports Docker images up to 100 GB uncompressed.
	//
	// If you use `BUILD_GENERAL1_SMALL` :
	//
	// - For environment type `LINUX_CONTAINER` , you can use up to 3 GB memory and 2 vCPUs for builds.
	// - For environment type `LINUX_GPU_CONTAINER` , you can use up to 16 GB memory, 4 vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.
	// - For environment type `ARM_CONTAINER` , you can use up to 4 GB memory and 2 vCPUs on ARM-based processors for builds.
	//
	// If you use `BUILD_GENERAL1_LARGE` :
	//
	// - For environment type `LINUX_CONTAINER` , you can use up to 15 GB memory and 8 vCPUs for builds.
	// - For environment type `LINUX_GPU_CONTAINER` , you can use up to 255 GB memory, 32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.
	// - For environment type `ARM_CONTAINER` , you can use up to 16 GB memory and 8 vCPUs on ARM-based processors for builds.
	//
	// For more information, see [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild User Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The environment type of the compute fleet.
	//
	// - The environment type `ARM_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), EU (Frankfurt), and South America (São Paulo).
	// - The environment type `LINUX_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo), and Asia Pacific (Mumbai).
	// - The environment type `LINUX_GPU_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Tokyo), and Asia Pacific (Sydney).
	// - The environment type `WINDOWS_SERVER_2019_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific (Sydney), Asia Pacific (Tokyo), Asia Pacific (Mumbai) and EU (Ireland).
	// - The environment type `WINDOWS_SERVER_2022_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Sydney), Asia Pacific (Singapore), Asia Pacific (Tokyo), South America (São Paulo) and Asia Pacific (Mumbai).
	//
	// For more information, see [Build environment compute types](https://docs.aws.amazon.com//codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild user guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-environmenttype
	//
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
	// The name of the compute fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of tag key and value pairs associated with this compute fleet.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild compute fleet tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

