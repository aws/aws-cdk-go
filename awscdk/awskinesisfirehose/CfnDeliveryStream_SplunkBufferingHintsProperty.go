package awskinesisfirehose


// The buffering options.
//
// If no value is specified, the default values for Splunk are used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splunkBufferingHintsProperty := &SplunkBufferingHintsProperty{
//   	IntervalInSeconds: jsii.Number(123),
//   	SizeInMBs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkbufferinghints.html
//
type CfnDeliveryStream_SplunkBufferingHintsProperty struct {
	// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination.
	//
	// The default value is 60 (1 minute).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkbufferinghints.html#cfn-kinesisfirehose-deliverystream-splunkbufferinghints-intervalinseconds
	//
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Buffer incoming data to the specified size, in MBs, before delivering it to the destination.
	//
	// The default value is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkbufferinghints.html#cfn-kinesisfirehose-deliverystream-splunkbufferinghints-sizeinmbs
	//
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

