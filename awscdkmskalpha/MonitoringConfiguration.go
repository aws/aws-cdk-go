package awscdkmskalpha


// Monitoring Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   monitoringConfiguration := &MonitoringConfiguration{
//   	ClusterMonitoringLevel: msk_alpha.ClusterMonitoringLevel_DEFAULT,
//   	EnablePrometheusJmxExporter: jsii.Boolean(false),
//   	EnablePrometheusNodeExporter: jsii.Boolean(false),
//   }
//
// Experimental.
type MonitoringConfiguration struct {
	// Specifies the level of monitoring for the MSK cluster.
	// Default: DEFAULT.
	//
	// Experimental.
	ClusterMonitoringLevel ClusterMonitoringLevel `field:"optional" json:"clusterMonitoringLevel" yaml:"clusterMonitoringLevel"`
	// Indicates whether you want to enable or disable the JMX Exporter.
	// Default: false.
	//
	// Experimental.
	EnablePrometheusJmxExporter *bool `field:"optional" json:"enablePrometheusJmxExporter" yaml:"enablePrometheusJmxExporter"`
	// Indicates whether you want to enable or disable the Prometheus Node Exporter.
	//
	// You can use the Prometheus Node Exporter to get CPU and disk metrics for the broker nodes.
	// Default: false.
	//
	// Experimental.
	EnablePrometheusNodeExporter *bool `field:"optional" json:"enablePrometheusNodeExporter" yaml:"enablePrometheusNodeExporter"`
}

