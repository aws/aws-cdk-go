package previewawslexmixins


// Specifies a custom vocabulary.
//
// A custom vocabulary is a list of words that you expect to be used during a conversation with your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customVocabularyProperty := &CustomVocabularyProperty{
//   	CustomVocabularyItems: []interface{}{
//   		&CustomVocabularyItemProperty{
//   			DisplayAs: jsii.String("displayAs"),
//   			Phrase: jsii.String("phrase"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabulary.html
//
type CfnBotPropsMixin_CustomVocabularyProperty struct {
	// Specifies a list of words that you expect to be used during a conversation with your bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabulary.html#cfn-lex-bot-customvocabulary-customvocabularyitems
	//
	CustomVocabularyItems interface{} `field:"optional" json:"customVocabularyItems" yaml:"customVocabularyItems"`
}

