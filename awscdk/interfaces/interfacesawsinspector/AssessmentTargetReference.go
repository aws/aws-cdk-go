package interfacesawsinspector


// A reference to a AssessmentTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentTargetReference := &AssessmentTargetReference{
//   	AssessmentTargetArn: jsii.String("assessmentTargetArn"),
//   }
//
type AssessmentTargetReference struct {
	// The Arn of the AssessmentTarget resource.
	AssessmentTargetArn *string `field:"required" json:"assessmentTargetArn" yaml:"assessmentTargetArn"`
}

