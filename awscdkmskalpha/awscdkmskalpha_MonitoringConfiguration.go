// The CDK Construct Library for AWS::MSK
package awscdkmskalpha


// Monitoring Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   monitoringConfiguration := &monitoringConfiguration{
//   	clusterMonitoringLevel: msk_alpha.clusterMonitoringLevel_DEFAULT,
//   	enablePrometheusJmxExporter: jsii.Boolean(false),
//   	enablePrometheusNodeExporter: jsii.Boolean(false),
//   }
//
// Experimental.
type MonitoringConfiguration struct {
	// Specifies the level of monitoring for the MSK cluster.
	// Experimental.
	ClusterMonitoringLevel ClusterMonitoringLevel `field:"optional" json:"clusterMonitoringLevel" yaml:"clusterMonitoringLevel"`
	// Indicates whether you want to enable or disable the JMX Exporter.
	// Experimental.
	EnablePrometheusJmxExporter *bool `field:"optional" json:"enablePrometheusJmxExporter" yaml:"enablePrometheusJmxExporter"`
	// Indicates whether you want to enable or disable the Prometheus Node Exporter.
	//
	// You can use the Prometheus Node Exporter to get CPU and disk metrics for the broker nodes.
	// Experimental.
	EnablePrometheusNodeExporter *bool `field:"optional" json:"enablePrometheusNodeExporter" yaml:"enablePrometheusNodeExporter"`
}

