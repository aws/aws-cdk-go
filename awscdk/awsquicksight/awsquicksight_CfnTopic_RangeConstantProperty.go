package awsquicksight


// The value of the constant that is used to specify the endpoints of a range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangeConstantProperty := &RangeConstantProperty{
//   	Maximum: jsii.String("maximum"),
//   	Minimum: jsii.String("minimum"),
//   }
//
type CfnTopic_RangeConstantProperty struct {
	// The maximum value for a range constant.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// The minimum value for a range constant.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
}

