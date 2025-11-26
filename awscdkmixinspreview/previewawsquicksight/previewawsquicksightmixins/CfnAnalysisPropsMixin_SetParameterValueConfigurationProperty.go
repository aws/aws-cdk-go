package previewawsquicksightmixins


// The configuration of adding parameters in action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   setParameterValueConfigurationProperty := &SetParameterValueConfigurationProperty{
//   	DestinationParameterName: jsii.String("destinationParameterName"),
//   	Value: &DestinationParameterValueConfigurationProperty{
//   		CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   			CustomValues: &CustomParameterValuesProperty{
//   				DateTimeValues: []*string{
//   					jsii.String("dateTimeValues"),
//   				},
//   				DecimalValues: []interface{}{
//   					jsii.Number(123),
//   				},
//   				IntegerValues: []interface{}{
//   					jsii.Number(123),
//   				},
//   				StringValues: []*string{
//   					jsii.String("stringValues"),
//   				},
//   			},
//   			IncludeNullValue: jsii.Boolean(false),
//   		},
//   		SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   		SourceColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		SourceField: jsii.String("sourceField"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-setparametervalueconfiguration.html
//
type CfnAnalysisPropsMixin_SetParameterValueConfigurationProperty struct {
	// The destination parameter name of the `SetParameterValueConfiguration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-setparametervalueconfiguration.html#cfn-quicksight-analysis-setparametervalueconfiguration-destinationparametername
	//
	DestinationParameterName *string `field:"optional" json:"destinationParameterName" yaml:"destinationParameterName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-setparametervalueconfiguration.html#cfn-quicksight-analysis-setparametervalueconfiguration-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

