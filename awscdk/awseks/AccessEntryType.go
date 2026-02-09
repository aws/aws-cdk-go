package awseks


// Represents the different types of access entries that can be used in an Amazon EKS cluster.
//
// Example:
//   var cluster Cluster
//   var nodeRole Role
//
//
//   // For EKS Auto Mode node roles
//   cluster.GrantAccess(jsii.String("NodeAccess"), nodeRole.roleArn, []IAccessPolicy{
//   }, &GrantAccessOptions{
//   	AccessEntryType: eks.AccessEntryType_EC2,
//   })
//
type AccessEntryType string

const (
	// Represents a standard access entry.
	//
	// Use this type for standard IAM principals that need cluster access with policies.
	AccessEntryType_STANDARD AccessEntryType = "STANDARD"
	// Represents a Fargate Linux access entry.
	//
	// Use this type for AWS Fargate profiles running Linux containers.
	AccessEntryType_FARGATE_LINUX AccessEntryType = "FARGATE_LINUX"
	// Represents an EC2 Linux access entry.
	//
	// Use this type for self-managed EC2 instances running Linux that join the cluster as worker nodes.
	AccessEntryType_EC2_LINUX AccessEntryType = "EC2_LINUX"
	// Represents an EC2 Windows access entry.
	//
	// Use this type for self-managed EC2 instances running Windows that join the cluster as worker nodes.
	AccessEntryType_EC2_WINDOWS AccessEntryType = "EC2_WINDOWS"
	// Represents an EC2 access entry for EKS Auto Mode.
	//
	// Use this type for node roles in EKS Auto Mode clusters where AWS automatically manages
	// the compute infrastructure. This type cannot have access policies attached.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/eks-auto-mode.html
	//
	AccessEntryType_EC2 AccessEntryType = "EC2"
	// Represents a Hybrid Linux access entry for EKS Hybrid Nodes.
	//
	// Use this type for on-premises or edge infrastructure running Linux that connects
	// to your EKS cluster. This type cannot have access policies attached.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/hybrid-nodes.html
	//
	AccessEntryType_HYBRID_LINUX AccessEntryType = "HYBRID_LINUX"
	// Represents a HyperPod Linux access entry for Amazon SageMaker HyperPod.
	//
	// Use this type for SageMaker HyperPod clusters that need access to your EKS cluster
	// for distributed machine learning workloads. This type cannot have access policies attached.
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod.html
	//
	AccessEntryType_HYPERPOD_LINUX AccessEntryType = "HYPERPOD_LINUX"
)

