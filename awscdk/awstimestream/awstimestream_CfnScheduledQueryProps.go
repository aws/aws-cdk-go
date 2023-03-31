package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnScheduledQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledQueryProps := &cfnScheduledQueryProps{
//   	errorReportConfiguration: &errorReportConfigurationProperty{
//   		s3Configuration: &s3ConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			encryptionOption: jsii.String("encryptionOption"),
//   			objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		},
//   	},
//   	notificationConfiguration: &notificationConfigurationProperty{
//   		snsConfiguration: &snsConfigurationProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   	},
//   	queryString: jsii.String("queryString"),
//   	scheduleConfiguration: &scheduleConfigurationProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	scheduledQueryExecutionRoleArn: jsii.String("scheduledQueryExecutionRoleArn"),
//
//   	// the properties below are optional
//   	clientToken: jsii.String("clientToken"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	scheduledQueryName: jsii.String("scheduledQueryName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetConfiguration: &targetConfigurationProperty{
//   		timestreamConfiguration: &timestreamConfigurationProperty{
//   			databaseName: jsii.String("databaseName"),
//   			dimensionMappings: []interface{}{
//   				&dimensionMappingProperty{
//   					dimensionValueType: jsii.String("dimensionValueType"),
//   					name: jsii.String("name"),
//   				},
//   			},
//   			tableName: jsii.String("tableName"),
//   			timeColumn: jsii.String("timeColumn"),
//
//   			// the properties below are optional
//   			measureNameColumn: jsii.String("measureNameColumn"),
//   			mixedMeasureMappings: []interface{}{
//   				&mixedMeasureMappingProperty{
//   					measureValueType: jsii.String("measureValueType"),
//
//   					// the properties below are optional
//   					measureName: jsii.String("measureName"),
//   					multiMeasureAttributeMappings: []interface{}{
//   						&multiMeasureAttributeMappingProperty{
//   							measureValueType: jsii.String("measureValueType"),
//   							sourceColumn: jsii.String("sourceColumn"),
//
//   							// the properties below are optional
//   							targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   						},
//   					},
//   					sourceColumn: jsii.String("sourceColumn"),
//   					targetMeasureName: jsii.String("targetMeasureName"),
//   				},
//   			},
//   			multiMeasureMappings: &multiMeasureMappingsProperty{
//   				multiMeasureAttributeMappings: []interface{}{
//   					&multiMeasureAttributeMappingProperty{
//   						measureValueType: jsii.String("measureValueType"),
//   						sourceColumn: jsii.String("sourceColumn"),
//
//   						// the properties below are optional
//   						targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				targetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   			},
//   		},
//   	},
//   }
//
type CfnScheduledQueryProps struct {
	// Configuration for error reporting.
	//
	// Error reports will be generated when a problem is encountered when writing the query results.
	ErrorReportConfiguration interface{} `field:"required" json:"errorReportConfiguration" yaml:"errorReportConfiguration"`
	// Notification configuration for the scheduled query.
	//
	// A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
	NotificationConfiguration interface{} `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The query string to run.
	//
	// Parameter names can be specified in the query string `@` character followed by an identifier. The named Parameter `@scheduled_runtime` is reserved and can be used in the query to get the time at which the query is scheduled to run.
	//
	// The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of `@scheduled_runtime` paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the `@scheduled_runtime` parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Schedule configuration.
	ScheduleConfiguration interface{} `field:"required" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN for the IAM role that Timestream will assume when running the scheduled query.
	ScheduledQueryExecutionRoleArn *string `field:"required" json:"scheduledQueryExecutionRoleArn" yaml:"scheduledQueryExecutionRoleArn"`
	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.
	//
	// Making multiple identical CreateScheduledQuery requests has the same effect as making a single request.
	//
	// - If CreateScheduledQuery is called without a `ClientToken` , the Query SDK generates a `ClientToken` on your behalf.
	// - After 8 hours, any request with the same `ClientToken` is treated as a new request.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest.
	//
	// If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with *alias/*
	//
	// If ErrorReportConfiguration uses `SSE_KMS` as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A name for the query.
	//
	// Scheduled query names must be unique within each Region.
	ScheduledQueryName *string `field:"optional" json:"scheduledQueryName" yaml:"scheduledQueryName"`
	// A list of key-value pairs to label the scheduled query.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Scheduled query target store configuration.
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
}

