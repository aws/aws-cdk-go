package previewawsresourceexplorer2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filtersProperty := &FiltersProperty{
//   	FilterString: jsii.String("filterString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-filters.html
//
type CfnViewPropsMixin_FiltersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-filters.html#cfn-resourceexplorer2-view-filters-filterstring
	//
	FilterString *string `field:"optional" json:"filterString" yaml:"filterString"`
}

