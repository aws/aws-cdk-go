package previewawslogsmixins


// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyName: jsii.String("policyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// The details of the policy.
	//
	// It must be formatted in JSON, and you must use backslashes to escape characters that need to be escaped in JSON strings, such as double quote marks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The name of the resource policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-resourcepolicy.html#cfn-logs-resourcepolicy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

