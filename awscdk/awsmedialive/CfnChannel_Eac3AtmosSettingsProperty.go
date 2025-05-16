package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eac3AtmosSettingsProperty := &Eac3AtmosSettingsProperty{
//   	Bitrate: jsii.Number(123),
//   	CodingMode: jsii.String("codingMode"),
//   	Dialnorm: jsii.Number(123),
//   	DrcLine: jsii.String("drcLine"),
//   	DrcRf: jsii.String("drcRf"),
//   	HeightTrim: jsii.Number(123),
//   	SurroundTrim: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html
//
type CfnChannel_Eac3AtmosSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-codingmode
	//
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-dialnorm
	//
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-drcline
	//
	DrcLine *string `field:"optional" json:"drcLine" yaml:"drcLine"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-drcrf
	//
	DrcRf *string `field:"optional" json:"drcRf" yaml:"drcRf"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-heighttrim
	//
	HeightTrim *float64 `field:"optional" json:"heightTrim" yaml:"heightTrim"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-surroundtrim
	//
	SurroundTrim *float64 `field:"optional" json:"surroundTrim" yaml:"surroundTrim"`
}

