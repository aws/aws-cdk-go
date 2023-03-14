package awsmedialive


// The configuration of DVB NIT.
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbNitSettingsProperty := &dvbNitSettingsProperty{
//   	networkId: jsii.Number(123),
//   	networkName: jsii.String("networkName"),
//   	repInterval: jsii.Number(123),
//   }
//
type CfnChannel_DvbNitSettingsProperty struct {
	// The numeric value placed in the Network Information Table (NIT).
	NetworkId *float64 `field:"optional" json:"networkId" yaml:"networkId"`
	// The network name text placed in the networkNameDescriptor inside the Network Information Table (NIT).
	//
	// The maximum length is 256 characters.
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

