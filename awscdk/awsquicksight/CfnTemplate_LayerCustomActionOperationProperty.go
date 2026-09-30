package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layerCustomActionOperationProperty := &LayerCustomActionOperationProperty{
//   	FilterOperation: &CustomActionFilterOperationProperty{
//   		SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   			SelectedColumns: []interface{}{
//   				&ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//
//   					// the properties below are optional
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					TopicIdentifier: jsii.String("topicIdentifier"),
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
//
//   						// the properties below are optional
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						TopicIdentifier: jsii.String("topicIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layercustomactionoperation.html
//
type CfnTemplate_LayerCustomActionOperationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layercustomactionoperation.html#cfn-quicksight-template-layercustomactionoperation-filteroperation
	//
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layercustomactionoperation.html#cfn-quicksight-template-layercustomactionoperation-navigationoperation
	//
	NavigationOperation interface{} `field:"optional" json:"navigationOperation" yaml:"navigationOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layercustomactionoperation.html#cfn-quicksight-template-layercustomactionoperation-setparametersoperation
	//
	SetParametersOperation interface{} `field:"optional" json:"setParametersOperation" yaml:"setParametersOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layercustomactionoperation.html#cfn-quicksight-template-layercustomactionoperation-urloperation
	//
	UrlOperation interface{} `field:"optional" json:"urlOperation" yaml:"urlOperation"`
}

