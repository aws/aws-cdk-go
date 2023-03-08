package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableProps := &cfnTableProps{
//   	keyspaceName: jsii.String("keyspaceName"),
//   	partitionKeyColumns: []interface{}{
//   		&columnProperty{
//   			columnName: jsii.String("columnName"),
//   			columnType: jsii.String("columnType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	billingMode: &billingModeProperty{
//   		mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		provisionedThroughput: &provisionedThroughputProperty{
//   			readCapacityUnits: jsii.Number(123),
//   			writeCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	clusteringKeyColumns: []interface{}{
//   		&clusteringKeyColumnProperty{
//   			column: &columnProperty{
//   				columnName: jsii.String("columnName"),
//   				columnType: jsii.String("columnType"),
//   			},
//
//   			// the properties below are optional
//   			orderBy: jsii.String("orderBy"),
//   		},
//   	},
//   	defaultTimeToLive: jsii.Number(123),
//   	encryptionSpecification: &encryptionSpecificationProperty{
//   		encryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		kmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	},
//   	pointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	regularColumns: []interface{}{
//   		&columnProperty{
//   			columnName: jsii.String("columnName"),
//   			columnType: jsii.String("columnType"),
//   		},
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTableProps struct {
	// The name of the keyspace in which to create the table.
	//
	// The keyspace must already exist.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// One or more columns that uniquely identify every row in the table.
	//
	// Every table must have a partition key.
	PartitionKeyColumns interface{} `field:"required" json:"partitionKeyColumns" yaml:"partitionKeyColumns"`
	// The billing mode for the table, which determines how you'll be charged for reads and writes:.
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	BillingMode interface{} `field:"optional" json:"billingMode" yaml:"billingMode"`
	// One or more columns that determine how the table data is sorted.
	ClusteringKeyColumns interface{} `field:"optional" json:"clusteringKeyColumns" yaml:"clusteringKeyColumns"`
	// The default Time To Live (TTL) value for all rows in a table in seconds.
	//
	// The maximum configurable value is 630,720,000 seconds, which is the equivalent of 20 years. By default, the TTL value for a table is 0, which means data does not expire.
	//
	// For more information, see [Setting the default TTL value for a table](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl) in the *Amazon Keyspaces Developer Guide* .
	DefaultTimeToLive *float64 `field:"optional" json:"defaultTimeToLive" yaml:"defaultTimeToLive"`
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces.
	// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
	//
	// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
	//
	// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
	EncryptionSpecification interface{} `field:"optional" json:"encryptionSpecification" yaml:"encryptionSpecification"`
	// Specifies if point-in-time recovery is enabled or disabled for the table.
	//
	// The options are `PointInTimeRecoveryEnabled=true` and `PointInTimeRecoveryEnabled=false` . If not specified, the default is `PointInTimeRecoveryEnabled=false` .
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// One or more columns that are not part of the primary key - that is, columns that are *not* defined as partition key columns or clustering key columns.
	//
	// You can add regular columns to existing tables by adding them to the template.
	RegularColumns interface{} `field:"optional" json:"regularColumns" yaml:"regularColumns"`
	// The name of the table to be created.
	//
	// The table name is case sensitive. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// *Length constraints:* Minimum length of 3. Maximum length of 255.
	//
	// *Pattern:* `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// A list of key-value pair tags to be attached to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

