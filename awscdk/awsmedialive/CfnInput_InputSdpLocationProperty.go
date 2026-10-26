package awsmedialive


// The location of the SDP file for one of the SMPTE 2110 streams in a receiver group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSdpLocationProperty := &InputSdpLocationProperty{
//   	MediaIndex: jsii.Number(123),
//   	SdpUrl: jsii.String("sdpUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsdplocation.html
//
type CfnInput_InputSdpLocationProperty struct {
	// The index of the media stream in the SDP file for one SMPTE 2110 stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsdplocation.html#cfn-medialive-input-inputsdplocation-mediaindex
	//
	MediaIndex *float64 `field:"optional" json:"mediaIndex" yaml:"mediaIndex"`
	// The URL of the SDP file for one SMPTE 2110 stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsdplocation.html#cfn-medialive-input-inputsdplocation-sdpurl
	//
	SdpUrl *string `field:"optional" json:"sdpUrl" yaml:"sdpUrl"`
}

