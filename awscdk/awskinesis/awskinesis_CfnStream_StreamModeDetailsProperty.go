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
//   streamModeDetailsProperty := &streamModeDetailsProperty{
//   	streamMode: jsii.String("streamMode"),
//   }
//
type CfnStream_StreamModeDetailsProperty struct {
	// Specifies the capacity mode to which you want to set your data stream.
	//
	// Currently, in Kinesis Data Streams, you can choose between an *on-demand* capacity mode and a *provisioned* capacity mode for your data streams.
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

