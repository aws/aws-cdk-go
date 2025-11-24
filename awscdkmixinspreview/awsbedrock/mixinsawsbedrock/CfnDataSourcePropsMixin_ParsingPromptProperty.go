package mixinsawsbedrock


// Instructions for interpreting the contents of a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parsingPromptProperty := &ParsingPromptProperty{
//   	ParsingPromptText: jsii.String("parsingPromptText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingprompt.html
//
type CfnDataSourcePropsMixin_ParsingPromptProperty struct {
	// Instructions for interpreting the contents of a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingprompt.html#cfn-bedrock-datasource-parsingprompt-parsingprompttext
	//
	ParsingPromptText *string `field:"optional" json:"parsingPromptText" yaml:"parsingPromptText"`
}

