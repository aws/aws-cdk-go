package awskinesisfirehose


// Configures retry behavior in case Firehose is unable to deliver documents to Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftRetryOptionsProperty := &RedshiftRetryOptionsProperty{
//   	DurationInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-redshiftretryoptions.html
//
type CfnDeliveryStream_RedshiftRetryOptionsProperty struct {
	// The length of time during which Firehose retries delivery after a failure, starting from the initial request and including the first attempt.
	//
	// The default value is 3600 seconds (60 minutes). Firehose does not retry if the value of `DurationInSeconds` is 0 (zero) or if the first delivery attempt takes longer than the current value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-redshiftretryoptions.html#cfn-kinesisfirehose-deliverystream-redshiftretryoptions-durationinseconds
	//
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
}

