package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prometheusProperty := &PrometheusProperty{
//   	JmxExporter: &JmxExporterProperty{
//   		EnabledInBroker: jsii.Boolean(false),
//   	},
//   	NodeExporter: &NodeExporterProperty{
//   		EnabledInBroker: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-prometheus.html
//
type CfnCluster_PrometheusProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-prometheus.html#cfn-msk-cluster-prometheus-jmxexporter
	//
	JmxExporter interface{} `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-prometheus.html#cfn-msk-cluster-prometheus-nodeexporter
	//
	NodeExporter interface{} `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

