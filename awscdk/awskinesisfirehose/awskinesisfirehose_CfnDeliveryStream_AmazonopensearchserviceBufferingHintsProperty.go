package awskinesisfirehose


// Describes the buffering to perform before delivering data to the Amazon OpenSearch Service destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonopensearchserviceBufferingHintsProperty := &amazonopensearchserviceBufferingHintsProperty{
//   	intervalInSeconds: jsii.Number(123),
//   	sizeInMBs: jsii.Number(123),
//   }
//
type CfnDeliveryStream_AmazonopensearchserviceBufferingHintsProperty struct {
	// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination.
	//
	// The default value is 300 (5 minutes).
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Buffer incoming data to the specified size, in MBs, before delivering it to the destination.
	//
	// The default value is 5. We recommend setting this parameter to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec, the value should be 10 MB or higher.
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

