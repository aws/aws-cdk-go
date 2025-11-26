package previewawsapplicationinsightsmixins


// The `AWS::ApplicationInsights::Application HAClusterPrometheusExporter` property type defines the HA cluster Prometheus Exporter settings.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hAClusterPrometheusExporterProperty := &HAClusterPrometheusExporterProperty{
//   	PrometheusPort: jsii.String("prometheusPort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-haclusterprometheusexporter.html
//
type CfnApplicationPropsMixin_HAClusterPrometheusExporterProperty struct {
	// The target port to which Prometheus sends metrics.
	//
	// If not specified, the default port 9668 is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-haclusterprometheusexporter.html#cfn-applicationinsights-application-haclusterprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

