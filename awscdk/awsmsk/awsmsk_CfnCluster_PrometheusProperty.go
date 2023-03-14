package awsmsk


// Prometheus settings for open monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prometheusProperty := &prometheusProperty{
//   	jmxExporter: &jmxExporterProperty{
//   		enabledInBroker: jsii.Boolean(false),
//   	},
//   	nodeExporter: &nodeExporterProperty{
//   		enabledInBroker: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_PrometheusProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter interface{} `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter interface{} `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

