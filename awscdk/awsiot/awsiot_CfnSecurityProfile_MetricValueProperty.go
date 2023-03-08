package awsiot


// The value to be compared with the `metric` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricValueProperty := &metricValueProperty{
//   	cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	count: jsii.String("count"),
//   	number: jsii.Number(123),
//   	numbers: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ports: []interface{}{
//   		jsii.Number(123),
//   	},
//   	strings: []*string{
//   		jsii.String("strings"),
//   	},
//   }
//
type CfnSecurityProfile_MetricValueProperty struct {
	// If the `comparisonOperator` calls for a set of CIDRs, use this to specify that set to be compared with the `metric` .
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// If the `comparisonOperator` calls for a numeric value, use this to specify that numeric value to be compared with the `metric` .
	Count *string `field:"optional" json:"count" yaml:"count"`
	// The numeric values of a metric.
	Number *float64 `field:"optional" json:"number" yaml:"number"`
	// The numeric value of a metric.
	Numbers interface{} `field:"optional" json:"numbers" yaml:"numbers"`
	// If the `comparisonOperator` calls for a set of ports, use this to specify that set to be compared with the `metric` .
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// The string values of a metric.
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
}

