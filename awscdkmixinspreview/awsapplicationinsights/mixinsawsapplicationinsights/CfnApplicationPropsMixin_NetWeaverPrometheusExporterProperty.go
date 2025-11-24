package mixinsawsapplicationinsights


// The NetWeaver Prometheus Exporter Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   netWeaverPrometheusExporterProperty := &NetWeaverPrometheusExporterProperty{
//   	InstanceNumbers: []*string{
//   		jsii.String("instanceNumbers"),
//   	},
//   	PrometheusPort: jsii.String("prometheusPort"),
//   	Sapsid: jsii.String("sapsid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html
//
type CfnApplicationPropsMixin_NetWeaverPrometheusExporterProperty struct {
	// SAP instance numbers for ASCS, ERS, and App Servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-instancenumbers
	//
	InstanceNumbers *[]*string `field:"optional" json:"instanceNumbers" yaml:"instanceNumbers"`
	// Prometheus exporter port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-prometheusport
	//
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
	// SAP NetWeaver SID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-sapsid
	//
	Sapsid *string `field:"optional" json:"sapsid" yaml:"sapsid"`
}

