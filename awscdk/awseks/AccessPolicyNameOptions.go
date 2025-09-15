package awseks


// Represents the options required to create an Amazon EKS Access Policy using the `fromAccessPolicyName()` method.
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
type AccessPolicyNameOptions struct {
	// The scope of the access policy.
	//
	// This determines the level of access granted by the policy.
	AccessScopeType AccessScopeType `field:"required" json:"accessScopeType" yaml:"accessScopeType"`
	// An optional array of Kubernetes namespaces to which the access policy applies.
	// Default: - no specific namespaces for this scope.
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

