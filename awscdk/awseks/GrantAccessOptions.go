package awseks


// Options for granting access to a cluster.
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
type GrantAccessOptions struct {
	// The type of the access entry.
	//
	// Specify `AccessEntryType.EC2` for EKS Auto Mode node roles,
	// `AccessEntryType.HYBRID_LINUX` for EKS Hybrid Nodes, or
	// `AccessEntryType.HYPERPOD_LINUX` for SageMaker HyperPod.
	//
	// Note that EC2, HYBRID_LINUX, and HYPERPOD_LINUX types cannot
	// have access policies attached per AWS EKS API constraints.
	// Default: AccessEntryType.STANDARD - Standard access entry type that supports access policies
	//
	AccessEntryType AccessEntryType `field:"optional" json:"accessEntryType" yaml:"accessEntryType"`
}

