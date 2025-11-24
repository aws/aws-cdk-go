package mixinsawsquicksight


// The set parameter operation that sets parameters in custom action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionSetParametersOperationProperty := &CustomActionSetParametersOperationProperty{
//   	ParameterValueConfigurations: []interface{}{
//   		&SetParameterValueConfigurationProperty{
//   			DestinationParameterName: jsii.String("destinationParameterName"),
//   			Value: &DestinationParameterValueConfigurationProperty{
//   				CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   					CustomValues: &CustomParameterValuesProperty{
//   						DateTimeValues: []*string{
//   							jsii.String("dateTimeValues"),
//   						},
//   						DecimalValues: []interface{}{
//   							jsii.Number(123),
//   						},
//   						IntegerValues: []interface{}{
//   							jsii.Number(123),
//   						},
//   						StringValues: []*string{
//   							jsii.String("stringValues"),
//   						},
//   					},
//   					IncludeNullValue: jsii.Boolean(false),
//   				},
//   				SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   				SourceColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				SourceField: jsii.String("sourceField"),
//   				SourceParameterName: jsii.String("sourceParameterName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customactionsetparametersoperation.html
//
type CfnDashboardPropsMixin_CustomActionSetParametersOperationProperty struct {
	// The parameter that determines the value configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customactionsetparametersoperation.html#cfn-quicksight-dashboard-customactionsetparametersoperation-parametervalueconfigurations
	//
	ParameterValueConfigurations interface{} `field:"optional" json:"parameterValueConfigurations" yaml:"parameterValueConfigurations"`
}

