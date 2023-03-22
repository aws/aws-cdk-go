package awsvpclattice


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
	// `AWS::VpcLattice::ResourcePolicy.Policy`.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// `AWS::VpcLattice::ResourcePolicy.ResourceArn`.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

