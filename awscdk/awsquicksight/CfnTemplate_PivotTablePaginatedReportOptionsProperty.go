package awsquicksight


// The paginated report options for a pivot table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTablePaginatedReportOptionsProperty := &PivotTablePaginatedReportOptionsProperty{
//   	OverflowColumnHeaderVisibility: jsii.String("overflowColumnHeaderVisibility"),
//   	VerticalOverflowVisibility: jsii.String("verticalOverflowVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablepaginatedreportoptions.html
//
type CfnTemplate_PivotTablePaginatedReportOptionsProperty struct {
	// The visibility of the repeating header rows on each page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablepaginatedreportoptions.html#cfn-quicksight-template-pivottablepaginatedreportoptions-overflowcolumnheadervisibility
	//
	OverflowColumnHeaderVisibility *string `field:"optional" json:"overflowColumnHeaderVisibility" yaml:"overflowColumnHeaderVisibility"`
	// The visibility of the printing table overflow across pages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablepaginatedreportoptions.html#cfn-quicksight-template-pivottablepaginatedreportoptions-verticaloverflowvisibility
	//
	VerticalOverflowVisibility *string `field:"optional" json:"verticalOverflowVisibility" yaml:"verticalOverflowVisibility"`
}

