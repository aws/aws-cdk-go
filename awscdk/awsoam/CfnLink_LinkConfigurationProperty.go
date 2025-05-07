package awsoam


// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkConfigurationProperty := &LinkConfigurationProperty{
//   	LogGroupConfiguration: &LinkFilterProperty{
//   		Filter: jsii.String("filter"),
//   	},
//   	MetricConfiguration: &LinkFilterProperty{
//   		Filter: jsii.String("filter"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkconfiguration.html
//
type CfnLink_LinkConfigurationProperty struct {
	// Use this structure to filter which log groups are to share log events from this source account to the monitoring account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkconfiguration.html#cfn-oam-link-linkconfiguration-loggroupconfiguration
	//
	LogGroupConfiguration interface{} `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// Use this structure to filter which metric namespaces are to be shared from the source account to the monitoring account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkconfiguration.html#cfn-oam-link-linkconfiguration-metricconfiguration
	//
	MetricConfiguration interface{} `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

