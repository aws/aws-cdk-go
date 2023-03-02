package awskinesisfirehose


// Configures retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon OpenSearch Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonopensearchserviceRetryOptionsProperty := &amazonopensearchserviceRetryOptionsProperty{
//   	durationInSeconds: jsii.Number(123),
//   }
//
type CfnDeliveryStream_AmazonopensearchserviceRetryOptionsProperty struct {
	// After an initial failure to deliver to Amazon OpenSearch Service, the total amount of time during which Kinesis Data Firehose retries delivery (including the first attempt).
	//
	// After this time has elapsed, the failed documents are written to Amazon S3. Default value is 300 seconds (5 minutes). A value of 0 (zero) results in no retries.
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
}

