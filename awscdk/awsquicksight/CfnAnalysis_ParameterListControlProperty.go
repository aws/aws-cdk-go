package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterListControlProperty := &ParameterListControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
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
//   	DisplayOptions: &ListControlDisplayOptionsProperty{
//   		SearchOptions: &ListControlSearchOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
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
//   	SelectableValues: &ParameterSelectableValuesProperty{
//   		LinkToDataSetColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnAnalysis_ParameterListControlProperty struct {
	// `CfnAnalysis.ParameterListControlProperty.ParameterControlId`.
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// `CfnAnalysis.ParameterListControlProperty.SourceParameterName`.
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// `CfnAnalysis.ParameterListControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnAnalysis.ParameterListControlProperty.CascadingControlConfiguration`.
	CascadingControlConfiguration interface{} `field:"optional" json:"cascadingControlConfiguration" yaml:"cascadingControlConfiguration"`
	// `CfnAnalysis.ParameterListControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// `CfnAnalysis.ParameterListControlProperty.SelectableValues`.
	SelectableValues interface{} `field:"optional" json:"selectableValues" yaml:"selectableValues"`
	// `CfnAnalysis.ParameterListControlProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

