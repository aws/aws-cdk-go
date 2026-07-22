package awsmedialive


// SDI source mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sdiSourceMappingProperty := &SdiSourceMappingProperty{
//   	CardNumber: jsii.Number(123),
//   	ChannelNumber: jsii.Number(123),
//   	SdiSource: jsii.String("sdiSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-sdisourcemapping.html
//
type CfnNode_SdiSourceMappingProperty struct {
	// The card number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-sdisourcemapping.html#cfn-medialive-node-sdisourcemapping-cardnumber
	//
	CardNumber *float64 `field:"optional" json:"cardNumber" yaml:"cardNumber"`
	// The channel number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-sdisourcemapping.html#cfn-medialive-node-sdisourcemapping-channelnumber
	//
	ChannelNumber *float64 `field:"optional" json:"channelNumber" yaml:"channelNumber"`
	// The SDI source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-node-sdisourcemapping.html#cfn-medialive-node-sdisourcemapping-sdisource
	//
	SdiSource *string `field:"optional" json:"sdiSource" yaml:"sdiSource"`
}

