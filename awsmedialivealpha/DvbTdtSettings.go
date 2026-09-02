package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for inserting a DVB Time and Date Table (TDT).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbTdtSettings := &DvbTdtSettings{
//   	RepInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type DvbTdtSettings struct {
	// The interval between instances of this table in the output transport stream.
	// Default: - service default.
	//
	// Experimental.
	RepInterval awscdk.Duration `field:"optional" json:"repInterval" yaml:"repInterval"`
}

