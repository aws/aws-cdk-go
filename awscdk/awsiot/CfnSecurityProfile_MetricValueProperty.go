package awsiot


// The value to be compared with the `metric` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricValueProperty := &MetricValueProperty{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	Count: jsii.String("count"),
//   	Number: jsii.Number(123),
//   	Numbers: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Ports: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Strings: []*string{
//   		jsii.String("strings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html
//
type CfnSecurityProfile_MetricValueProperty struct {
	// If the `comparisonOperator` calls for a set of CIDRs, use this to specify that set to be compared with the `metric` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-cidrs
	//
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// If the `comparisonOperator` calls for a numeric value, use this to specify that numeric value to be compared with the `metric` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-count
	//
	Count *string `field:"optional" json:"count" yaml:"count"`
	// The numeric values of a metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-number
	//
	Number *float64 `field:"optional" json:"number" yaml:"number"`
	// The numeric value of a metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-numbers
	//
	Numbers interface{} `field:"optional" json:"numbers" yaml:"numbers"`
	// If the `comparisonOperator` calls for a set of ports, use this to specify that set to be compared with the `metric` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-ports
	//
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// The string values of a metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricvalue.html#cfn-iot-securityprofile-metricvalue-strings
	//
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
}

