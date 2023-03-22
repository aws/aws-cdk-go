package awsapplicationinsights


// The `AWS::ApplicationInsights::Application HAClusterPrometheusExporter` property type defines the HA cluster Prometheus Exporter settings.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hAClusterPrometheusExporterProperty := &hAClusterPrometheusExporterProperty{
//   	prometheusPort: jsii.String("prometheusPort"),
//   }
//
type CfnApplication_HAClusterPrometheusExporterProperty struct {
	// The target port to which Prometheus sends metrics.
	//
	// If not specified, the default port 9668 is used.
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

