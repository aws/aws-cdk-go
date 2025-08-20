package awsmedialive


// Complete these fields only if you want to insert watermarks of type Nielsen CBET.
//
// The parent of this entity is NielsenWatermarksSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenCBETProperty := &NielsenCBETProperty{
//   	CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   	CbetStepaside: jsii.String("cbetStepaside"),
//   	Csid: jsii.String("csid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsencbet.html
//
type CfnChannel_NielsenCBETProperty struct {
	// Enter the CBET check digits to use in the watermark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsencbet.html#cfn-medialive-channel-nielsencbet-cbetcheckdigitstring
	//
	CbetCheckDigitString *string `field:"optional" json:"cbetCheckDigitString" yaml:"cbetCheckDigitString"`
	// Determines the method of CBET insertion mode when prior encoding is detected on the same layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsencbet.html#cfn-medialive-channel-nielsencbet-cbetstepaside
	//
	CbetStepaside *string `field:"optional" json:"cbetStepaside" yaml:"cbetStepaside"`
	// Enter the CBET Source ID (CSID) to use in the watermark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsencbet.html#cfn-medialive-channel-nielsencbet-csid
	//
	Csid *string `field:"optional" json:"csid" yaml:"csid"`
}

