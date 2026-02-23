package awscdkeksv2alpha


// Represents the options required to create an Amazon EKS Access Policy using the `fromAccessPolicyName()` method.
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
// Deprecated.
type AccessPolicyNameOptions struct {
	// The scope of the access policy.
	//
	// This determines the level of access granted by the policy.
	// Deprecated.
	AccessScopeType AccessScopeType `field:"required" json:"accessScopeType" yaml:"accessScopeType"`
	// An optional array of Kubernetes namespaces to which the access policy applies.
	// Default: - no specific namespaces for this scope.
	//
	// Deprecated.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

