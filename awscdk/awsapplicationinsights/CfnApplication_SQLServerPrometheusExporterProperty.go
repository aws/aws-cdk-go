package awsapplicationinsights


// The SQL prometheus exporter settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sQLServerPrometheusExporterProperty := &SQLServerPrometheusExporterProperty{
//   	PrometheusPort: jsii.String("prometheusPort"),
//   	SqlSecretName: jsii.String("sqlSecretName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html
//
type CfnApplication_SQLServerPrometheusExporterProperty struct {
	// Prometheus exporter port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html#cfn-applicationinsights-application-sqlserverprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"required" json:"prometheusPort" yaml:"prometheusPort"`
	// Secret name which managers SQL exporter connection.
	//
	// e.g. {"data_source_name": "sqlserver://<USERNAME>:<PASSWORD>@localhost:1433"}
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html#cfn-applicationinsights-application-sqlserverprometheusexporter-sqlsecretname
	//
	SqlSecretName *string `field:"required" json:"sqlSecretName" yaml:"sqlSecretName"`
}

