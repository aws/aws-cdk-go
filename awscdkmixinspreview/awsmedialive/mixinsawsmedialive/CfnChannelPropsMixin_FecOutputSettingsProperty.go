package mixinsawsmedialive


// The settings for FEC.
//
// The parent of this entity is UdpOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fecOutputSettingsProperty := &FecOutputSettingsProperty{
//   	ColumnDepth: jsii.Number(123),
//   	IncludeFec: jsii.String("includeFec"),
//   	RowLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fecoutputsettings.html
//
type CfnChannelPropsMixin_FecOutputSettingsProperty struct {
	// The parameter D from SMPTE 2022-1.
	//
	// The height of the FEC protection matrix. The number of transport stream packets per column error correction packet. The number must be between 4 and 20, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fecoutputsettings.html#cfn-medialive-channel-fecoutputsettings-columndepth
	//
	ColumnDepth *float64 `field:"optional" json:"columnDepth" yaml:"columnDepth"`
	// Enables column only or column and row-based FEC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fecoutputsettings.html#cfn-medialive-channel-fecoutputsettings-includefec
	//
	IncludeFec *string `field:"optional" json:"includeFec" yaml:"includeFec"`
	// The parameter L from SMPTE 2022-1.
	//
	// The width of the FEC protection matrix. Must be between 1 and 20, inclusive. If only Column FEC is used, then larger values increase robustness. If Row FEC is used, then this is the number of transport stream packets per row error correction packet, and the value must be between 4 and 20, inclusive, if includeFec is columnAndRow. If includeFec is column, this value must be 1 to 20, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fecoutputsettings.html#cfn-medialive-channel-fecoutputsettings-rowlength
	//
	RowLength *float64 `field:"optional" json:"rowLength" yaml:"rowLength"`
}

