# Amazon DynamoDB Construct Library

Here is a minimal deployable DynamoDB table definition:

```go
table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
})
```

## Importing existing tables

To import an existing table into your CDK application, use the `Table.fromTableName`, `Table.fromTableArn` or `Table.fromTableAttributes`
factory method. This method accepts table name or table ARN which describes the properties of an already
existing table:

```go
var user user

table := dynamodb.Table_FromTableArn(this, jsii.String("ImportedTable"), jsii.String("arn:aws:dynamodb:us-east-1:111111111:table/my-table"))
// now you can just call methods on the table
table.GrantReadWriteData(user)
```

If you intend to use the `tableStreamArn` (including indirectly, for example by creating an
`@aws-cdk/aws-lambda-event-source.DynamoEventSource` on the imported table), you *must* use the
`Table.fromTableAttributes` method and the `tableStreamArn` property *must* be populated.

In order to grant permissions to indexes on imported tables you can either set `grantIndexPermissions` to `true`, or you can provide the indexes via the `globalIndexes` or `localIndexes` properties. This will enable `grant*` methods to also grant permissions to *all* table indexes.

## Keys

When a table is defined, you must define it's schema using the `partitionKey`
(required) and `sortKey` (optional) properties.

## Billing Mode

DynamoDB supports two billing modes:

* PROVISIONED - the default mode where the table and global secondary indexes have configured read and write capacity.
* PAY_PER_REQUEST - on-demand pricing and scaling. You only pay for what you use and there is no read and write capacity for the table or its global secondary indexes.

```go
table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.

## Table Class

DynamoDB supports two table classes:

* STANDARD - the default mode, and is recommended for the vast majority of workloads.
* STANDARD_INFREQUENT_ACCESS - optimized for tables where storage is the dominant cost.

```go
table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html

## Configure AutoScaling for your table

You can have DynamoDB automatically raise and lower the read and write capacities
of your table by setting up autoscaling. You can use this to either keep your
tables at a desired utilization level, or by scaling up and down at pre-configured
times of the day:

Auto-scaling is only relevant for tables with the billing mode, PROVISIONED.

```go
readScaling := table.AutoScaleReadCapacity(&EnableScalingProps{
	MinCapacity: jsii.Number(1),
	MaxCapacity: jsii.Number(50),
})

readScaling.ScaleOnUtilization(&UtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(50),
})

readScaling.ScaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &ScalingSchedule{
	Schedule: appscaling.Schedule_Cron(&CronOptions{
		Hour: jsii.String("8"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(20),
})

readScaling.ScaleOnSchedule(jsii.String("ScaleDownAtNight"), &ScalingSchedule{
	Schedule: appscaling.Schedule_*Cron(&CronOptions{
		Hour: jsii.String("20"),
		Minute: jsii.String("0"),
	}),
	MaxCapacity: jsii.Number(20),
})
```

Further reading:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/AutoScaling.html
https://aws.amazon.com/blogs/database/how-to-use-aws-cloudformation-to-configure-auto-scaling-for-amazon-dynamodb-tables-and-indexes/

## Amazon DynamoDB Global Tables

You can create DynamoDB Global Tables by setting the `replicationRegions` property on a `Table`:

```go
globalTable := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ReplicationRegions: []*string{
		jsii.String("us-east-1"),
		jsii.String("us-east-2"),
		jsii.String("us-west-2"),
	},
})
```

When doing so, a CloudFormation Custom Resource will be added to the stack in order to create the replica tables in the
selected regions.

The default billing mode for Global Tables is `PAY_PER_REQUEST`.
If you want to use `PROVISIONED`,
you have to make sure write auto-scaling is enabled for that Table:

```go
globalTable := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ReplicationRegions: []*string{
		jsii.String("us-east-1"),
		jsii.String("us-east-2"),
		jsii.String("us-west-2"),
	},
	BillingMode: dynamodb.BillingMode_PROVISIONED,
})

globalTable.AutoScaleWriteCapacity(&EnableScalingProps{
	MinCapacity: jsii.Number(1),
	MaxCapacity: jsii.Number(10),
}).ScaleOnUtilization(&UtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(75),
})
```

When adding a replica region for a large table, you might want to increase the
timeout for the replication operation:

```go
globalTable := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	ReplicationRegions: []*string{
		jsii.String("us-east-1"),
		jsii.String("us-east-2"),
		jsii.String("us-west-2"),
	},
	ReplicationTimeout: awscdk.Duration_Hours(jsii.Number(2)),
})
```

A maximum of 10 tables with replication can be added to a stack without a limit increase for
[managed policies attached to an IAM role](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entities).
This is because more than 10 managed policies will be attached to the DynamoDB service replication role - one policy per replication table.
Consider splitting your tables across multiple stacks if your reach this limit.

## Encryption

All user data stored in Amazon DynamoDB is fully encrypted at rest. When creating a new table, you can choose to encrypt using the following customer master keys (CMK) to encrypt your table:

* AWS owned CMK - By default, all tables are encrypted under an AWS owned customer master key (CMK) in the DynamoDB service account (no additional charges apply).
* AWS managed CMK - AWS KMS keys (one per region) are created in your account, managed, and used on your behalf by AWS DynamoDB (AWS KMS charges apply).
* Customer managed CMK - You have full control over the KMS key used to encrypt the DynamoDB Table (AWS KMS charges apply).

Creating a Table encrypted with a customer managed CMK:

```go
table := dynamodb.NewTable(this, jsii.String("MyTable"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryption_CUSTOMER_MANAGED,
})

// You can access the CMK that was added to the stack on your behalf by the Table construct via:
tableEncryptionKey := table.EncryptionKey
```

You can also supply your own key:

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


encryptionKey := kms.NewKey(this, jsii.String("Key"), &KeyProps{
	EnableKeyRotation: jsii.Boolean(true),
})
table := dynamodb.NewTable(this, jsii.String("MyTable"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryption_CUSTOMER_MANAGED,
	EncryptionKey: EncryptionKey,
})
```

In order to use the AWS managed CMK instead, change the code to:

```go
table := dynamodb.NewTable(this, jsii.String("MyTable"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	Encryption: dynamodb.TableEncryption_AWS_MANAGED,
})
```

## Get schema of table or secondary indexes

To get the partition key and sort key of the table or indexes you have configured:

```go
var table table

schema := table.Schema()
partitionKey := schema.PartitionKey
sortKey := schema.SortKey
```

## Kinesis Stream

A Kinesis Data Stream can be configured on the DynamoDB table to capture item-level changes.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"


stream := kinesis.NewStream(this, jsii.String("Stream"))

table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
	KinesisStream: stream,
})
```

## Alarm metrics

Alarms can be configured on the DynamoDB table to captured metric data

```go
// Example automatically generated from non-compiling source. May contain errors.
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
	PartitionKey: &Attribute{
		Name: jsii.String("id"),
		Type: dynamodb.AttributeType_STRING,
	},
})

metric := table.metricThrottledRequestsForOperations(&OperationsMetricOptions{
	Operations: []operation{
		dynamodb.*operation_PUT_ITEM,
	},
	Period: awscdk.Duration_Minutes(jsii.Number(1)),
})

cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &AlarmProps{
	Metric: metric,
	EvaluationPeriods: jsii.Number(1),
	Threshold: jsii.Number(1),
})
```
