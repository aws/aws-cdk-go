package awsinspector


// A reference to a AssessmentTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentTemplateReference := &AssessmentTemplateReference{
//   	AssessmentTemplateArn: jsii.String("assessmentTemplateArn"),
//   }
//
type AssessmentTemplateReference struct {
	// The Arn of the AssessmentTemplate resource.
	AssessmentTemplateArn *string `field:"required" json:"assessmentTemplateArn" yaml:"assessmentTemplateArn"`
}

