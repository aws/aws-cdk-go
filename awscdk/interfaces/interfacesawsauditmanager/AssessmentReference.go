package interfacesawsauditmanager


// A reference to a Assessment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentReference := &AssessmentReference{
//   	AssessmentArn: jsii.String("assessmentArn"),
//   	AssessmentId: jsii.String("assessmentId"),
//   }
//
type AssessmentReference struct {
	// The ARN of the Assessment resource.
	AssessmentArn *string `field:"required" json:"assessmentArn" yaml:"assessmentArn"`
	// The AssessmentId of the Assessment resource.
	AssessmentId *string `field:"required" json:"assessmentId" yaml:"assessmentId"`
}

