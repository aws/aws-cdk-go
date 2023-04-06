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
//   cfnResourcePolicyProps := &CfnResourcePolicyProps{
//   	Policy: policy,
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type CfnResourcePolicyProps struct {
	// The Amazon Resource Name (ARN) of the service network or service.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// An IAM policy.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

