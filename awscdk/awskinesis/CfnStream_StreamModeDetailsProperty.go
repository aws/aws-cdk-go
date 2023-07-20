package awskinesis


// Specifies the capacity mode to which you want to set your data stream.
//
// Currently, in Kinesis Data Streams, you can choose between an *on-demand* capacity mode and a *provisioned* capacity mode for your data streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamModeDetailsProperty := &StreamModeDetailsProperty{
//   	StreamMode: jsii.String("streamMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streammodedetails.html
//
type CfnStream_StreamModeDetailsProperty struct {
	// Specifies the capacity mode to which you want to set your data stream.
	//
	// Currently, in Kinesis Data Streams, you can choose between an *on-demand* capacity mode and a *provisioned* capacity mode for your data streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streammodedetails.html#cfn-kinesis-stream-streammodedetails-streammode
	//
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

