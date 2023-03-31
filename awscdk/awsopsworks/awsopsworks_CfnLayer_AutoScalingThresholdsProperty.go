package awsopsworks


// Describes a load-based auto scaling upscaling or downscaling threshold configuration, which specifies when AWS OpsWorks Stacks starts or stops load-based instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingThresholdsProperty := &autoScalingThresholdsProperty{
//   	cpuThreshold: jsii.Number(123),
//   	ignoreMetricsTime: jsii.Number(123),
//   	instanceCount: jsii.Number(123),
//   	loadThreshold: jsii.Number(123),
//   	memoryThreshold: jsii.Number(123),
//   	thresholdsWaitTime: jsii.Number(123),
//   }
//
type CfnLayer_AutoScalingThresholdsProperty struct {
	// The CPU utilization threshold, as a percent of the available CPU.
	//
	// A value of -1 disables the threshold.
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	//
	// For example, AWS OpsWorks Stacks adds new instances following an upscaling event but the instances won't start reducing the load until they have been booted and configured. There is no point in raising additional scaling events during that operation, which typically takes several minutes. `IgnoreMetricsTime` allows you to direct AWS OpsWorks Stacks to suppress scaling events long enough to get the new instances online.
	IgnoreMetricsTime *float64 `field:"optional" json:"ignoreMetricsTime" yaml:"ignoreMetricsTime"`
	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The load threshold.
	//
	// A value of -1 disables the threshold. For more information about how load is computed, see [Load (computing)](https://docs.aws.amazon.com/http://en.wikipedia.org/wiki/Load_%28computing%29) .
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
	// The memory utilization threshold, as a percent of the available memory.
	//
	// A value of -1 disables the threshold.
	MemoryThreshold *float64 `field:"optional" json:"memoryThreshold" yaml:"memoryThreshold"`
	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime *float64 `field:"optional" json:"thresholdsWaitTime" yaml:"thresholdsWaitTime"`
}

