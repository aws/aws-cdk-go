package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnScheduledQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledQueryProps := &CfnScheduledQueryProps{
//   	ErrorReportConfiguration: &ErrorReportConfigurationProperty{
//   		S3Configuration: &S3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			EncryptionOption: jsii.String("encryptionOption"),
//   			ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		},
//   	},
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		SnsConfiguration: &SnsConfigurationProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   	QueryString: jsii.String("queryString"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	ScheduledQueryExecutionRoleArn: jsii.String("scheduledQueryExecutionRoleArn"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ScheduledQueryName: jsii.String("scheduledQueryName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetConfiguration: &TargetConfigurationProperty{
//   		TimestreamConfiguration: &TimestreamConfigurationProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			DimensionMappings: []interface{}{
//   				&DimensionMappingProperty{
//   					DimensionValueType: jsii.String("dimensionValueType"),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			TableName: jsii.String("tableName"),
//   			TimeColumn: jsii.String("timeColumn"),
//
//   			// the properties below are optional
//   			MeasureNameColumn: jsii.String("measureNameColumn"),
//   			MixedMeasureMappings: []interface{}{
//   				&MixedMeasureMappingProperty{
//   					MeasureValueType: jsii.String("measureValueType"),
//
//   					// the properties below are optional
//   					MeasureName: jsii.String("measureName"),
//   					MultiMeasureAttributeMappings: []interface{}{
//   						&MultiMeasureAttributeMappingProperty{
//   							MeasureValueType: jsii.String("measureValueType"),
//   							SourceColumn: jsii.String("sourceColumn"),
//
//   							// the properties below are optional
//   							TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   						},
//   					},
//   					SourceColumn: jsii.String("sourceColumn"),
//   					TargetMeasureName: jsii.String("targetMeasureName"),
//   				},
//   			},
//   			MultiMeasureMappings: &MultiMeasureMappingsProperty{
//   				MultiMeasureAttributeMappings: []interface{}{
//   					&MultiMeasureAttributeMappingProperty{
//   						MeasureValueType: jsii.String("measureValueType"),
//   						SourceColumn: jsii.String("sourceColumn"),
//
//   						// the properties below are optional
//   						TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				TargetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html
//
type CfnScheduledQueryProps struct {
	// Configuration for error reporting.
	//
	// Error reports will be generated when a problem is encountered when writing the query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-errorreportconfiguration
	//
	ErrorReportConfiguration interface{} `field:"required" json:"errorReportConfiguration" yaml:"errorReportConfiguration"`
	// Notification configuration for the scheduled query.
	//
	// A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-notificationconfiguration
	//
	NotificationConfiguration interface{} `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The query string to run.
	//
	// Parameter names can be specified in the query string `@` character followed by an identifier. The named Parameter `@scheduled_runtime` is reserved and can be used in the query to get the time at which the query is scheduled to run.
	//
	// The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of `@scheduled_runtime` paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the `@scheduled_runtime` parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Schedule configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduleconfiguration
	//
	ScheduleConfiguration interface{} `field:"required" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN for the IAM role that Timestream will assume when running the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduledqueryexecutionrolearn
	//
	ScheduledQueryExecutionRoleArn *string `field:"required" json:"scheduledQueryExecutionRoleArn" yaml:"scheduledQueryExecutionRoleArn"`
	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.
	//
	// Making multiple identical CreateScheduledQuery requests has the same effect as making a single request.
	//
	// - If CreateScheduledQuery is called without a `ClientToken` , the Query SDK generates a `ClientToken` on your behalf.
	// - After 8 hours, any request with the same `ClientToken` is treated as a new request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest.
	//
	// If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with *alias/*
	//
	// If ErrorReportConfiguration uses `SSE_KMS` as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A name for the query.
	//
	// Scheduled query names must be unique within each Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduledqueryname
	//
	ScheduledQueryName *string `field:"optional" json:"scheduledQueryName" yaml:"scheduledQueryName"`
	// A list of key-value pairs to label the scheduled query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Scheduled query target store configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-targetconfiguration
	//
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
}

