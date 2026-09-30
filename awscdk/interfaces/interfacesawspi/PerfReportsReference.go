package interfacesawspi


// A reference to a PerfReports resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   perfReportsReference := &PerfReportsReference{
//   	PerfReportsArn: jsii.String("perfReportsArn"),
//   }
//
type PerfReportsReference struct {
	// The Arn of the PerfReports resource.
	PerfReportsArn *string `field:"required" json:"perfReportsArn" yaml:"perfReportsArn"`
}

