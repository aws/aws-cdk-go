package awsinspectorv2


// An object that describes the details of a port range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeFilterProperty := &portRangeFilterProperty{
//   	beginInclusive: jsii.Number(123),
//   	endInclusive: jsii.Number(123),
//   }
//
type CfnFilter_PortRangeFilterProperty struct {
	// The port number the port range begins at.
	BeginInclusive *float64 `field:"optional" json:"beginInclusive" yaml:"beginInclusive"`
	// The port number the port range ends at.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
}

