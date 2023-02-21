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
//   availConfigurationProperty := &availConfigurationProperty{
//   	availSettings: &availSettingsProperty{
//   		scte35SpliceInsert: &scte35SpliceInsertProperty{
//   			adAvailOffset: jsii.Number(123),
//   			noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   		scte35TimeSignalApos: &scte35TimeSignalAposProperty{
//   			adAvailOffset: jsii.Number(123),
//   			noRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			webDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   	},
//   }
//
type CfnChannel_AvailConfigurationProperty struct {
	// The setup of ad avail handling in the output.
	AvailSettings interface{} `field:"optional" json:"availSettings" yaml:"availSettings"`
}

