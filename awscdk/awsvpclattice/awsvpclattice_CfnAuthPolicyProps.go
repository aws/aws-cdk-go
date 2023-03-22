package awsvpclattice


// Properties for defining a `CfnAuthPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnAuthPolicyProps := &cfnAuthPolicyProps{
//   	policy: policy,
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
type CfnAuthPolicyProps struct {
	// `AWS::VpcLattice::AuthPolicy.Policy`.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// `AWS::VpcLattice::AuthPolicy.ResourceIdentifier`.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

