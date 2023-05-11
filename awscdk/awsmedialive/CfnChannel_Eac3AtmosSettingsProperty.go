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
type CfnChannel_Eac3AtmosSettingsProperty struct {
	// `CfnChannel.Eac3AtmosSettingsProperty.Bitrate`.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// `CfnChannel.Eac3AtmosSettingsProperty.CodingMode`.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// `CfnChannel.Eac3AtmosSettingsProperty.Dialnorm`.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// `CfnChannel.Eac3AtmosSettingsProperty.DrcLine`.
	DrcLine *string `field:"optional" json:"drcLine" yaml:"drcLine"`
	// `CfnChannel.Eac3AtmosSettingsProperty.DrcRf`.
	DrcRf *string `field:"optional" json:"drcRf" yaml:"drcRf"`
	// `CfnChannel.Eac3AtmosSettingsProperty.HeightTrim`.
	HeightTrim *float64 `field:"optional" json:"heightTrim" yaml:"heightTrim"`
	// `CfnChannel.Eac3AtmosSettingsProperty.SurroundTrim`.
	SurroundTrim *float64 `field:"optional" json:"surroundTrim" yaml:"surroundTrim"`
}

