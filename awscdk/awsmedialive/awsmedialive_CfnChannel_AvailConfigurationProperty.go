package awsmedialive


// The setup of ad avail handling in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availConfigurationProperty := &AvailConfigurationProperty{
//   	AvailSettings: &AvailSettingsProperty{
//   		Scte35SpliceInsert: &Scte35SpliceInsertProperty{
//   			AdAvailOffset: jsii.Number(123),
//   			NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   		Scte35TimeSignalApos: &Scte35TimeSignalAposProperty{
//   			AdAvailOffset: jsii.Number(123),
//   			NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   	},
//   }
//
type CfnChannel_AvailConfigurationProperty struct {
	// The setup of ad avail handling in the output.
	AvailSettings interface{} `field:"optional" json:"availSettings" yaml:"availSettings"`
}

