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
//   customVocabularyProperty := &customVocabularyProperty{
//   	customVocabularyItems: []interface{}{
//   		&customVocabularyItemProperty{
//   			phrase: jsii.String("phrase"),
//
//   			// the properties below are optional
//   			weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBot_CustomVocabularyProperty struct {
	// Specifies a list of words that you expect to be used during a conversation with your bot.
	CustomVocabularyItems interface{} `field:"required" json:"customVocabularyItems" yaml:"customVocabularyItems"`
}

