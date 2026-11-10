package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRetrievalInstructionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRetrievalInstructionMixinProps := &CfnRetrievalInstructionMixinProps{
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	Description: jsii.String("description"),
//   	Instruction: &InstructionProperty{
//   		AnswerFormat: jsii.String("answerFormat"),
//   		AnswerPolicy: jsii.String("answerPolicy"),
//   		Persona: jsii.String("persona"),
//   		RetrievalGuidance: jsii.String("retrievalGuidance"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html
//
type CfnRetrievalInstructionMixinProps struct {
	// The ARN of the customer managed KMS key protecting this content, when one was supplied at create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html#cfn-bedrock-retrievalinstruction-customerencryptionkeyarn
	//
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// A description of the retrieval instruction.
	//
	// On a numbered version, the description of that version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html#cfn-bedrock-retrievalinstruction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The reviewed content this instruction serves — what AgenticRetrieveStream resolves while status is ACTIVE.
	//
	// Absent until the first create review passes. Always present on a numbered version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html#cfn-bedrock-retrievalinstruction-instruction
	//
	Instruction interface{} `field:"optional" json:"instruction" yaml:"instruction"`
	// The name of the retrieval instruction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html#cfn-bedrock-retrievalinstruction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html#cfn-bedrock-retrievalinstruction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

