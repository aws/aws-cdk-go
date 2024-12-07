package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parsingPromptProperty := &ParsingPromptProperty{
//   	ParsingPromptText: jsii.String("parsingPromptText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingprompt.html
//
type CfnKnowledgeBase_ParsingPromptProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingprompt.html#cfn-wisdom-knowledgebase-parsingprompt-parsingprompttext
	//
	ParsingPromptText *string `field:"required" json:"parsingPromptText" yaml:"parsingPromptText"`
}

