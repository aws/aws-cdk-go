package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterRelativeDateTimeControlProperty := &FilterRelativeDateTimeControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &RelativeDateTimeControlDisplayOptionsProperty{
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
//   }
//
type CfnDashboard_FilterRelativeDateTimeControlProperty struct {
	// `CfnDashboard.FilterRelativeDateTimeControlProperty.FilterControlId`.
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// `CfnDashboard.FilterRelativeDateTimeControlProperty.SourceFilterId`.
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// `CfnDashboard.FilterRelativeDateTimeControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnDashboard.FilterRelativeDateTimeControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

