package awsbedrockagentcore


// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	Policy: jsii.String("policy"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// The resource policy to create or update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-resourcepolicy.html#cfn-bedrockagentcore-resourcepolicy-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the resource for which to create or update the resource policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-resourcepolicy.html#cfn-bedrockagentcore-resourcepolicy-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

