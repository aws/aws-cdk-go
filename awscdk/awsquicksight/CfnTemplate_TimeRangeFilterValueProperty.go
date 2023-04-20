package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeFilterValueProperty := &TimeRangeFilterValueProperty{
//   	Parameter: jsii.String("parameter"),
//   	RollingDate: &RollingDateConfigurationProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	StaticValue: jsii.String("staticValue"),
//   }
//
type CfnTemplate_TimeRangeFilterValueProperty struct {
	// `CfnTemplate.TimeRangeFilterValueProperty.Parameter`.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// `CfnTemplate.TimeRangeFilterValueProperty.RollingDate`.
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// `CfnTemplate.TimeRangeFilterValueProperty.StaticValue`.
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
}

