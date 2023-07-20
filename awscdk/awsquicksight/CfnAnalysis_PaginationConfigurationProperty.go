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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paginationconfiguration.html
//
type CfnAnalysis_PaginationConfigurationProperty struct {
	// Indicates the page number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paginationconfiguration.html#cfn-quicksight-analysis-paginationconfiguration-pagenumber
	//
	PageNumber *float64 `field:"required" json:"pageNumber" yaml:"pageNumber"`
	// Indicates how many items render in one page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paginationconfiguration.html#cfn-quicksight-analysis-paginationconfiguration-pagesize
	//
	PageSize *float64 `field:"required" json:"pageSize" yaml:"pageSize"`
}

