package awsquicksight


// A list of Amazon QuickSight parameters and the list's override values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &ParametersProperty{
//   	DateTimeParameters: []interface{}{
//   		&DateTimeParameterProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	DecimalParameters: []interface{}{
//   		&DecimalParameterProperty{
//   			Name: jsii.String("name"),
//   			Values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	IntegerParameters: []interface{}{
//   		&IntegerParameterProperty{
//   			Name: jsii.String("name"),
//   			Values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	StringParameters: []interface{}{
//   		&StringParameterProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
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

