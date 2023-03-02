package awslex


// Specifies an entry in a custom vocabulary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customVocabularyItemProperty := &customVocabularyItemProperty{
//   	phrase: jsii.String("phrase"),
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type CfnBot_CustomVocabularyItemProperty struct {
	// Specifies 1 - 4 words that should be recognized.
	Phrase *string `field:"required" json:"phrase" yaml:"phrase"`
	// Specifies the degree to which the phrase recognition is boosted.
	//
	// The default value is 1.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

