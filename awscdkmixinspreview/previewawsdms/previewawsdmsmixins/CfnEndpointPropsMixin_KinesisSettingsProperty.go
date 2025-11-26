package previewawsdmsmixins


// Provides information that describes an Amazon Kinesis Data Stream endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisSettingsProperty := &KinesisSettingsProperty{
//   	IncludeControlDetails: jsii.Boolean(false),
//   	IncludeNullAndEmpty: jsii.Boolean(false),
//   	IncludePartitionValue: jsii.Boolean(false),
//   	IncludeTableAlterOperations: jsii.Boolean(false),
//   	IncludeTransactionDetails: jsii.Boolean(false),
//   	MessageFormat: jsii.String("messageFormat"),
//   	NoHexPrefix: jsii.Boolean(false),
//   	PartitionIncludeSchemaTable: jsii.Boolean(false),
//   	ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html
//
type CfnEndpointPropsMixin_KinesisSettingsProperty struct {
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-includecontroldetails
	//
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-includenullandempty
	//
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Shows the partition value within the Kinesis message output, unless the partition type is `schema-table-type` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-includepartitionvalue
	//
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-includetablealteroperations
	//
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-includetransactiondetails
	//
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// The output format for the records created on the endpoint.
	//
	// The message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-messageformat
	//
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to an Amazon Kinesis target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-nohexprefix
	//
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kinesis shards. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same shard, which causes throttling. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-partitionincludeschematable
	//
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The Amazon Resource Name (ARN) for the IAM role that AWS DMS uses to write to the Kinesis data stream.
	//
	// The role must allow the `iam:PassRole` action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-serviceaccessrolearn
	//
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The Amazon Resource Name (ARN) for the Amazon Kinesis Data Streams endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

