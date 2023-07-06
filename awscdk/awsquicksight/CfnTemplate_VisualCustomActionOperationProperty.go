package awsquicksight


// The operation that is defined by the custom action.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualCustomActionOperationProperty := &VisualCustomActionOperationProperty{
//   	FilterOperation: &CustomActionFilterOperationProperty{
//   		SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   			SelectedColumns: []interface{}{
//   				&ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   			SelectedFields: []*string{
//   				jsii.String("selectedFields"),
//   			},
//   		},
//   		TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   			SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   				TargetVisualOptions: jsii.String("targetVisualOptions"),
//   				TargetVisuals: []*string{
//   					jsii.String("targetVisuals"),
//   				},
//   			},
//   		},
//   	},
//   	NavigationOperation: &CustomActionNavigationOperationProperty{
//   		LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   			TargetSheetId: jsii.String("targetSheetId"),
//   		},
//   	},
//   	SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   		ParameterValueConfigurations: []interface{}{
//   			&SetParameterValueConfigurationProperty{
//   				DestinationParameterName: jsii.String("destinationParameterName"),
//   				Value: &DestinationParameterValueConfigurationProperty{
//   					CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   						CustomValues: &CustomParameterValuesProperty{
//   							DateTimeValues: []*string{
//   								jsii.String("dateTimeValues"),
//   							},
//   							DecimalValues: []interface{}{
//   								jsii.Number(123),
//   							},
//   							IntegerValues: []interface{}{
//   								jsii.Number(123),
//   							},
//   							StringValues: []*string{
//   								jsii.String("stringValues"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						IncludeNullValue: jsii.Boolean(false),
//   					},
//   					SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   					SourceColumn: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceField: jsii.String("sourceField"),
//   					SourceParameterName: jsii.String("sourceParameterName"),
//   				},
//   			},
//   		},
//   	},
//   	UrlOperation: &CustomActionURLOperationProperty{
//   		UrlTarget: jsii.String("urlTarget"),
//   		UrlTemplate: jsii.String("urlTemplate"),
//   	},
//   }
//
type CfnTemplate_VisualCustomActionOperationProperty struct {
	// The filter operation that filters data included in a visual or in an entire sheet.
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// The navigation operation that navigates between different sheets in the same analysis.
	NavigationOperation interface{} `field:"optional" json:"navigationOperation" yaml:"navigationOperation"`
	// The set parameter operation that sets parameters in custom action.
	SetParametersOperation interface{} `field:"optional" json:"setParametersOperation" yaml:"setParametersOperation"`
	// The URL operation that opens a link to another webpage.
	UrlOperation interface{} `field:"optional" json:"urlOperation" yaml:"urlOperation"`
}

