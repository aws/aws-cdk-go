package awseks


// Represents the scope of an access policy.
//
// The scope defines the namespaces or cluster-level access granted by the policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessScope := &AccessScope{
//   	Type: awscdk.Aws_eks.AccessScopeType_NAMESPACE,
//
//   	// the properties below are optional
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
type AccessScope struct {
	// The scope type of the policy, either 'namespace' or 'cluster'.
	Type AccessScopeType `field:"required" json:"type" yaml:"type"`
	// A Kubernetes namespace that an access policy is scoped to.
	//
	// A value is required if you specified
	// namespace for Type.
	// Default: - no specific namespaces for this scope.
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

