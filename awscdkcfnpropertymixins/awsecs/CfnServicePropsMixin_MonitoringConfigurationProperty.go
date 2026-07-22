package awsecs


// The optional monitoring configuration for a service, which defines the resolution for the service-level ``CPUUtilization`` and ``MemoryUtilization`` Amazon CloudWatch metrics.
//
// When not specified, Amazon ECS uses the default resolution of ``60`` seconds.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   monitoringConfigurationProperty := &MonitoringConfigurationProperty{
//   	MetricConfigurations: []interface{}{
//   		&MetricConfigurationProperty{
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   			ResolutionSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-monitoringconfiguration.html
//
type CfnServicePropsMixin_MonitoringConfigurationProperty struct {
	// The list of metric configurations for the service monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-monitoringconfiguration.html#cfn-ecs-service-monitoringconfiguration-metricconfigurations
	//
	MetricConfigurations interface{} `field:"optional" json:"metricConfigurations" yaml:"metricConfigurations"`
}

