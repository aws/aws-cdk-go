package awsbatch


// Provides information used to select Amazon Machine Images (AMIs) for instances in the compute environment.
//
// If `Ec2Configuration` isn't specified, the default is `ECS_AL2` ( [Amazon Linux 2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) ).
//
// > This object isn't applicable to jobs that are running on Fargate resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2ConfigurationObjectProperty := &ec2ConfigurationObjectProperty{
//   	imageType: jsii.String("imageType"),
//
//   	// the properties below are optional
//   	imageIdOverride: jsii.String("imageIdOverride"),
//   	imageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   }
//
type CfnComputeEnvironment_Ec2ConfigurationObjectProperty struct {
	// The image type to match with the instance type to select an AMI.
	//
	// If the `imageIdOverride` parameter isn't specified, then a recent [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) ( `ECS_AL2` ) is used. If a new image type is specified in an update, but neither an `imageId` nor a `imageIdOverride` parameter is specified, then the latest Amazon ECS optimized AMI for that image type that's supported by AWS Batch is used.
	//
	// - **ECS_AL2** - [Amazon Linux 2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) − Default for all non-GPU instance families.
	// - **ECS_AL2_NVIDIA** - [Amazon Linux 2 (GPU)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#gpuami) −Default for all GPU instance families (for example `P4` and `G4` ) and can be used for all non AWS Graviton-based instance types.
	// - **ECS_AL1** - [Amazon Linux](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#alami) . Amazon Linux is reaching the end-of-life of standard support. For more information, see [Amazon Linux AMI](https://docs.aws.amazon.com/amazon-linux-ami/) .
	ImageType *string `field:"required" json:"imageType" yaml:"imageType"`
	// The AMI ID used for instances launched in the compute environment that match the image type.
	//
	// This setting overrides the `imageId` set in the `computeResource` object.
	//
	// > The AMI that you choose for a compute environment must match the architecture of the instance types that you intend to use for that compute environment. For example, if your compute environment uses A1 instance types, the compute resource AMI that you choose must support ARM instances. Amazon ECS vends both x86 and ARM versions of the Amazon ECS-optimized Amazon Linux 2 AMI. For more information, see [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#ecs-optimized-ami-linux-variants.html) in the *Amazon Elastic Container Service Developer Guide* .
	ImageIdOverride *string `field:"optional" json:"imageIdOverride" yaml:"imageIdOverride"`
	// `CfnComputeEnvironment.Ec2ConfigurationObjectProperty.ImageKubernetesVersion`.
	ImageKubernetesVersion *string `field:"optional" json:"imageKubernetesVersion" yaml:"imageKubernetesVersion"`
}

