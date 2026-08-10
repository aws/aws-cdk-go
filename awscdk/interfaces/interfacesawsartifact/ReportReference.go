package interfacesawsartifact


// A reference to a Report resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportReference := &ReportReference{
//   	ReportArn: jsii.String("reportArn"),
//   }
//
type ReportReference struct {
	// The Arn of the Report resource.
	ReportArn *string `field:"required" json:"reportArn" yaml:"reportArn"`
}

