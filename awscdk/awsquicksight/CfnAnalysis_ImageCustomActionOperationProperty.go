package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageCustomActionOperationProperty := &ImageCustomActionOperationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomactionoperation.html
//
type CfnAnalysis_ImageCustomActionOperationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomactionoperation.html#cfn-quicksight-analysis-imagecustomactionoperation-navigationoperation
	//
	NavigationOperation interface{} `field:"optional" json:"navigationOperation" yaml:"navigationOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomactionoperation.html#cfn-quicksight-analysis-imagecustomactionoperation-setparametersoperation
	//
	SetParametersOperation interface{} `field:"optional" json:"setParametersOperation" yaml:"setParametersOperation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomactionoperation.html#cfn-quicksight-analysis-imagecustomactionoperation-urloperation
	//
	UrlOperation interface{} `field:"optional" json:"urlOperation" yaml:"urlOperation"`
}

