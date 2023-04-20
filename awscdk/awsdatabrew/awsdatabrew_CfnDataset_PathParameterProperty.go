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
//   pathParameterProperty := &PathParameterProperty{
//   	DatasetParameter: &DatasetParameterProperty{
//   		Name: jsii.String("name"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		CreateColumn: jsii.Boolean(false),
//   		DatetimeOptions: &DatetimeOptionsProperty{
//   			Format: jsii.String("format"),
//
//   			// the properties below are optional
//   			LocaleCode: jsii.String("localeCode"),
//   			TimezoneOffset: jsii.String("timezoneOffset"),
//   		},
//   		Filter: &FilterExpressionProperty{
//   			Expression: jsii.String("expression"),
//   			ValuesMap: []interface{}{
//   				&FilterValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   	},
//   	PathParameterName: jsii.String("pathParameterName"),
//   }
//
type CfnDataset_PathParameterProperty struct {
	// The path parameter definition.
	DatasetParameter interface{} `field:"required" json:"datasetParameter" yaml:"datasetParameter"`
	// The name of the path parameter.
	PathParameterName *string `field:"required" json:"pathParameterName" yaml:"pathParameterName"`
}

