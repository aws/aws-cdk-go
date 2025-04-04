package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeliveryStream`.
//
// Example:
//   destinationBucket := s3.NewBucket(this, jsii.String("Bucket"))
//   deliveryStreamRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//
//   stream := firehose.NewCfnDeliveryStream(this, jsii.String("MyStream"), &CfnDeliveryStreamProps{
//   	DeliveryStreamName: jsii.String("amazon-apigateway-delivery-stream"),
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		BucketArn: destinationBucket.BucketArn,
//   		RoleArn: deliveryStreamRole.RoleArn,
//   	},
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: apigateway.NewFirehoseLogDestination(stream),
//   		AccessLogFormat: apigateway.AccessLogFormat_JsonWithStandardFields(),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
//
type CfnDeliveryStreamProps struct {
	// Describes the configuration of a destination in the Serverless offering for Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration
	//
	AmazonOpenSearchServerlessDestinationConfiguration interface{} `field:"optional" json:"amazonOpenSearchServerlessDestinationConfiguration" yaml:"amazonOpenSearchServerlessDestinationConfiguration"`
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-amazonopensearchservicedestinationconfiguration
	//
	AmazonopensearchserviceDestinationConfiguration interface{} `field:"optional" json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	// The top level object for configuring streams with database as a source.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration
	//
	DatabaseSourceConfiguration interface{} `field:"optional" json:"databaseSourceConfiguration" yaml:"databaseSourceConfiguration"`
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput
	//
	DeliveryStreamEncryptionConfigurationInput interface{} `field:"optional" json:"deliveryStreamEncryptionConfigurationInput" yaml:"deliveryStreamEncryptionConfigurationInput"`
	// The name of the Firehose stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamname
	//
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The Firehose stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the Firehose stream directly.
	// - `KinesisStreamAsSource` : The Firehose stream uses a Kinesis data stream as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamtype
	//
	DeliveryStreamType *string `field:"optional" json:"deliveryStreamType" yaml:"deliveryStreamType"`
	// The structure that configures parameters such as `ThroughputHintInMBs` for a stream configured with Direct PUT as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-directputsourceconfiguration
	//
	DirectPutSourceConfiguration interface{} `field:"optional" json:"directPutSourceConfiguration" yaml:"directPutSourceConfiguration"`
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration
	//
	ElasticsearchDestinationConfiguration interface{} `field:"optional" json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration
	//
	ExtendedS3DestinationConfiguration interface{} `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration
	//
	HttpEndpointDestinationConfiguration interface{} `field:"optional" json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	// Specifies the destination configure settings for Apache Iceberg Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-icebergdestinationconfiguration
	//
	IcebergDestinationConfiguration interface{} `field:"optional" json:"icebergDestinationConfiguration" yaml:"icebergDestinationConfiguration"`
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration
	//
	KinesisStreamSourceConfiguration interface{} `field:"optional" json:"kinesisStreamSourceConfiguration" yaml:"kinesisStreamSourceConfiguration"`
	// The configuration for the Amazon MSK cluster to be used as the source for a delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration
	//
	MskSourceConfiguration interface{} `field:"optional" json:"mskSourceConfiguration" yaml:"mskSourceConfiguration"`
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
	//
	RedshiftDestinationConfiguration interface{} `field:"optional" json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
	//
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Configure Snowflake destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration
	//
	SnowflakeDestinationConfiguration interface{} `field:"optional" json:"snowflakeDestinationConfiguration" yaml:"snowflakeDestinationConfiguration"`
	// The configuration of a destination in Splunk for the delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-splunkdestinationconfiguration
	//
	SplunkDestinationConfiguration interface{} `field:"optional" json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
	// A set of tags to assign to the Firehose stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the Firehose stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a Firehose stream.
	//
	// If you specify tags in the `CreateDeliveryStream` action, Amazon Data Firehose performs an additional authorization on the `firehose:TagDeliveryStream` action to verify if users have permissions to create tags. If you do not provide this permission, requests to create new Firehose streams with IAM resource tags will fail with an `AccessDeniedException` such as following.
	//
	// *AccessDeniedException*
	//
	// User: arn:aws:sts::x:assumed-role/x/x is not authorized to perform: firehose:TagDeliveryStream on resource: arn:aws:firehose:us-east-1:x:deliverystream/x with an explicit deny in an identity-based policy.
	//
	// For an example IAM policy, see [Tag example.](https://docs.aws.amazon.com/firehose/latest/APIReference/API_CreateDeliveryStream.html#API_CreateDeliveryStream_Examples)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

