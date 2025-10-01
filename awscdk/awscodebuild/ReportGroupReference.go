package awscodebuild


// A reference to a ReportGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportGroupReference := &ReportGroupReference{
//   	ReportGroupArn: jsii.String("reportGroupArn"),
//   	ReportGroupId: jsii.String("reportGroupId"),
//   }
//
type ReportGroupReference struct {
	// The ARN of the ReportGroup resource.
	ReportGroupArn *string `field:"required" json:"reportGroupArn" yaml:"reportGroupArn"`
	// The Id of the ReportGroup resource.
	ReportGroupId *string `field:"required" json:"reportGroupId" yaml:"reportGroupId"`
}

