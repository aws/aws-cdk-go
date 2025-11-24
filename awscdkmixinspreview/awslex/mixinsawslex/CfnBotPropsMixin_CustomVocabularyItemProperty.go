package mixinsawslex


// Specifies an entry in a custom vocabulary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customVocabularyItemProperty := &CustomVocabularyItemProperty{
//   	DisplayAs: jsii.String("displayAs"),
//   	Phrase: jsii.String("phrase"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabularyitem.html
//
type CfnBotPropsMixin_CustomVocabularyItemProperty struct {
	// The DisplayAs value for the custom vocabulary item from the custom vocabulary list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabularyitem.html#cfn-lex-bot-customvocabularyitem-displayas
	//
	DisplayAs *string `field:"optional" json:"displayAs" yaml:"displayAs"`
	// Specifies 1 - 4 words that should be recognized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabularyitem.html#cfn-lex-bot-customvocabularyitem-phrase
	//
	Phrase *string `field:"optional" json:"phrase" yaml:"phrase"`
	// Specifies the degree to which the phrase recognition is boosted.
	//
	// The default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabularyitem.html#cfn-lex-bot-customvocabularyitem-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

