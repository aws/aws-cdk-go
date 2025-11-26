package previewawsvpclatticemixins


// Properties for CfnAuthPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnAuthPolicyMixinProps := &CfnAuthPolicyMixinProps{
//   	Policy: policy,
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html
//
type CfnAuthPolicyMixinProps struct {
	// The auth policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html#cfn-vpclattice-authpolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The ID or ARN of the service network or service for which the policy is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html#cfn-vpclattice-authpolicy-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

