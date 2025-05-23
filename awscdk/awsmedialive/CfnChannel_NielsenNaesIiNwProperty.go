package awsmedialive


// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
//
// The parent of this entity is NielsenWatermarksSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenNaesIiNwProperty := &NielsenNaesIiNwProperty{
//   	CheckDigitString: jsii.String("checkDigitString"),
//   	Sid: jsii.Number(123),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennaesiinw.html
//
type CfnChannel_NielsenNaesIiNwProperty struct {
	// Enter the check digit string for the watermark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennaesiinw.html#cfn-medialive-channel-nielsennaesiinw-checkdigitstring
	//
	CheckDigitString *string `field:"optional" json:"checkDigitString" yaml:"checkDigitString"`
	// Enter the Nielsen Source ID (SID) to include in the watermark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennaesiinw.html#cfn-medialive-channel-nielsennaesiinw-sid
	//
	Sid *float64 `field:"optional" json:"sid" yaml:"sid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennaesiinw.html#cfn-medialive-channel-nielsennaesiinw-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

