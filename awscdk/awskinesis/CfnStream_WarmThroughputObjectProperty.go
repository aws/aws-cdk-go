package awskinesis


// Represents the warm throughput configuration on the stream.
//
// This is only present for On-Demand Kinesis Data Streams in accounts that have `MinimumThroughputBillingCommitment` enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   warmThroughputObjectProperty := &WarmThroughputObjectProperty{
//   	CurrentMiBps: jsii.Number(123),
//   	TargetMiBps: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-warmthroughputobject.html
//
type CfnStream_WarmThroughputObjectProperty struct {
	// The current warm throughput value on the stream.
	//
	// This is the write throughput in MiBps that the stream is currently scaled to handle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-warmthroughputobject.html#cfn-kinesis-stream-warmthroughputobject-currentmibps
	//
	CurrentMiBps *float64 `field:"optional" json:"currentMiBps" yaml:"currentMiBps"`
	// The target warm throughput value on the stream.
	//
	// This indicates that the stream is currently scaling towards this target value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-warmthroughputobject.html#cfn-kinesis-stream-warmthroughputobject-targetmibps
	//
	TargetMiBps *float64 `field:"optional" json:"targetMiBps" yaml:"targetMiBps"`
}

