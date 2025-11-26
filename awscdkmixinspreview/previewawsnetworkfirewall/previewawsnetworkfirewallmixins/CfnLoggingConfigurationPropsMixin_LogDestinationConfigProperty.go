package previewawsnetworkfirewallmixins


// Defines where AWS Network Firewall sends logs for the firewall for one log type.
//
// This is used in logging configuration. You can send each type of log to an Amazon S3 bucket, a CloudWatch log group, or a Kinesis Data Firehose delivery stream.
//
// Network Firewall generates logs for stateful rule groups. You can save alert and flow log types. The stateful rules engine records flow logs for all network traffic that it receives. It records alert logs for traffic that matches stateful rules that have the rule action set to `DROP` or `ALERT` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logDestinationConfigProperty := &LogDestinationConfigProperty{
//   	LogDestination: map[string]*string{
//   		"logDestinationKey": jsii.String("logDestination"),
//   	},
//   	LogDestinationType: jsii.String("logDestinationType"),
//   	LogType: jsii.String("logType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html
//
type CfnLoggingConfigurationPropsMixin_LogDestinationConfigProperty struct {
	// The named location for the logs, provided in a key:value mapping that is specific to the chosen destination type.
	//
	// - For an Amazon S3 bucket, provide the name of the bucket, with key `bucketName` , and optionally provide a prefix, with key `prefix` .
	//
	// The following example specifies an Amazon S3 bucket named `DOC-EXAMPLE-BUCKET` and the prefix `alerts` :
	//
	// `"LogDestination": { "bucketName": "DOC-EXAMPLE-BUCKET", "prefix": "alerts" }`
	// - For a CloudWatch log group, provide the name of the CloudWatch log group, with key `logGroup` . The following example specifies a log group named `alert-log-group` :
	//
	// `"LogDestination": { "logGroup": "alert-log-group" }`
	// - For a Firehose delivery stream, provide the name of the delivery stream, with key `deliveryStream` . The following example specifies a delivery stream named `alert-delivery-stream` :
	//
	// `"LogDestination": { "deliveryStream": "alert-delivery-stream" }`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestination
	//
	LogDestination interface{} `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The type of storage destination to send these logs to.
	//
	// You can send logs to an Amazon S3 bucket, a CloudWatch log group, or a Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestinationtype
	//
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// The type of log to record.
	//
	// You can record the following types of logs from your Network Firewall stateful engine.
	//
	// - `ALERT` - Logs for traffic that matches your stateful rules and that have an action that sends an alert. A stateful rule sends alerts for the rule actions DROP, ALERT, and REJECT. For more information, see the `StatefulRule` property.
	// - `FLOW` - Standard network traffic flow logs. The stateful rules engine records flow logs for all network traffic that it receives. Each flow log record captures the network flow for a specific standard stateless rule group.
	// - `TLS` - Logs for events that are related to TLS inspection. For more information, see [Inspecting SSL/TLS traffic with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-configurations.html) in the *Network Firewall Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logtype
	//
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

