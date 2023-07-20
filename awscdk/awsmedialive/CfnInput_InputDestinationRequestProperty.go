package awsmedialive


// Settings that apply only if the input is a push type of input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDestinationRequestProperty := &InputDestinationRequestProperty{
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html
//
type CfnInput_InputDestinationRequestProperty struct {
	// The stream name (application name/application instance) for the location the RTMP source content will be pushed to in MediaLive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationrequest.html#cfn-medialive-input-inputdestinationrequest-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

