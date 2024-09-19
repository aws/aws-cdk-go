package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openMonitoringProperty := &OpenMonitoringProperty{
//   	Prometheus: &PrometheusProperty{
//   		JmxExporter: &JmxExporterProperty{
//   			EnabledInBroker: jsii.Boolean(false),
//   		},
//   		NodeExporter: &NodeExporterProperty{
//   			EnabledInBroker: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-openmonitoring.html
//
type CfnCluster_OpenMonitoringProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-openmonitoring.html#cfn-msk-cluster-openmonitoring-prometheus
	//
	Prometheus interface{} `field:"required" json:"prometheus" yaml:"prometheus"`
}

