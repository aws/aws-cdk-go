package awsquicksight


// The customized parameter values.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customParameterValuesProperty := &CustomParameterValuesProperty{
//   	DateTimeValues: []*string{
//   		jsii.String("dateTimeValues"),
//   	},
//   	DecimalValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	IntegerValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	StringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   }
//
type CfnAnalysis_CustomParameterValuesProperty struct {
	// A list of datetime-type parameter values.
	DateTimeValues *[]*string `field:"optional" json:"dateTimeValues" yaml:"dateTimeValues"`
	// A list of decimal-type parameter values.
	DecimalValues interface{} `field:"optional" json:"decimalValues" yaml:"decimalValues"`
	// A list of integer-type parameter values.
	IntegerValues interface{} `field:"optional" json:"integerValues" yaml:"integerValues"`
	// A list of string-type parameter values.
	StringValues *[]*string `field:"optional" json:"stringValues" yaml:"stringValues"`
}

