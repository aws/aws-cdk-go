package awsauditmanager


// The `AssessmentReportsDestination` property type specifies the location in which AWS Audit Manager saves assessment reports for the given assessment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentReportsDestinationProperty := &assessmentReportsDestinationProperty{
//   	destination: jsii.String("destination"),
//   	destinationType: jsii.String("destinationType"),
//   }
//
type CfnAssessment_AssessmentReportsDestinationProperty struct {
	// The destination of the assessment report.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The destination type, such as Amazon S3.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
}

