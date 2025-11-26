package previewawsmskmixins


// Prometheus settings for open monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnClusterPropsMixin_PrometheusProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-prometheus.html#cfn-msk-cluster-prometheus-jmxexporter
	//
	JmxExporter interface{} `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// Indicates whether you want to enable or disable the Node Exporter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-prometheus.html#cfn-msk-cluster-prometheus-nodeexporter
	//
	NodeExporter interface{} `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

