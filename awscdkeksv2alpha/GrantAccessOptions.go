package awscdkeksv2alpha


// Options for granting access to a cluster.
//
// Example:
//   var cluster Cluster
//   var nodeRole Role
//
//
//   // Grant access with EC2 type for Auto Mode node role
//   cluster.GrantAccess(jsii.String("nodeAccess"), nodeRole.roleArn, []IAccessPolicy{
//   	eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAutoNodePolicy"), &AccessPolicyNameOptions{
//   		AccessScopeType: eks.AccessScopeType_CLUSTER,
//   	}),
//   }, &GrantAccessOptions{
//   	AccessEntryType: eks.AccessEntryType_EC2,
//   })
//
// Experimental.
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
	// Experimental.
	AccessEntryType AccessEntryType `field:"optional" json:"accessEntryType" yaml:"accessEntryType"`
}

