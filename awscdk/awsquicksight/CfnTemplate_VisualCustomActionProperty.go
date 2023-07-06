package awsquicksight


// A custom action defined on a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualCustomActionProperty := &VisualCustomActionProperty{
//   	ActionOperations: []interface{}{
//   		&VisualCustomActionOperationProperty{
//   			FilterOperation: &CustomActionFilterOperationProperty{
//   				SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   					SelectedColumns: []interface{}{
//   						&ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   					},
//   					SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   					SelectedFields: []*string{
//   						jsii.String("selectedFields"),
//   					},
//   				},
//   				TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   					SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   						TargetVisualOptions: jsii.String("targetVisualOptions"),
//   						TargetVisuals: []*string{
//   							jsii.String("targetVisuals"),
//   						},
//   					},
//   				},
//   			},
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
type CfnTemplate_VisualCustomActionProperty struct {
	// A list of `VisualCustomActionOperations` .
	//
	// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
	ActionOperations interface{} `field:"required" json:"actionOperations" yaml:"actionOperations"`
	// The ID of the `VisualCustomAction` .
	CustomActionId *string `field:"required" json:"customActionId" yaml:"customActionId"`
	// The name of the `VisualCustomAction` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The trigger of the `VisualCustomAction` .
	//
	// Valid values are defined as follows:
	//
	// - `DATA_POINT_CLICK` : Initiates a custom action by a left pointer click on a data point.
	// - `DATA_POINT_MENU` : Initiates a custom action by right pointer click from the menu.
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// The status of the `VisualCustomAction` .
	Status *string `field:"optional" json:"status" yaml:"status"`
}

