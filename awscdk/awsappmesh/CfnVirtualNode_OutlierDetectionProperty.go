package awsappmesh


// An object that represents the outlier detection for a virtual node's listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outlierDetectionProperty := &OutlierDetectionProperty{
//   	BaseEjectionDuration: &DurationProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	Interval: &DurationProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	MaxEjectionPercent: jsii.Number(123),
//   	MaxServerErrors: jsii.Number(123),
//   }
//
type CfnVirtualNode_OutlierDetectionProperty struct {
	// The base amount of time for which a host is ejected.
	BaseEjectionDuration interface{} `field:"required" json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	Interval interface{} `field:"required" json:"interval" yaml:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at least one host regardless of the value.
	MaxEjectionPercent *float64 `field:"required" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Number of consecutive `5xx` errors required for ejection.
	MaxServerErrors *float64 `field:"required" json:"maxServerErrors" yaml:"maxServerErrors"`
}

