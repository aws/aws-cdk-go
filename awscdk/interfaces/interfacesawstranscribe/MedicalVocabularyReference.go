package interfacesawstranscribe


// A reference to a MedicalVocabulary resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   medicalVocabularyReference := &MedicalVocabularyReference{
//   	MedicalVocabularyArn: jsii.String("medicalVocabularyArn"),
//   }
//
type MedicalVocabularyReference struct {
	// The Arn of the MedicalVocabulary resource.
	MedicalVocabularyArn *string `field:"required" json:"medicalVocabularyArn" yaml:"medicalVocabularyArn"`
}

