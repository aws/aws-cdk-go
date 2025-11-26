package previewawsopensearchserverlessmixins


// Properties for CfnLifecyclePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLifecyclePolicyMixinProps := &CfnLifecyclePolicyMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-lifecyclepolicy.html
//
type CfnLifecyclePolicyMixinProps struct {
	// The description of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-lifecyclepolicy.html#cfn-opensearchserverless-lifecyclepolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-lifecyclepolicy.html#cfn-opensearchserverless-lifecyclepolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-lifecyclepolicy.html#cfn-opensearchserverless-lifecyclepolicy-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The type of lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-lifecyclepolicy.html#cfn-opensearchserverless-lifecyclepolicy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

