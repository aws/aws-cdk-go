package previewawswisdommixins


// Configuration for an external Bedrock knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   externalBedrockKnowledgeBaseConfigProperty := &ExternalBedrockKnowledgeBaseConfigProperty{
//   	AccessRoleArn: jsii.String("accessRoleArn"),
//   	BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html
//
type CfnAssistantAssociationPropsMixin_ExternalBedrockKnowledgeBaseConfigProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role used to access the external Bedrock knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html#cfn-wisdom-assistantassociation-externalbedrockknowledgebaseconfig-accessrolearn
	//
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// The Amazon Resource Name (ARN) of the external Bedrock knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html#cfn-wisdom-assistantassociation-externalbedrockknowledgebaseconfig-bedrockknowledgebasearn
	//
	BedrockKnowledgeBaseArn *string `field:"optional" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
}

