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
//   ec2ConfigurationObjectProperty := &Ec2ConfigurationObjectProperty{
//   	ImageType: jsii.String("imageType"),
//
//   	// the properties below are optional
//   	ImageIdOverride: jsii.String("imageIdOverride"),
//   	ImageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html
//
type CfnComputeEnvironment_Ec2ConfigurationObjectProperty struct {
	// The image type to match with the instance type to select an AMI.
	//
	// The supported values are different for `ECS` and `EKS` resources.
	//
	// - **ECS** - If the `imageIdOverride` parameter isn't specified, then a recent [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) ( `ECS_AL2` ) is used. If a new image type is specified in an update, but neither an `imageId` nor a `imageIdOverride` parameter is specified, then the latest Amazon ECS optimized AMI for that image type that's supported by AWS Batch is used.
	//
	// - **ECS_AL2** - [Amazon Linux 2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) : Default for all non-GPU instance families.
	// - **ECS_AL2_NVIDIA** - [Amazon Linux 2 (GPU)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#gpuami) : Default for all GPU instance families (for example `P4` and `G4` ) and can be used for all non AWS Graviton-based instance types.
	// - **ECS_AL2023** - [Amazon Linux 2023](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) : AWS Batch supports Amazon Linux 2023.
	//
	// > Amazon Linux 2023 does not support `A1` instances.
	// - **ECS_AL1** - [Amazon Linux](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#alami) . Amazon Linux has reached the end-of-life of standard support. For more information, see [Amazon Linux AMI](https://docs.aws.amazon.com/amazon-linux-ami/) .
	// - **EKS** - If the `imageIdOverride` parameter isn't specified, then a recent [Amazon EKS-optimized Amazon Linux AMI](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) ( `EKS_AL2` ) is used. If a new image type is specified in an update, but neither an `imageId` nor a `imageIdOverride` parameter is specified, then the latest Amazon EKS optimized AMI for that image type that AWS Batch supports is used.
	//
	// - **EKS_AL2** - [Amazon Linux 2](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Default for all non-GPU instance families.
	// - **EKS_AL2_NVIDIA** - [Amazon Linux 2 (accelerated)](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Default for all GPU instance families (for example, `P4` and `G4` ) and can be used for all non AWS Graviton-based instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imagetype
	//
	ImageType *string `field:"required" json:"imageType" yaml:"imageType"`
	// The AMI ID used for instances launched in the compute environment that match the image type.
	//
	// This setting overrides the `imageId` set in the `computeResource` object.
	//
	// > The AMI that you choose for a compute environment must match the architecture of the instance types that you intend to use for that compute environment. For example, if your compute environment uses A1 instance types, the compute resource AMI that you choose must support ARM instances. Amazon ECS vends both x86 and ARM versions of the Amazon ECS-optimized Amazon Linux 2 AMI. For more information, see [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#ecs-optimized-ami-linux-variants.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imageidoverride
	//
	ImageIdOverride *string `field:"optional" json:"imageIdOverride" yaml:"imageIdOverride"`
	// The Kubernetes version for the compute environment.
	//
	// If you don't specify a value, the latest version that AWS Batch supports is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imagekubernetesversion
	//
	ImageKubernetesVersion *string `field:"optional" json:"imageKubernetesVersion" yaml:"imageKubernetesVersion"`
}

