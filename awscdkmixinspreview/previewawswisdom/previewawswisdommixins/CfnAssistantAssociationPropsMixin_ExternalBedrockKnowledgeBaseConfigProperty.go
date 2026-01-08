package previewawswisdommixins


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html#cfn-wisdom-assistantassociation-externalbedrockknowledgebaseconfig-accessrolearn
	//
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html#cfn-wisdom-assistantassociation-externalbedrockknowledgebaseconfig-bedrockknowledgebasearn
	//
	BedrockKnowledgeBaseArn *string `field:"optional" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
}

