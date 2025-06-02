package awslex


// Specifies a custom vocabulary.
//
// A custom vocabulary is a list of words that you expect to be used during a conversation with your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customVocabularyProperty := &CustomVocabularyProperty{
//   	CustomVocabularyItems: []interface{}{
//   		&CustomVocabularyItemProperty{
//   			Phrase: jsii.String("phrase"),
//
//   			// the properties below are optional
//   			DisplayAs: jsii.String("displayAs"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabulary.html
//
type CfnBot_CustomVocabularyProperty struct {
	// Specifies a list of words that you expect to be used during a conversation with your bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-customvocabulary.html#cfn-lex-bot-customvocabulary-customvocabularyitems
	//
	CustomVocabularyItems interface{} `field:"required" json:"customVocabularyItems" yaml:"customVocabularyItems"`
}

