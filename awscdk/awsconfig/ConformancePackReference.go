package awsconfig


// A reference to a ConformancePack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conformancePackReference := &ConformancePackReference{
//   	ConformancePackName: jsii.String("conformancePackName"),
//   }
//
type ConformancePackReference struct {
	// The ConformancePackName of the ConformancePack resource.
	ConformancePackName *string `field:"required" json:"conformancePackName" yaml:"conformancePackName"`
}

