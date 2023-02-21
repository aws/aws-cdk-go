package awsmedialive


// The settings for the ad avail setup in the output.
//
// The parent of this entity is AvailConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSettingsProperty := &availSettingsProperty{
//   	scte35SpliceInsert: &scte35SpliceInsertProperty{
//   		adAvailOffset: jsii.Number(123),
//   		noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   	scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   		adAvailOffset: jsii.Number(123),
//   		noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   }
//
type CfnChannel_AvailSettingsProperty struct {
	// The setup for SCTE-35 splice insert handling.
	Scte35SpliceInsert interface{} `field:"optional" json:"scte35SpliceInsert" yaml:"scte35SpliceInsert"`
	// The setup for SCTE-35 time signal APOS handling.
	Scte35TimeSignalApos interface{} `field:"optional" json:"scte35TimeSignalApos" yaml:"scte35TimeSignalApos"`
}

