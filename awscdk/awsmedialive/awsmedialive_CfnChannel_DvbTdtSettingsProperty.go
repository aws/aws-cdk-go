package awsmedialive


// The DVB Time and Date Table (TDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbTdtSettingsProperty := &dvbTdtSettingsProperty{
//   	repInterval: jsii.Number(123),
//   }
//
type CfnChannel_DvbTdtSettingsProperty struct {
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

