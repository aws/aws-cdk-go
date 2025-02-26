package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioDolbyEDecodeProperty := &AudioDolbyEDecodeProperty{
//   	ProgramSelection: jsii.String("programSelection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodolbyedecode.html
//
type CfnChannel_AudioDolbyEDecodeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodolbyedecode.html#cfn-medialive-channel-audiodolbyedecode-programselection
	//
	ProgramSelection *string `field:"optional" json:"programSelection" yaml:"programSelection"`
}

