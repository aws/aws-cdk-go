package awscur


// A reference to a ReportDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportDefinitionReference := &ReportDefinitionReference{
//   	ReportName: jsii.String("reportName"),
//   }
//
type ReportDefinitionReference struct {
	// The ReportName of the ReportDefinition resource.
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
}

