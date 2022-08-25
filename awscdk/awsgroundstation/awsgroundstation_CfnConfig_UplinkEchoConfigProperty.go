package awsgroundstation


// Provides information about how AWS Ground Station should echo back uplink transmissions to a dataflow endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkEchoConfigProperty := &uplinkEchoConfigProperty{
//   	antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnConfig_UplinkEchoConfigProperty struct {
	// Defines the ARN of the uplink config to echo back to a dataflow endpoint.
	AntennaUplinkConfigArn *string `field:"optional" json:"antennaUplinkConfigArn" yaml:"antennaUplinkConfigArn"`
	// Whether or not uplink echo is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

