package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for inserting a DVB Network Information Table (NIT).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbNitSettings := &DvbNitSettings{
//   	NetworkId: jsii.Number(123),
//   	NetworkName: jsii.String("networkName"),
//   	RepInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type DvbNitSettings struct {
	// The numeric value placed in the Network Information Table (NIT).
	// Default: - no network ID.
	//
	// Experimental.
	NetworkId *float64 `field:"optional" json:"networkId" yaml:"networkId"`
	// The network name placed in the networkNameDescriptor inside the NIT (max 256 characters).
	// Default: - no network name.
	//
	// Experimental.
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The interval between instances of this table in the output transport stream.
	// Default: - service default.
	//
	// Experimental.
	RepInterval awscdk.Duration `field:"optional" json:"repInterval" yaml:"repInterval"`
}

