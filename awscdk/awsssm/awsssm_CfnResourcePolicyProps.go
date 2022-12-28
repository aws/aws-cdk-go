package awsssm


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnResourcePolicyProps := &cfnResourcePolicyProps{
//   	policy: policy,
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnResourcePolicyProps struct {
	// A policy you want to associate with a resource.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// Amazon Resource Name (ARN) of the resource to which you want to attach a policy.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

