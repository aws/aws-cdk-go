package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of a DynamoDB Table.
//
// Use `TableProps` for all table properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var inputFormat inputFormat
//   var key key
//
//   tableOptions := &TableOptions{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("name"),
//   		Type: awscdk.Aws_dynamodb.AttributeType_BINARY,
//   	},
//
//   	// the properties below are optional
//   	BillingMode: awscdk.*Aws_dynamodb.BillingMode_PAY_PER_REQUEST,
//   	ContributorInsightsEnabled: jsii.Boolean(false),
//   	DeletionProtection: jsii.Boolean(false),
//   	Encryption: awscdk.*Aws_dynamodb.TableEncryption_DEFAULT,
//   	EncryptionKey: key,
//   	ImportSource: &ImportSourceSpecification{
//   		Bucket: bucket,
//   		InputFormat: inputFormat,
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		CompressionType: awscdk.*Aws_dynamodb.InputCompressionType_GZIP,
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   	PointInTimeRecovery: jsii.Boolean(false),
//   	ReadCapacity: jsii.Number(123),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	ReplicationRegions: []*string{
//   		jsii.String("replicationRegions"),
//   	},
//   	ReplicationTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	SortKey: &Attribute{
//   		Name: jsii.String("name"),
//   		Type: awscdk.*Aws_dynamodb.AttributeType_BINARY,
//   	},
//   	Stream: awscdk.*Aws_dynamodb.StreamViewType_NEW_IMAGE,
//   	TableClass: awscdk.*Aws_dynamodb.TableClass_STANDARD,
//   	TimeToLiveAttribute: jsii.String("timeToLiveAttribute"),
//   	WaitForReplicationToFinish: jsii.Boolean(false),
//   	WriteCapacity: jsii.Number(123),
//   }
//
type TableOptions struct {
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Default: no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Default: PROVISIONED if `replicationRegions` is not specified, PAY_PER_REQUEST otherwise.
	//
	BillingMode BillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	// Default: false.
	//
	ContributorInsightsEnabled *bool `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Enables deletion protection for the table.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	//
	// > **NOTE**: if you set this to `CUSTOMER_MANAGED` and `encryptionKey` is not
	// > specified, the key that the Tablet generates for you will be created with
	// > default permissions. If you are using CDKv2, these permissions will be
	// > sufficient to enable the key for use with DynamoDB tables.  If you are
	// > using CDKv1, make sure the feature flag
	// > `@aws-cdk/aws-kms:defaultKeyPolicies` is set to `true` in your `cdk.json`.
	// Default: - The table is encrypted with an encryption key managed by DynamoDB, and you are not charged any fee for using it.
	//
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Default: - If `encryption` is set to `TableEncryption.CUSTOMER_MANAGED` and this
	// property is undefined, a new KMS key will be created and associated with this table.
	// If `encryption` and this property are both undefined, then the table is encrypted with
	// an encryption key managed by DynamoDB, and you are not charged any fee for using it.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The properties of data being imported from the S3 bucket source to the table.
	// Default: - no data import from the S3 bucket.
	//
	ImportSource *ImportSourceSpecification `field:"optional" json:"importSource" yaml:"importSource"`
	// Whether point-in-time recovery is enabled.
	// Default: - point-in-time recovery is disabled.
	//
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Default: 5.
	//
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	// Default: - no replica tables are created.
	//
	ReplicationRegions *[]*string `field:"optional" json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Default: Duration.minutes(30)
	//
	ReplicationTimeout awscdk.Duration `field:"optional" json:"replicationTimeout" yaml:"replicationTimeout"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Default: - streams are disabled unless `replicationRegions` is specified.
	//
	Stream StreamViewType `field:"optional" json:"stream" yaml:"stream"`
	// Specify the table class.
	// Default: STANDARD.
	//
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	// Default: - TTL is disabled.
	//
	TimeToLiveAttribute *string `field:"optional" json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// Indicates whether CloudFormation stack waits for replication to finish.
	//
	// If set to false, the CloudFormation resource will mark the resource as
	// created and replication will be completed asynchronously. This property is
	// ignored if replicationRegions property is not set.
	//
	// WARNING:
	// DO NOT UNSET this property if adding/removing multiple replicationRegions
	// in one deployment, as CloudFormation only supports one region replication
	// at a time. CDK overcomes this limitation by waiting for replication to
	// finish before starting new replicationRegion.
	//
	// If the custom resource which handles replication has a physical resource
	// ID with the format `region` instead of `tablename-region` (this would happen
	// if the custom resource hasn't received an event since v1.91.0), DO NOT SET
	// this property to false without making a change to the table name.
	// This will cause the existing replicas to be deleted.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	//
	// Default: true.
	//
	WaitForReplicationToFinish *bool `field:"optional" json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Default: 5.
	//
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

