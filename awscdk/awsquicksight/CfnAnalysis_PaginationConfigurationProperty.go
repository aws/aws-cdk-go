package awsquicksight


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
	// `CfnAnalysis.PaginationConfigurationProperty.PageNumber`.
	PageNumber *float64 `field:"required" json:"pageNumber" yaml:"pageNumber"`
	// `CfnAnalysis.PaginationConfigurationProperty.PageSize`.
	PageSize *float64 `field:"required" json:"pageSize" yaml:"pageSize"`
}

