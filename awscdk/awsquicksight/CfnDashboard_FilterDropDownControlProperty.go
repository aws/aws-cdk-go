package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDropDownControlProperty := &FilterDropDownControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   		SourceControls: []interface{}{
//   			&CascadingControlSourceProperty{
//   				ColumnToMatch: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   			},
//   		},
//   	},
//   	DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   		SelectAllOptions: &ListControlSelectAllOptionsProperty{
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
//   	SelectableValues: &FilterSelectableValuesProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnDashboard_FilterDropDownControlProperty struct {
	// `CfnDashboard.FilterDropDownControlProperty.FilterControlId`.
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// `CfnDashboard.FilterDropDownControlProperty.SourceFilterId`.
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// `CfnDashboard.FilterDropDownControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnDashboard.FilterDropDownControlProperty.CascadingControlConfiguration`.
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// `CfnDashboard.FilterDropDownControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// `CfnDashboard.FilterDropDownControlProperty.SelectableValues`.
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// `CfnDashboard.FilterDropDownControlProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

