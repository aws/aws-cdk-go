package interfacesawstranscribe


// A reference to a VocabularyFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vocabularyFilterReference := &VocabularyFilterReference{
//   	VocabularyFilterArn: jsii.String("vocabularyFilterArn"),
//   }
//
type VocabularyFilterReference struct {
	// The Arn of the VocabularyFilter resource.
	VocabularyFilterArn *string `field:"required" json:"vocabularyFilterArn" yaml:"vocabularyFilterArn"`
}

