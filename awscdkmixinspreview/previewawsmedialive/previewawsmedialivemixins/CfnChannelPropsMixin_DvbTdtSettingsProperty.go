package previewawsmedialivemixins


// The DVB Time and Date Table (TDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dvbTdtSettingsProperty := &DvbTdtSettingsProperty{
//   	RepInterval: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbtdtsettings.html
//
type CfnChannelPropsMixin_DvbTdtSettingsProperty struct {
	// The number of milliseconds between instances of this table in the output transport stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbtdtsettings.html#cfn-medialive-channel-dvbtdtsettings-repinterval
	//
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

