package awsecs


// The configuration for a specific set of metrics to collect for a service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricConfigurationProperty := &MetricConfigurationProperty{
//   	MetricNames: []*string{
//   		jsii.String("metricNames"),
//   	},
//   	ResolutionSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-metricconfiguration.html
//
type CfnServicePropsMixin_MetricConfigurationProperty struct {
	// The list of metric names to configure.
	//
	// The supported metric names are ``CPUUtilization`` and ``MemoryUtilization``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-metricconfiguration.html#cfn-ecs-service-metricconfiguration-metricnames
	//
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// The resolution, in seconds, at which to collect the metrics.
	//
	// The valid values are ``20`` and ``60``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-metricconfiguration.html#cfn-ecs-service-metricconfiguration-resolutionseconds
	//
	ResolutionSeconds *float64 `field:"optional" json:"resolutionSeconds" yaml:"resolutionSeconds"`
}

