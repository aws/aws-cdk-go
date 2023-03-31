package awskinesisfirehose


// The `SplunkRetryOptions` property type specifies retry behavior in case Kinesis Data Firehose is unable to deliver documents to Splunk or if it doesn't receive an acknowledgment from Splunk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splunkRetryOptionsProperty := &splunkRetryOptionsProperty{
//   	durationInSeconds: jsii.Number(123),
//   }
//
type CfnDeliveryStream_SplunkRetryOptionsProperty struct {
	// The total amount of time that Kinesis Data Firehose spends on retries.
	//
	// This duration starts after the initial attempt to send data to Splunk fails. It doesn't include the periods during which Kinesis Data Firehose waits for acknowledgment from Splunk after each attempt.
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
}

