package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualCustomActionOperationProperty := &VisualCustomActionOperationProperty{
//   	FilterOperation: &CustomActionFilterOperationProperty{
//   		SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
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
	// `CfnTemplate.VisualCustomActionOperationProperty.FilterOperation`.
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// `CfnTemplate.VisualCustomActionOperationProperty.NavigationOperation`.
	NavigationOperation interface{} `field:"optional" json:"navigationOperation" yaml:"navigationOperation"`
	// `CfnTemplate.VisualCustomActionOperationProperty.SetParametersOperation`.
	SetParametersOperation interface{} `field:"optional" json:"setParametersOperation" yaml:"setParametersOperation"`
	// `CfnTemplate.VisualCustomActionOperationProperty.URLOperation`.
	UrlOperation interface{} `field:"optional" json:"urlOperation" yaml:"urlOperation"`
}

