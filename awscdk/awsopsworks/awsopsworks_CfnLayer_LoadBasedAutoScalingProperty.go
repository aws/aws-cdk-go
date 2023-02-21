package awsopsworks


// Describes a layer's load-based auto scaling configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBasedAutoScalingProperty := &loadBasedAutoScalingProperty{
//   	downScaling: &autoScalingThresholdsProperty{
//   		cpuThreshold: jsii.Number(123),
//   		ignoreMetricsTime: jsii.Number(123),
//   		instanceCount: jsii.Number(123),
//   		loadThreshold: jsii.Number(123),
//   		memoryThreshold: jsii.Number(123),
//   		thresholdsWaitTime: jsii.Number(123),
//   	},
//   	enable: jsii.Boolean(false),
//   	upScaling: &autoScalingThresholdsProperty{
//   		cpuThreshold: jsii.Number(123),
//   		ignoreMetricsTime: jsii.Number(123),
//   		instanceCount: jsii.Number(123),
//   		loadThreshold: jsii.Number(123),
//   		memoryThreshold: jsii.Number(123),
//   		thresholdsWaitTime: jsii.Number(123),
//   	},
//   }
//
type CfnLayer_LoadBasedAutoScalingProperty struct {
	// An `AutoScalingThresholds` object that describes the downscaling configuration, which defines how and when AWS OpsWorks Stacks reduces the number of instances.
	DownScaling interface{} `field:"optional" json:"downScaling" yaml:"downScaling"`
	// Whether load-based auto scaling is enabled for the layer.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// An `AutoScalingThresholds` object that describes the upscaling configuration, which defines how and when AWS OpsWorks Stacks increases the number of instances.
	UpScaling interface{} `field:"optional" json:"upScaling" yaml:"upScaling"`
}

