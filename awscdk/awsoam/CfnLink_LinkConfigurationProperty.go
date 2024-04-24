package awsoam


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkconfiguration.html#cfn-oam-link-linkconfiguration-loggroupconfiguration
	//
	LogGroupConfiguration interface{} `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-oam-link-linkconfiguration.html#cfn-oam-link-linkconfiguration-metricconfiguration
	//
	MetricConfiguration interface{} `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

