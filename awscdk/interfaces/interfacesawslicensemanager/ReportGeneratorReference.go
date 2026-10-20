package interfacesawslicensemanager


// A reference to a ReportGenerator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportGeneratorReference := &ReportGeneratorReference{
//   	ReportGeneratorArn: jsii.String("reportGeneratorArn"),
//   }
//
type ReportGeneratorReference struct {
	// The Arn of the ReportGenerator resource.
	ReportGeneratorArn *string `field:"required" json:"reportGeneratorArn" yaml:"reportGeneratorArn"`
}

