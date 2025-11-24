package mixinsawsbatch


// Provides information used to select Amazon Machine Images (AMIs) for instances in the compute environment.
//
// If `Ec2Configuration` isn't specified, the default is `ECS_AL2` ( [Amazon Linux 2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) ).
//
// > This object isn't applicable to jobs that are running on Fargate resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ec2ConfigurationObjectProperty := &Ec2ConfigurationObjectProperty{
//   	ImageIdOverride: jsii.String("imageIdOverride"),
//   	ImageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   	ImageType: jsii.String("imageType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html
//
type CfnComputeEnvironmentPropsMixin_Ec2ConfigurationObjectProperty struct {
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
	// The image type to match with the instance type to select an AMI.
	//
	// The supported values are different for `ECS` and `EKS` resources.
	//
	// - **ECS** - If the `imageIdOverride` parameter isn't specified, then a recent [Amazon ECS-optimized Amazon Linux 2 AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) ( `ECS_AL2` ) is used. If a new image type is specified in an update, but neither an `imageId` nor a `imageIdOverride` parameter is specified, then the latest Amazon ECS optimized AMI for that image type that's supported by AWS Batch is used.
	//
	// > AWS will end support for Amazon ECS optimized AL2-optimized and AL2-accelerated AMIs. Starting in January 2026, AWS Batch will change the default AMI for new Amazon ECS compute environments from Amazon Linux 2 to Amazon Linux 2023. We recommend migrating AWS Batch Amazon ECS compute environments to Amazon Linux 2023 to maintain optimal performance and security. For more information on upgrading from AL2 to AL2023, see [How to migrate from ECS AL2 to ECS AL2023](https://docs.aws.amazon.com/batch/latest/userguide/ecs-migration-2023.html) in the *AWS Batch User Guide* .
	//
	// - **ECS_AL2** - [Amazon Linux 2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#al2ami) : Default for all non-GPU instance families.
	// - **ECS_AL2_NVIDIA** - [Amazon Linux 2 (GPU)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#gpuami) : Default for all GPU instance families (for example `P4` and `G4` ) and can be used for all non AWS Graviton-based instance types.
	// - **ECS_AL2023** - [Amazon Linux 2023](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) : AWS Batch supports Amazon Linux 2023.
	//
	// > Amazon Linux 2023 does not support `A1` instances.
	// - **ECS_AL2023_NVIDIA** - [Amazon Linux 2023 (GPU)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html#gpuami) : For all GPU instance families and can be used for all non AWS Graviton-based instance types.
	//
	// > ECS_AL2023_NVIDIA doesn't support `p3` and `g3` instance types.
	// - **EKS** - If the `imageIdOverride` parameter isn't specified, then a recent [Amazon EKS-optimized Amazon Linux 2023 AMI](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) ( `EKS_AL2023` ) is used. If a new image type is specified in an update, but neither an `imageId` nor a `imageIdOverride` parameter is specified, then the latest Amazon EKS optimized AMI for that image type that AWS Batch supports is used.
	//
	// > Amazon Linux 2023 AMIs are the default on AWS Batch for Amazon EKS.
	// >
	// > AWS will end support for Amazon EKS AL2-optimized and AL2-accelerated AMIs, starting 11/26/25. You can continue using AWS Batch -provided Amazon EKS optimized Amazon Linux 2 AMIs on your Amazon EKS compute environments beyond the 11/26/25 end-of-support date, these compute environments will no longer receive any new software updates, security patches, or bug fixes from AWS . For more information on upgrading from AL2 to AL2023, see [How to upgrade from EKS AL2 to EKS AL2023](https://docs.aws.amazon.com/batch/latest/userguide/eks-migration-2023.html) in the *AWS Batch User Guide* .
	//
	// - **EKS_AL2** - [Amazon Linux 2](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Used for non-GPU instance families.
	// - **EKS_AL2_NVIDIA** - [Amazon Linux 2 (accelerated)](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Used for GPU instance families (for example, `P4` and `G4` ) and can be used for all non AWS Graviton-based instance types.
	// - **EKS_AL2023** - [Amazon Linux 2023](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Default for non-GPU instance families.
	//
	// > Amazon Linux 2023 does not support `A1` instances.
	// - **EKS_AL2023_NVIDIA** - [Amazon Linux 2023 (accelerated)](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) : Default for GPU instance families and can be used for all non AWS Graviton-based instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imagetype
	//
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
}

