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
//   cfnAuthPolicyProps := &CfnAuthPolicyProps{
//   	Policy: policy,
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html
//
type CfnAuthPolicyProps struct {
	// The auth policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html#cfn-vpclattice-authpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The ID or Amazon Resource Name (ARN) of the service network or service for which the policy is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html#cfn-vpclattice-authpolicy-resourceidentifier
	//
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

