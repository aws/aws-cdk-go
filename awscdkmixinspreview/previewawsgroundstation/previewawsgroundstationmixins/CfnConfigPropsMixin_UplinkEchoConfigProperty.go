package previewawsgroundstationmixins


// Provides information about how AWS Ground Station should echo back uplink transmissions to a dataflow endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uplinkEchoConfigProperty := &UplinkEchoConfigProperty{
//   	AntennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html
//
type CfnConfigPropsMixin_UplinkEchoConfigProperty struct {
	// Defines the ARN of the uplink config to echo back to a dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-antennauplinkconfigarn
	//
	AntennaUplinkConfigArn *string `field:"optional" json:"antennaUplinkConfigArn" yaml:"antennaUplinkConfigArn"`
	// Whether or not uplink echo is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

