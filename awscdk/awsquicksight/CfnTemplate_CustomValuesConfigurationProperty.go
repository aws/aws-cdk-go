package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customValuesConfigurationProperty := &CustomValuesConfigurationProperty{
//   	CustomValues: &CustomParameterValuesProperty{
//   		DateTimeValues: []*string{
//   			jsii.String("dateTimeValues"),
//   		},
//   		DecimalValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IntegerValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		StringValues: []*string{
//   			jsii.String("stringValues"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	IncludeNullValue: jsii.Boolean(false),
//   }
//
type CfnTemplate_CustomValuesConfigurationProperty struct {
	// `CfnTemplate.CustomValuesConfigurationProperty.CustomValues`.
	CustomValues interface{} `field:"required" json:"customValues" yaml:"customValues"`
	// `CfnTemplate.CustomValuesConfigurationProperty.IncludeNullValue`.
	IncludeNullValue interface{} `field:"optional" json:"includeNullValue" yaml:"includeNullValue"`
}

