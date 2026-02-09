package awscdkeksv2alpha


// Represents the scope type of an access policy.
//
// The scope type determines the level of access granted by the policy.
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
type AccessScopeType string

const (
	// The policy applies to a specific namespace within the cluster.
	// Experimental.
	AccessScopeType_NAMESPACE AccessScopeType = "NAMESPACE"
	// The policy applies to the entire cluster.
	// Experimental.
	AccessScopeType_CLUSTER AccessScopeType = "CLUSTER"
)

