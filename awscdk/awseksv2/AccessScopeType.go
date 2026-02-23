package awseksv2


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
type AccessScopeType string

const (
	// The policy applies to a specific namespace within the cluster.
	AccessScopeType_NAMESPACE AccessScopeType = "NAMESPACE"
	// The policy applies to the entire cluster.
	AccessScopeType_CLUSTER AccessScopeType = "CLUSTER"
)

