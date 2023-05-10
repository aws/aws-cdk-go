package awsquicksight


// The value of a time range filter.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
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
type CfnDashboard_TimeRangeFilterValueProperty struct {
	// The parameter type input value.
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The rolling date input value.
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// The static input value.
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
}

