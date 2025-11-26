package previewawsopensearchserverlessmixins


// Properties for CfnSecurityPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityPolicyMixinProps := &CfnSecurityPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html
//
type CfnSecurityPolicyMixinProps struct {
	// The description of the security policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The type of security policy.
	//
	// Can be either `encryption` or `network` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

