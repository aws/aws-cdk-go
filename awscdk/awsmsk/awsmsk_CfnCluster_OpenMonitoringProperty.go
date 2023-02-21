package awsmsk


// JMX and Node monitoring for the MSK cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openMonitoringProperty := &openMonitoringProperty{
//   	prometheus: &prometheusProperty{
//   		jmxExporter: &jmxExporterProperty{
//   			enabledInBroker: jsii.Boolean(false),
//   		},
//   		nodeExporter: &nodeExporterProperty{
//   			enabledInBroker: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCluster_OpenMonitoringProperty struct {
	// Prometheus exporter settings.
	Prometheus interface{} `field:"required" json:"prometheus" yaml:"prometheus"`
}

