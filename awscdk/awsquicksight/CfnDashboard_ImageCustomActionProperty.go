package awsquicksight


// A custom action defined on an image.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html
//
type CfnDashboard_ImageCustomActionProperty struct {
	// A list of `ImageCustomActionOperations` .
	//
	// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html#cfn-quicksight-dashboard-imagecustomaction-actionoperations
	//
	ActionOperations interface{} `field:"required" json:"actionOperations" yaml:"actionOperations"`
	// The ID of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html#cfn-quicksight-dashboard-imagecustomaction-customactionid
	//
	CustomActionId *string `field:"required" json:"customActionId" yaml:"customActionId"`
	// The name of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html#cfn-quicksight-dashboard-imagecustomaction-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The trigger of the `VisualCustomAction` .
	//
	// Valid values are defined as follows:
	//
	// - `CLICK` : Initiates a custom action by a left pointer click on a data point.
	// - `MENU` : Initiates a custom action by right pointer click from the menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html#cfn-quicksight-dashboard-imagecustomaction-trigger
	//
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// The status of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagecustomaction.html#cfn-quicksight-dashboard-imagecustomaction-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

