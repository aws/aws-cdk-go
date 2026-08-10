package interfacesawstranscribe


// A reference to a MedicalTranscriptionJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   medicalTranscriptionJobReference := &MedicalTranscriptionJobReference{
//   	MedicalTranscriptionJobArn: jsii.String("medicalTranscriptionJobArn"),
//   }
//
type MedicalTranscriptionJobReference struct {
	// The Arn of the MedicalTranscriptionJob resource.
	MedicalTranscriptionJobArn *string `field:"required" json:"medicalTranscriptionJobArn" yaml:"medicalTranscriptionJobArn"`
}

