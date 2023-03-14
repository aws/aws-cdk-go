package awsnetworkfirewall


// Defines where AWS Network Firewall sends logs for the firewall for one log type.
//
// This is used in `LoggingConfiguration` . You can send each type of log to an Amazon S3 bucket, a CloudWatch log group, or a Kinesis Data Firehose delivery stream.
//
// Network Firewall generates logs for stateful rule groups. You can save alert and flow log types. The stateful rules engine records flow logs for all network traffic that it receives. It records alert logs for traffic that matches stateful rules that have the rule action set to `DROP` or `ALERT` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDestinationConfigProperty := &logDestinationConfigProperty{
//   	logDestination: map[string]*string{
//   		"logDestinationKey": jsii.String("logDestination"),
//   	},
//   	logDestinationType: jsii.String("logDestinationType"),
//   	logType: jsii.String("logType"),
//   }
//
type CfnLoggingConfiguration_LogDestinationConfigProperty struct {
	// The named location for the logs, provided in a key:value mapping that is specific to the chosen destination type.
	//
	// - For an Amazon S3 bucket, provide the name of the bucket, with key `bucketName` , and optionally provide a prefix, with key `prefix` . The following example specifies an Amazon S3 bucket named `DOC-EXAMPLE-BUCKET` and the prefix `alerts` :
	//
	// `"LogDestination": { "bucketName": "DOC-EXAMPLE-BUCKET", "prefix": "alerts" }`
	// - For a CloudWatch log group, provide the name of the CloudWatch log group, with key `logGroup` . The following example specifies a log group named `alert-log-group` :
	//
	// `"LogDestination": { "logGroup": "alert-log-group" }`
	// - For a Kinesis Data Firehose delivery stream, provide the name of the delivery stream, with key `deliveryStream` . The following example specifies a delivery stream named `alert-delivery-stream` :
	//
	// `"LogDestination": { "deliveryStream": "alert-delivery-stream" }`.
	LogDestination interface{} `field:"required" json:"logDestination" yaml:"logDestination"`
	// The type of storage destination to send these logs to.
	//
	// You can send logs to an Amazon S3 bucket, a CloudWatch log group, or a Kinesis Data Firehose delivery stream.
	LogDestinationType *string `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// The type of log to send.
	//
	// Alert logs report traffic that matches a stateful rule with an action setting that sends an alert log message. Flow logs are standard network traffic flow logs.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

