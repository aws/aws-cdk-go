package awsapplicationinsights


// The `AWS::ApplicationInsights::Application HANAPrometheusExporter` property type defines the HANA DB Prometheus Exporter settings.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hANAPrometheusExporterProperty := &HANAPrometheusExporterProperty{
//   	AgreeToInstallHanadbClient: jsii.Boolean(false),
//   	HanaPort: jsii.String("hanaPort"),
//   	HanaSecretName: jsii.String("hanaSecretName"),
//   	Hanasid: jsii.String("hanasid"),
//
//   	// the properties below are optional
//   	PrometheusPort: jsii.String("prometheusPort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html
//
type CfnApplication_HANAPrometheusExporterProperty struct {
	// Designates whether you agree to install the HANA DB client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-agreetoinstallhanadbclient
	//
	AgreeToInstallHanadbClient interface{} `field:"required" json:"agreeToInstallHanadbClient" yaml:"agreeToInstallHanadbClient"`
	// The HANA database port by which the exporter will query HANA metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanaport
	//
	HanaPort *string `field:"required" json:"hanaPort" yaml:"hanaPort"`
	// The AWS Secrets Manager secret that stores HANA monitoring user credentials.
	//
	// The HANA Prometheus exporter uses these credentials to connect to the database and query HANA metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanasecretname
	//
	HanaSecretName *string `field:"required" json:"hanaSecretName" yaml:"hanaSecretName"`
	// The three-character SAP system ID (SID) of the SAP HANA system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanasid
	//
	Hanasid *string `field:"required" json:"hanasid" yaml:"hanasid"`
	// The target port to which Prometheus sends metrics.
	//
	// If not specified, the default port 9668 is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

