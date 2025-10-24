package awscdkeks-v2alpha


// Properties for configuring an Amazon EKS Access Policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   var accessPolicyArn AccessPolicyArn
//
//   accessPolicyProps := &AccessPolicyProps{
//   	AccessScope: &AccessScope{
//   		Type: eks_v2_alpha.AccessScopeType_NAMESPACE,
//
//   		// the properties below are optional
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   	Policy: accessPolicyArn,
//   }
//
// Experimental.
type AccessPolicyProps struct {
	// The scope of the access policy, which determines the level of access granted.
	// Experimental.
	AccessScope *AccessScope `field:"required" json:"accessScope" yaml:"accessScope"`
	// The access policy itself, which defines the specific permissions.
	// Experimental.
	Policy AccessPolicyArn `field:"required" json:"policy" yaml:"policy"`
}

