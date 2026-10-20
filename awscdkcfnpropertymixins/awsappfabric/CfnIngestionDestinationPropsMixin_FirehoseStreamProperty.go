package awsappfabric


// Contains information about an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   firehoseStreamProperty := &FirehoseStreamProperty{
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-firehosestream.html
//
type CfnIngestionDestinationPropsMixin_FirehoseStreamProperty struct {
	// The name of the Amazon Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-firehosestream.html#cfn-appfabric-ingestiondestination-firehosestream-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

