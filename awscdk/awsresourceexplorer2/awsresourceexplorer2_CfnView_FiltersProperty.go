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
type CfnView_FiltersProperty struct {
	// `CfnView.FiltersProperty.FilterString`.
	FilterString *string `field:"required" json:"filterString" yaml:"filterString"`
}

