package mixinsawsverifiedpermissions


// Properties for CfnPolicyTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPolicyTemplateMixinProps := &CfnPolicyTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   	Statement: jsii.String("statement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html
//
type CfnPolicyTemplateMixinProps struct {
	// The description to attach to the new or updated policy template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the policy store that contains the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-policystoreid
	//
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
	// Specifies the content that you want to use for the new policy template, written in the Cedar policy language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-statement
	//
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}

