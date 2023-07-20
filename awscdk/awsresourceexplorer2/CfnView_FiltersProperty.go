package awsresourceexplorer2


// An object with a `FilterString` that specifies which resources to include in the results of queries made using this view.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filtersProperty := &FiltersProperty{
//   	FilterString: jsii.String("filterString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-filters.html
//
type CfnView_FiltersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-filters.html#cfn-resourceexplorer2-view-filters-filterstring
	//
	FilterString *string `field:"required" json:"filterString" yaml:"filterString"`
}

