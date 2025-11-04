package awscdkeksv2alpha


// Represents the scope of an access policy.
//
// The scope defines the namespaces or cluster-level access granted by the policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   accessScope := &AccessScope{
//   	Type: eks_v2_alpha.AccessScopeType_NAMESPACE,
//
//   	// the properties below are optional
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
// Experimental.
type AccessScope struct {
	// The scope type of the policy, either 'namespace' or 'cluster'.
	// Experimental.
	Type AccessScopeType `field:"required" json:"type" yaml:"type"`
	// A Kubernetes namespace that an access policy is scoped to.
	//
	// A value is required if you specified
	// namespace for Type.
	// Default: - no specific namespaces for this scope.
	//
	// Experimental.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

