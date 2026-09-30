package awskinesis


// The configuration that describes how records on the source stream are encoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recordConfigurationProperty := &RecordConfigurationProperty{
//   	GsrSchemaArn: jsii.String("gsrSchemaArn"),
//   	RecordFormatType: jsii.String("recordFormatType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-recordconfiguration.html
//
type CfnChannelPropsMixin_RecordConfigurationProperty struct {
	// The ARN of the AWS Glue Schema Registry (GSR) schema.
	//
	// Required for the S3 Tables destination, where it is used to create the S3 Table and to validate that the record format matches the table schema. Also used when RecordFormatType is GSR_JSON to interpret records read from the source stream. Vanilla S3 delivery writes records as S3 objects and does not need a schema. The schema must be in the same account and region as the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-recordconfiguration.html#cfn-kinesis-channel-recordconfiguration-gsrschemaarn
	//
	GsrSchemaArn *string `field:"optional" json:"gsrSchemaArn" yaml:"gsrSchemaArn"`
	// The format used to interpret records read from the source stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-recordconfiguration.html#cfn-kinesis-channel-recordconfiguration-recordformattype
	//
	RecordFormatType *string `field:"optional" json:"recordFormatType" yaml:"recordFormatType"`
}

