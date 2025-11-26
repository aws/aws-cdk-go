package previewawsmskmixins


// You can configure your MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingInfoProperty := &LoggingInfoProperty{
//   	BrokerLogs: &BrokerLogsProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-logginginfo.html
//
type CfnClusterPropsMixin_LoggingInfoProperty struct {
	// You can configure your MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-logginginfo.html#cfn-msk-cluster-logginginfo-brokerlogs
	//
	BrokerLogs interface{} `field:"optional" json:"brokerLogs" yaml:"brokerLogs"`
}

