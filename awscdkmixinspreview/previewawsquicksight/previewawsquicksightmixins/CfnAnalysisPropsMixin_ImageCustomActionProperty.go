package previewawsquicksightmixins


// A custom action defined on an image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   	Status: jsii.String("status"),
//   	Trigger: jsii.String("trigger"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html
//
type CfnAnalysisPropsMixin_ImageCustomActionProperty struct {
	// A list of `ImageCustomActionOperations` .
	//
	// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html#cfn-quicksight-analysis-imagecustomaction-actionoperations
	//
	ActionOperations interface{} `field:"optional" json:"actionOperations" yaml:"actionOperations"`
	// The ID of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html#cfn-quicksight-analysis-imagecustomaction-customactionid
	//
	CustomActionId *string `field:"optional" json:"customActionId" yaml:"customActionId"`
	// The name of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html#cfn-quicksight-analysis-imagecustomaction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html#cfn-quicksight-analysis-imagecustomaction-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The trigger of the `VisualCustomAction` .
	//
	// Valid values are defined as follows:
	//
	// - `CLICK` : Initiates a custom action by a left pointer click on a data point.
	// - `MENU` : Initiates a custom action by right pointer click from the menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagecustomaction.html#cfn-quicksight-analysis-imagecustomaction-trigger
	//
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
}

