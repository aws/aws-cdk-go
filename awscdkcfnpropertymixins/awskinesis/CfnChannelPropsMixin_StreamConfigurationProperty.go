package awskinesis


// The configuration of a Kinesis stream that the channel reads from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   streamConfigurationProperty := &StreamConfigurationProperty{
//   	RecordConfiguration: &RecordConfigurationProperty{
//   		GsrSchemaArn: jsii.String("gsrSchemaArn"),
//   		RecordFormatType: jsii.String("recordFormatType"),
//   	},
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-streamconfiguration.html
//
type CfnChannelPropsMixin_StreamConfigurationProperty struct {
	// The configuration that describes how records on the source stream are encoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-streamconfiguration.html#cfn-kinesis-channel-streamconfiguration-recordconfiguration
	//
	RecordConfiguration interface{} `field:"optional" json:"recordConfiguration" yaml:"recordConfiguration"`
	// The Amazon resource name (ARN) of the Kinesis data stream that the channel reads from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-streamconfiguration.html#cfn-kinesis-channel-streamconfiguration-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

