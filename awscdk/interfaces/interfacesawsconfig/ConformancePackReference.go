package interfacesawsconfig


// A reference to a ConformancePack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conformancePackReference := &ConformancePackReference{
//   	ConformancePackArn: jsii.String("conformancePackArn"),
//   	ConformancePackName: jsii.String("conformancePackName"),
//   }
//
type ConformancePackReference struct {
	// The ARN of the ConformancePack resource.
	ConformancePackArn *string `field:"required" json:"conformancePackArn" yaml:"conformancePackArn"`
	// The ConformancePackName of the ConformancePack resource.
	ConformancePackName *string `field:"required" json:"conformancePackName" yaml:"conformancePackName"`
}

