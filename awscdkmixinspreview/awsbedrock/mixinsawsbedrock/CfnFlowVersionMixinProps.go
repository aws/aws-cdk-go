package mixinsawsbedrock


// Properties for CfnFlowVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowVersionMixinProps := &CfnFlowVersionMixinProps{
//   	Description: jsii.String("description"),
//   	FlowArn: jsii.String("flowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowversion.html
//
type CfnFlowVersionMixinProps struct {
	// The description of the flow version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowversion.html#cfn-bedrock-flowversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the flow that the version belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flowversion.html#cfn-bedrock-flowversion-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

