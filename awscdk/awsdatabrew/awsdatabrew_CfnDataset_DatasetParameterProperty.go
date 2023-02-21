package awsdatabrew


// Represents a dataset paramater that defines type and conditions for a parameter in the Amazon S3 path of the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetParameterProperty := &datasetParameterProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	createColumn: jsii.Boolean(false),
//   	datetimeOptions: &datetimeOptionsProperty{
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		localeCode: jsii.String("localeCode"),
//   		timezoneOffset: jsii.String("timezoneOffset"),
//   	},
//   	filter: &filterExpressionProperty{
//   		expression: jsii.String("expression"),
//   		valuesMap: []interface{}{
//   			&filterValueProperty{
//   				value: jsii.String("value"),
//   				valueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_DatasetParameterProperty struct {
	// The name of the parameter that is used in the dataset's Amazon S3 path.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the dataset parameter, can be one of a 'String', 'Number' or 'Datetime'.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional boolean value that defines whether the captured value of this parameter should be loaded as an additional column in the dataset.
	CreateColumn interface{} `field:"optional" json:"createColumn" yaml:"createColumn"`
	// Additional parameter options such as a format and a timezone.
	//
	// Required for datetime parameters.
	DatetimeOptions interface{} `field:"optional" json:"datetimeOptions" yaml:"datetimeOptions"`
	// The optional filter expression structure to apply additional matching criteria to the parameter.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

