package awsdatabrew


// Represents a single entry in the path parameters of a dataset.
//
// Each `PathParameter` consists of a name and a parameter definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathParameterProperty := &pathParameterProperty{
//   	datasetParameter: &datasetParameterProperty{
//   		name: jsii.String("name"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		createColumn: jsii.Boolean(false),
//   		datetimeOptions: &datetimeOptionsProperty{
//   			format: jsii.String("format"),
//
//   			// the properties below are optional
//   			localeCode: jsii.String("localeCode"),
//   			timezoneOffset: jsii.String("timezoneOffset"),
//   		},
//   		filter: &filterExpressionProperty{
//   			expression: jsii.String("expression"),
//   			valuesMap: []interface{}{
//   				&filterValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   	},
//   	pathParameterName: jsii.String("pathParameterName"),
//   }
//
type CfnDataset_PathParameterProperty struct {
	// The path parameter definition.
	DatasetParameter interface{} `field:"required" json:"datasetParameter" yaml:"datasetParameter"`
	// The name of the path parameter.
	PathParameterName *string `field:"required" json:"pathParameterName" yaml:"pathParameterName"`
}

