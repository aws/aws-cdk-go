package awsquicksight


// A control to display a text box that is used to enter multiple entries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterTextAreaControlProperty := &FilterTextAreaControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Delimiter: jsii.String("delimiter"),
//   	DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   		PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnDashboard_FilterTextAreaControlProperty struct {
	// The ID of the `FilterTextAreaControl` .
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterTextAreaControl` .
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterTextAreaControl` .
	Title *string `field:"required" json:"title" yaml:"title"`
	// The delimiter that is used to separate the lines in text.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The display options of a control.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

