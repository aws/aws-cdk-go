package awsquicksight


// A control from a date filter that is used to specify date and time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDateTimePickerControlProperty := &FilterDateTimePickerControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
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
//   	Type: jsii.String("type"),
//   }
//
type CfnDashboard_FilterDateTimePickerControlProperty struct {
	// The ID of the `FilterDateTimePickerControl` .
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// The source filter ID of the `FilterDateTimePickerControl` .
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The title of the `FilterDateTimePickerControl` .
	Title *string `field:"required" json:"title" yaml:"title"`
	// The display options of a control.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The date time picker type of a `FilterDateTimePickerControl` . Choose one of the following options:.
	//
	// - `SINGLE_VALUED` : The filter condition is a fixed date.
	// - `DATE_RANGE` : The filter condition is a date time range.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

