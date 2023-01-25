package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// The ID of the resource to monitor.
	//
	// For example, if the resource type is `VPC` , specify the ID of the VPC.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The type of resource to monitor.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account.
	//
	// This parameter is required if the destination type is `cloud-watch-logs` and unsupported otherwise.
	DeliverLogsPermissionArn *string `field:"optional" json:"deliverLogsPermissionArn" yaml:"deliverLogsPermissionArn"`
	// The destination options. The following options are supported:.
	//
	// - `FileFormat` - The format for the flow log ( `plain-text` | `parquet` ). The default is `plain-text` .
	// - `HiveCompatiblePartitions` - Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3 ( `true` | `false` ). The default is `false` .
	// - `PerHourPartition` - Indicates whether to partition the flow log per hour ( `true` | `false` ). The default is `false` .
	DestinationOptions interface{} `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// The destination for the flow log data. The meaning of this parameter depends on the destination type.
	//
	// - If the destination type is `cloud-watch-logs` , specify the ARN of a CloudWatch Logs log group. For example:
	//
	// arn:aws:logs: *region* : *account_id* :log-group: *my_group*
	//
	// Alternatively, use the `LogGroupName` parameter.
	// - If the destination type is `s3` , specify the ARN of an S3 bucket. For example:
	//
	// arn:aws:s3::: *my_bucket* / *my_subfolder* /
	//
	// The subfolder is optional. Note that you can't use `AWSLogs` as a subfolder name.
	// - If the destination type is `kinesis-data-firehose` , specify the ARN of a Kinesis Data Firehose delivery stream. For example:
	//
	// arn:aws:firehose: *region* : *account_id* :deliverystream: *my_stream*.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The type of destination for the flow log data.
	//
	// Default: `cloud-watch-logs`.
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// If you omit this parameter, the flow log is created using the default format. If you specify this parameter, you must include at least one field. For more information about the available fields, see [Flow log records](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records) in the *Amazon VPC User Guide* or [Transit Gateway Flow Log records](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-flow-logs.html#flow-log-records) in the *AWS Transit Gateway Guide* .
	//
	// Specify the fields using the `${field-id}` format, separated by spaces.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.
	//
	// This parameter is valid only if the destination type is `cloud-watch-logs` .
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// The possible values are 60 seconds (1 minute) or 600 seconds (10 minutes). This parameter must be 60 seconds for transit gateway resource types.
	//
	// When a network interface is attached to a [Nitro-based instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) , the aggregation interval is always 60 seconds or less, regardless of the value that you specify.
	//
	// Default: 600.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The tags to apply to the flow logs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of traffic to monitor (accepted traffic, rejected traffic, or all traffic).
	//
	// This parameter is not supported for transit gateway resource types. It is required for the other resource types.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

