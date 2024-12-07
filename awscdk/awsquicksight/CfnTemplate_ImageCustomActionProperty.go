package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageCustomActionProperty := &ImageCustomActionProperty{
//   	ActionOperations: []interface{}{
//   		&ImageCustomActionOperationProperty{
//   			NavigationOperation: &CustomActionNavigationOperationProperty{
//   				LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   					TargetSheetId: jsii.String("targetSheetId"),
//   				},
//   			},
//   			SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   				ParameterValueConfigurations: []interface{}{
//   					&SetParameterValueConfigurationProperty{
//   						DestinationParameterName: jsii.String("destinationParameterName"),
//   						Value: &DestinationParameterValueConfigurationProperty{
//   							CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   								CustomValues: &CustomParameterValuesProperty{
//   									DateTimeValues: []*string{
//   										jsii.String("dateTimeValues"),
//   									},
//   									DecimalValues: []interface{}{
//   										jsii.Number(123),
//   									},
//   									IntegerValues: []interface{}{
//   										jsii.Number(123),
//   									},
//   									StringValues: []*string{
//   										jsii.String("stringValues"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								IncludeNullValue: jsii.Boolean(false),
//   							},
//   							SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   							SourceColumn: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							SourceField: jsii.String("sourceField"),
//   							SourceParameterName: jsii.String("sourceParameterName"),
//   						},
//   					},
//   				},
//   			},
//   			UrlOperation: &CustomActionURLOperationProperty{
//   				UrlTarget: jsii.String("urlTarget"),
//   				UrlTemplate: jsii.String("urlTemplate"),
//   			},
//   		},
//   	},
//   	CustomActionId: jsii.String("customActionId"),
//   	Name: jsii.String("name"),
//   	Trigger: jsii.String("trigger"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html
//
type CfnTemplate_ImageCustomActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html#cfn-quicksight-template-imagecustomaction-actionoperations
	//
	ActionOperations interface{} `field:"required" json:"actionOperations" yaml:"actionOperations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html#cfn-quicksight-template-imagecustomaction-customactionid
	//
	CustomActionId *string `field:"required" json:"customActionId" yaml:"customActionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html#cfn-quicksight-template-imagecustomaction-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html#cfn-quicksight-template-imagecustomaction-trigger
	//
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomaction.html#cfn-quicksight-template-imagecustomaction-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

