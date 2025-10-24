package awseks


// Properties for configuring an Amazon EKS Access Policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicyArn AccessPolicyArn
//
//   accessPolicyProps := &AccessPolicyProps{
//   	AccessScope: &AccessScope{
//   		Type: awscdk.Aws_eks.AccessScopeType_NAMESPACE,
//
//   		// the properties below are optional
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   	Policy: accessPolicyArn,
//   }
//
type AccessPolicyProps struct {
	// The scope of the access policy, which determines the level of access granted.
	AccessScope *AccessScope `field:"required" json:"accessScope" yaml:"accessScope"`
	// The access policy itself, which defines the specific permissions.
	Policy AccessPolicyArn `field:"required" json:"policy" yaml:"policy"`
}

