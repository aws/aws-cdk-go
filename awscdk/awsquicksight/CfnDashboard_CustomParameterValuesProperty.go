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
type CfnDashboard_CustomParameterValuesProperty struct {
	// `CfnDashboard.CustomParameterValuesProperty.DateTimeValues`.
	DateTimeValues *[]*string `field:"optional" json:"dateTimeValues" yaml:"dateTimeValues"`
	// `CfnDashboard.CustomParameterValuesProperty.DecimalValues`.
	DecimalValues interface{} `field:"optional" json:"decimalValues" yaml:"decimalValues"`
	// `CfnDashboard.CustomParameterValuesProperty.IntegerValues`.
	IntegerValues interface{} `field:"optional" json:"integerValues" yaml:"integerValues"`
	// `CfnDashboard.CustomParameterValuesProperty.StringValues`.
	StringValues *[]*string `field:"optional" json:"stringValues" yaml:"stringValues"`
}

