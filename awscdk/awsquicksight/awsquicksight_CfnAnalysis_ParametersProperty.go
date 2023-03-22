package awsquicksight


// A list of Amazon QuickSight parameters and the list's override values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &parametersProperty{
//   	dateTimeParameters: []interface{}{
//   		&dateTimeParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	decimalParameters: []interface{}{
//   		&decimalParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	integerParameters: []interface{}{
//   		&integerParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	stringParameters: []interface{}{
//   		&stringParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_ParametersProperty struct {
	// The parameters that have a data type of date-time.
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// The parameters that have a data type of decimal.
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// The parameters that have a data type of integer.
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// The parameters that have a data type of string.
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

