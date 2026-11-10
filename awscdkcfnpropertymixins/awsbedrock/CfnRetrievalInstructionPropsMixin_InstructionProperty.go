package awsbedrock


// The reviewed content this instruction serves — what AgenticRetrieveStream resolves while status is ACTIVE.
//
// Absent until the first create review passes. Always present on a numbered version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instructionProperty := &InstructionProperty{
//   	AnswerFormat: jsii.String("answerFormat"),
//   	AnswerPolicy: jsii.String("answerPolicy"),
//   	Persona: jsii.String("persona"),
//   	RetrievalGuidance: jsii.String("retrievalGuidance"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-retrievalinstruction-instruction.html
//
type CfnRetrievalInstructionPropsMixin_InstructionProperty struct {
	// How the answer is structured, including how sources are cited.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-retrievalinstruction-instruction.html#cfn-bedrock-retrievalinstruction-instruction-answerformat
	//
	AnswerFormat *string `field:"optional" json:"answerFormat" yaml:"answerFormat"`
	// Whether and how to answer: how strongly an answer must be supported, and what to do with questions outside scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-retrievalinstruction-instruction.html#cfn-bedrock-retrievalinstruction-instruction-answerpolicy
	//
	AnswerPolicy *string `field:"optional" json:"answerPolicy" yaml:"answerPolicy"`
	// Who the assistant is: identity, tone and language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-retrievalinstruction-instruction.html#cfn-bedrock-retrievalinstruction-instruction-persona
	//
	Persona *string `field:"optional" json:"persona" yaml:"persona"`
	// How to interpret and decompose the question, and what to prioritise when searching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-retrievalinstruction-instruction.html#cfn-bedrock-retrievalinstruction-instruction-retrievalguidance
	//
	RetrievalGuidance *string `field:"optional" json:"retrievalGuidance" yaml:"retrievalGuidance"`
}

