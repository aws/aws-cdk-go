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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-paginationconfiguration.html
//
type CfnTemplate_PaginationConfigurationProperty struct {
	// Indicates the page number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-paginationconfiguration.html#cfn-quicksight-template-paginationconfiguration-pagenumber
	//
	PageNumber *float64 `field:"required" json:"pageNumber" yaml:"pageNumber"`
	// Indicates how many items render in one page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-paginationconfiguration.html#cfn-quicksight-template-paginationconfiguration-pagesize
	//
	PageSize *float64 `field:"required" json:"pageSize" yaml:"pageSize"`
}

