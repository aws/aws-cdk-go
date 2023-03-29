package awsquicksight


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
	// `CfnTemplate.VisualCustomActionProperty.ActionOperations`.
	ActionOperations interface{} `field:"required" json:"actionOperations" yaml:"actionOperations"`
	// `CfnTemplate.VisualCustomActionProperty.CustomActionId`.
	CustomActionId *string `field:"required" json:"customActionId" yaml:"customActionId"`
	// `CfnTemplate.VisualCustomActionProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnTemplate.VisualCustomActionProperty.Trigger`.
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// `CfnTemplate.VisualCustomActionProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

