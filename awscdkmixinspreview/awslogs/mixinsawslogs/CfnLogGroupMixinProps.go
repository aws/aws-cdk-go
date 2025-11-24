package mixinsawslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLogGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataProtectionPolicy interface{}
//   var fieldIndexPolicies interface{}
//   var resourcePolicyDocument interface{}
//
//   cfnLogGroupMixinProps := &CfnLogGroupMixinProps{
//   	DataProtectionPolicy: dataProtectionPolicy,
//   	FieldIndexPolicies: []interface{}{
//   		fieldIndexPolicies,
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogGroupClass: jsii.String("logGroupClass"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	ResourcePolicyDocument: resourcePolicyDocument,
//   	RetentionInDays: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
//
type CfnLogGroupMixinProps struct {
	// Creates a data protection policy and assigns it to the log group.
	//
	// A data protection policy can help safeguard sensitive data that's ingested by the log group by auditing and masking the sensitive log data. When a user who does not have permission to view masked data views a log event that includes masked data, the sensitive data is replaced by asterisks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-dataprotectionpolicy
	//
	DataProtectionPolicy interface{} `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// Creates or updates a *field index policy* for the specified log group.
	//
	// Only log groups in the Standard log class support field index policies. For more information about log classes, see [Log classes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html) .
	//
	// You can use field index policies to create *field indexes* on fields found in log events in the log group. Creating field indexes lowers the costs for CloudWatch Logs Insights queries that reference those field indexes, because these queries attempt to skip the processing of log events that are known to not match the indexed field. Good fields to index are fields that you often need to query for and fields that have high cardinality of values Common examples of indexes include request ID, session ID, userID, and instance IDs. For more information, see [Create field indexes to improve query performance and reduce costs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html) .
	//
	// Currently, this array supports only one field index policy object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-fieldindexpolicies
	//
	FieldIndexPolicies interface{} `field:"optional" json:"fieldIndexPolicies" yaml:"fieldIndexPolicies"`
	// The Amazon Resource Name (ARN) of the AWS  key to use when encrypting log data.
	//
	// To associate an AWS  key with the log group, specify the ARN of that KMS key here. If you do so, ingested data is encrypted using this key. This association is stored as long as the data encrypted with the KMS key is still within CloudWatch Logs . This enables CloudWatch Logs to decrypt this data whenever it is requested.
	//
	// If you attempt to associate a KMS key with the log group but the KMS key doesn't exist or is deactivated, you will receive an `InvalidParameterException` error.
	//
	// Log group data is always encrypted in CloudWatch Logs . If you omit this key, the encryption does not use AWS  . For more information, see [Encrypt log data in CloudWatch Logs using AWS Key Management Service](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/encrypt-log-data-kms.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the log group class for this log group. There are two classes:.
	//
	// - The `Standard` log class supports all CloudWatch Logs features.
	// - The `Infrequent Access` log class supports a subset of CloudWatch Logs features and incurs lower costs.
	//
	// For details about the features supported by each class, see [Log classes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-loggroupclass
	//
	// Default: - "STANDARD".
	//
	LogGroupClass *string `field:"optional" json:"logGroupClass" yaml:"logGroupClass"`
	// The name of the log group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Creates or updates a resource policy for the specified log group that allows other services to put log events to this account.
	//
	// A LogGroup can have 1 resource policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-resourcepolicydocument
	//
	ResourcePolicyDocument interface{} `field:"optional" json:"resourcePolicyDocument" yaml:"resourcePolicyDocument"`
	// The number of days to retain the log events in the specified log group.
	//
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, and 3653.
	//
	// To set a log group so that its log events do not expire, do not specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// An array of key-value pairs to apply to the log group.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-logs-loggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

