package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFlowLog`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destinationOptions interface{}
//
//   cfnFlowLogProps := &cfnFlowLogProps{
//   	resourceId: jsii.String("resourceId"),
//   	resourceType: jsii.String("resourceType"),
//
//   	// the properties below are optional
//   	deliverLogsPermissionArn: jsii.String("deliverLogsPermissionArn"),
//   	destinationOptions: destinationOptions,
//   	logDestination: jsii.String("logDestination"),
//   	logDestinationType: jsii.String("logDestinationType"),
//   	logFormat: jsii.String("logFormat"),
//   	logGroupName: jsii.String("logGroupName"),
//   	maxAggregationInterval: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trafficType: jsii.String("trafficType"),
//   }
//
type CfnFlowLogProps struct {
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The type of resource for which to create the flow log.
	//
	// For example, if you specified a VPC ID for the `ResourceId` property, specify `VPC` for this property.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account.
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	DeliverLogsPermissionArn *string `field:"optional" json:"deliverLogsPermissionArn" yaml:"deliverLogsPermissionArn"`
	// The destination options. The following options are supported:.
	//
	// - `FileFormat` - The format for the flow log ( `plain-text` | `parquet` ). The default is `plain-text` .
	// - `HiveCompatiblePartitions` - Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3 ( `true` | `false` ). The default is `false` .
	// - `PerHourPartition` - Indicates whether to partition the flow log per hour ( `true` | `false` ). The default is `false` .
	DestinationOptions interface{} `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// The destination to which the flow log data is to be published.
	//
	// Flow log data can be published to a CloudWatch Logs log group or an Amazon S3 bucket. The value specified for this parameter depends on the value specified for `LogDestinationType` .
	//
	// If `LogDestinationType` is not specified or `cloud-watch-logs` , specify the Amazon Resource Name (ARN) of the CloudWatch Logs log group. For example, to publish to a log group called `my-logs` , specify `arn:aws:logs:us-east-1:123456789012:log-group:my-logs` . Alternatively, use `LogGroupName` instead.
	//
	// If LogDestinationType is `s3` , specify the ARN of the Amazon S3 bucket. You can also specify a subfolder in the bucket. To specify a subfolder in the bucket, use the following ARN format: `bucket_ARN/subfolder_name/` . For example, to specify a subfolder named `my-logs` in a bucket named `my-bucket` , use the following ARN: `arn:aws:s3:::my-bucket/my-logs/` . You cannot use `AWSLogs` as a subfolder name. This is a reserved term.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3. To publish flow log data to CloudWatch Logs, specify `cloud-watch-logs` . To publish flow log data to Amazon S3, specify `s3` .
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	//
	// Default: `cloud-watch-logs`.
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// For a list of available fields, see [Flow Log Records](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records) . If you omit this parameter, the flow log is created using the default format. If you specify this parameter, you must specify at least one field.
	//
	// Specify the fields using the `${field-id}` format, separated by spaces.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	//
	// When a network interface is attached to a [Nitro-based instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) , the aggregation interval is always 60 seconds or less, regardless of the value that you specify.
	//
	// Default: 600.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The tags to apply to the flow logs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

