package mixinsawsappmesh


// An object that represents the outlier detection for a virtual node's listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-outlierdetection.html
//
type CfnVirtualNodePropsMixin_OutlierDetectionProperty struct {
	// The base amount of time for which a host is ejected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-outlierdetection.html#cfn-appmesh-virtualnode-outlierdetection-baseejectionduration
	//
	BaseEjectionDuration interface{} `field:"optional" json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-outlierdetection.html#cfn-appmesh-virtualnode-outlierdetection-interval
	//
	Interval interface{} `field:"optional" json:"interval" yaml:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at least one host regardless of the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-outlierdetection.html#cfn-appmesh-virtualnode-outlierdetection-maxejectionpercent
	//
	MaxEjectionPercent *float64 `field:"optional" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Number of consecutive `5xx` errors required for ejection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-outlierdetection.html#cfn-appmesh-virtualnode-outlierdetection-maxservererrors
	//
	MaxServerErrors *float64 `field:"optional" json:"maxServerErrors" yaml:"maxServerErrors"`
}

