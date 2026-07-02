package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-monitoringconfiguration.html#cfn-ecs-service-monitoringconfiguration-metricconfigurations
	//
	MetricConfigurations interface{} `field:"optional" json:"metricConfigurations" yaml:"metricConfigurations"`
}

