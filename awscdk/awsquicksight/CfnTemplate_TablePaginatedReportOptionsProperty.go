package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tablePaginatedReportOptionsProperty := &TablePaginatedReportOptionsProperty{
//   	OverflowColumnHeaderVisibility: jsii.String("overflowColumnHeaderVisibility"),
//   	VerticalOverflowVisibility: jsii.String("verticalOverflowVisibility"),
//   }
//
type CfnTemplate_TablePaginatedReportOptionsProperty struct {
	// `CfnTemplate.TablePaginatedReportOptionsProperty.OverflowColumnHeaderVisibility`.
	OverflowColumnHeaderVisibility *string `field:"optional" json:"overflowColumnHeaderVisibility" yaml:"overflowColumnHeaderVisibility"`
	// `CfnTemplate.TablePaginatedReportOptionsProperty.VerticalOverflowVisibility`.
	VerticalOverflowVisibility *string `field:"optional" json:"verticalOverflowVisibility" yaml:"verticalOverflowVisibility"`
}

