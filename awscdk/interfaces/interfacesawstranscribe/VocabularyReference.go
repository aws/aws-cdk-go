package interfacesawstranscribe


// A reference to a Vocabulary resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vocabularyReference := &VocabularyReference{
//   	VocabularyArn: jsii.String("vocabularyArn"),
//   }
//
type VocabularyReference struct {
	// The Arn of the Vocabulary resource.
	VocabularyArn *string `field:"required" json:"vocabularyArn" yaml:"vocabularyArn"`
}

