package awscdkeksv2alpha


// Represents the scope type of an access policy.
//
// The scope type determines the level of access granted by the policy.
//
// Example:
//   // AmazonEKSClusterAdminPolicy with `cluster` scope
//   eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSClusterAdminPolicy"), &AccessPolicyNameOptions{
//   	AccessScopeType: eks.AccessScopeType_CLUSTER,
//   })
//   // AmazonEKSAdminPolicy with `namespace` scope
//   eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAdminPolicy"), &AccessPolicyNameOptions{
//   	AccessScopeType: eks.AccessScopeType_NAMESPACE,
//   	Namespaces: []*string{
//   		jsii.String("foo"),
//   		jsii.String("bar"),
//   	},
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

