package previewawsopensearchserverlessmixins


// Properties for CfnAccessPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessPolicyMixinProps := &CfnAccessPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-accesspolicy.html
//
type CfnAccessPolicyMixinProps struct {
	// The description of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-accesspolicy.html#cfn-opensearchserverless-accesspolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-accesspolicy.html#cfn-opensearchserverless-accesspolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-accesspolicy.html#cfn-opensearchserverless-accesspolicy-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The type of access policy.
	//
	// Currently the only option is `data` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-accesspolicy.html#cfn-opensearchserverless-accesspolicy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

