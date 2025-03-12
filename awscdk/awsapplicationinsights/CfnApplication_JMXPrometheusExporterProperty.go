package awsapplicationinsights


// The `AWS::ApplicationInsights::Application JMXPrometheusExporter` property type defines the JMXPrometheus Exporter configuration.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jMXPrometheusExporterProperty := &JMXPrometheusExporterProperty{
//   	HostPort: jsii.String("hostPort"),
//   	Jmxurl: jsii.String("jmxurl"),
//   	PrometheusPort: jsii.String("prometheusPort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-jmxprometheusexporter.html
//
type CfnApplication_JMXPrometheusExporterProperty struct {
	// The host and port to connect to through remote JMX.
	//
	// Only one of `jmxURL` and `hostPort` can be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-jmxprometheusexporter.html#cfn-applicationinsights-application-jmxprometheusexporter-hostport
	//
	HostPort *string `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The complete JMX URL to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-jmxprometheusexporter.html#cfn-applicationinsights-application-jmxprometheusexporter-jmxurl
	//
	Jmxurl *string `field:"optional" json:"jmxurl" yaml:"jmxurl"`
	// The target port to send Prometheus metrics to.
	//
	// If not specified, the default port `9404` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-jmxprometheusexporter.html#cfn-applicationinsights-application-jmxprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

