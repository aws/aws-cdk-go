package interfacesawsauditmanager


// A reference to a AssessmentFramework resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentFrameworkReference := &AssessmentFrameworkReference{
//   	AssessmentFrameworkArn: jsii.String("assessmentFrameworkArn"),
//   }
//
type AssessmentFrameworkReference struct {
	// The Arn of the AssessmentFramework resource.
	AssessmentFrameworkArn *string `field:"required" json:"assessmentFrameworkArn" yaml:"assessmentFrameworkArn"`
}

