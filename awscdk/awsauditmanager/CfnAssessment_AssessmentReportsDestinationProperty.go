package awsauditmanager


// The `AssessmentReportsDestination` property type specifies the location in which AWS Audit Manager saves assessment reports for the given assessment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assessmentReportsDestinationProperty := &AssessmentReportsDestinationProperty{
//   	Destination: jsii.String("destination"),
//   	DestinationType: jsii.String("destinationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html
//
type CfnAssessment_AssessmentReportsDestinationProperty struct {
	// The destination bucket where Audit Manager stores assessment reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html#cfn-auditmanager-assessment-assessmentreportsdestination-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The destination type, such as Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html#cfn-auditmanager-assessment-assessmentreportsdestination-destinationtype
	//
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
}

