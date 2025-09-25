package awsapplicationinsights


// The NetWeaver Prometheus Exporter Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   netWeaverPrometheusExporterProperty := &NetWeaverPrometheusExporterProperty{
//   	InstanceNumbers: []*string{
//   		jsii.String("instanceNumbers"),
//   	},
//   	Sapsid: jsii.String("sapsid"),
//
//   	// the properties below are optional
//   	PrometheusPort: jsii.String("prometheusPort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html
//
type CfnApplication_NetWeaverPrometheusExporterProperty struct {
	// SAP instance numbers for ASCS, ERS, and App Servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-instancenumbers
	//
	InstanceNumbers *[]*string `field:"required" json:"instanceNumbers" yaml:"instanceNumbers"`
	// SAP NetWeaver SID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-sapsid
	//
	Sapsid *string `field:"required" json:"sapsid" yaml:"sapsid"`
	// Prometheus exporter port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

