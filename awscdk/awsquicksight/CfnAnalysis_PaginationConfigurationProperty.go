package awsquicksight


// The pagination configuration for a table visual or boxplot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paginationConfigurationProperty := &PaginationConfigurationProperty{
//   	PageNumber: jsii.Number(123),
//   	PageSize: jsii.Number(123),
//   }
//
type CfnAnalysis_PaginationConfigurationProperty struct {
	// Indicates the page number.
	PageNumber *float64 `field:"required" json:"pageNumber" yaml:"pageNumber"`
	// Indicates how many items render in one page.
	PageSize *float64 `field:"required" json:"pageSize" yaml:"pageSize"`
}

