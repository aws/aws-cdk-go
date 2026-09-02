package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenNwOnlyProperty := &NielsenNwOnlyProperty{
//   	CheckDigitString: jsii.String("checkDigitString"),
//   	Sid: jsii.Number(123),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennwonly.html
//
type CfnChannel_NielsenNwOnlyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennwonly.html#cfn-medialive-channel-nielsennwonly-checkdigitstring
	//
	CheckDigitString *string `field:"optional" json:"checkDigitString" yaml:"checkDigitString"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennwonly.html#cfn-medialive-channel-nielsennwonly-sid
	//
	Sid *float64 `field:"optional" json:"sid" yaml:"sid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsennwonly.html#cfn-medialive-channel-nielsennwonly-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

