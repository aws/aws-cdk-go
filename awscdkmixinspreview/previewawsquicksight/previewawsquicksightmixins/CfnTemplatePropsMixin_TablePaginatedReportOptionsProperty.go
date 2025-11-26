package previewawsquicksightmixins


// The paginated report options for a table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tablePaginatedReportOptionsProperty := &TablePaginatedReportOptionsProperty{
//   	OverflowColumnHeaderVisibility: jsii.String("overflowColumnHeaderVisibility"),
//   	VerticalOverflowVisibility: jsii.String("verticalOverflowVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablepaginatedreportoptions.html
//
type CfnTemplatePropsMixin_TablePaginatedReportOptionsProperty struct {
	// The visibility of repeating header rows on each page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablepaginatedreportoptions.html#cfn-quicksight-template-tablepaginatedreportoptions-overflowcolumnheadervisibility
	//
	OverflowColumnHeaderVisibility *string `field:"optional" json:"overflowColumnHeaderVisibility" yaml:"overflowColumnHeaderVisibility"`
	// The visibility of printing table overflow across pages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablepaginatedreportoptions.html#cfn-quicksight-template-tablepaginatedreportoptions-verticaloverflowvisibility
	//
	VerticalOverflowVisibility *string `field:"optional" json:"verticalOverflowVisibility" yaml:"verticalOverflowVisibility"`
}

