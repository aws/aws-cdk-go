package awsquicksight


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
type CfnTemplate_CustomParameterValuesProperty struct {
	// `CfnTemplate.CustomParameterValuesProperty.DateTimeValues`.
	DateTimeValues *[]*string `field:"optional" json:"dateTimeValues" yaml:"dateTimeValues"`
	// `CfnTemplate.CustomParameterValuesProperty.DecimalValues`.
	DecimalValues interface{} `field:"optional" json:"decimalValues" yaml:"decimalValues"`
	// `CfnTemplate.CustomParameterValuesProperty.IntegerValues`.
	IntegerValues interface{} `field:"optional" json:"integerValues" yaml:"integerValues"`
	// `CfnTemplate.CustomParameterValuesProperty.StringValues`.
	StringValues *[]*string `field:"optional" json:"stringValues" yaml:"stringValues"`
}

