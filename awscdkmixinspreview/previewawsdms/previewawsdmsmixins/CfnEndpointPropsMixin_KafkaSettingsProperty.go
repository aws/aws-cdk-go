package previewawsdmsmixins


// Provides information that describes an Apache Kafka endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kafkaSettingsProperty := &KafkaSettingsProperty{
//   	Broker: jsii.String("broker"),
//   	IncludeControlDetails: jsii.Boolean(false),
//   	IncludeNullAndEmpty: jsii.Boolean(false),
//   	IncludePartitionValue: jsii.Boolean(false),
//   	IncludeTableAlterOperations: jsii.Boolean(false),
//   	IncludeTransactionDetails: jsii.Boolean(false),
//   	MessageFormat: jsii.String("messageFormat"),
//   	MessageMaxBytes: jsii.Number(123),
//   	NoHexPrefix: jsii.Boolean(false),
//   	PartitionIncludeSchemaTable: jsii.Boolean(false),
//   	SaslPassword: jsii.String("saslPassword"),
//   	SaslUserName: jsii.String("saslUserName"),
//   	SecurityProtocol: jsii.String("securityProtocol"),
//   	SslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   	SslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   	SslClientKeyArn: jsii.String("sslClientKeyArn"),
//   	SslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   	Topic: jsii.String("topic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html
//
type CfnEndpointPropsMixin_KafkaSettingsProperty struct {
	// A comma-separated list of one or more broker locations in your Kafka cluster that host your Kafka instance.
	//
	// Specify each broker location in the form `*broker-hostname-or-ip* : *port*` . For example, `"ec2-12-345-678-901.compute-1.amazonaws.com:2345"` . For more information and examples of specifying a list of broker locations, see [Using Apache Kafka as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-broker
	//
	Broker *string `field:"optional" json:"broker" yaml:"broker"`
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-includecontroldetails
	//
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-includenullandempty
	//
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Shows the partition value within the Kafka message output unless the partition type is `schema-table-type` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-includepartitionvalue
	//
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table` , `drop-table` , `add-column` , `drop-column` , and `rename-column` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-includetablealteroperations
	//
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// This information includes a commit timestamp, a log position, and values for `transaction_id` , previous `transaction_id` , and `transaction_record_id` (the record offset within a transaction). The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-includetransactiondetails
	//
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// The output format for the records created on the endpoint.
	//
	// The message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-messageformat
	//
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// The maximum size in bytes for records created on the endpoint The default is 1,000,000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-messagemaxbytes
	//
	MessageMaxBytes *float64 `field:"optional" json:"messageMaxBytes" yaml:"messageMaxBytes"`
	// Set this optional parameter to `true` to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the `NoHexPrefix` endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-nohexprefix
	//
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type` .
	//
	// Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-partitionincludeschematable
	//
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The secure password that you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-saslpassword
	//
	SaslPassword *string `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// The secure user name you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-saslusername
	//
	SaslUserName *string `field:"optional" json:"saslUserName" yaml:"saslUserName"`
	// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS).
	//
	// Options include `ssl-encryption` , `ssl-authentication` , and `sasl-ssl` . `sasl-ssl` requires `SaslUsername` and `SaslPassword` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-securityprotocol
	//
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-sslcacertificatearn
	//
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-sslclientcertificatearn
	//
	SslClientCertificateArn *string `field:"optional" json:"sslClientCertificateArn" yaml:"sslClientCertificateArn"`
	// The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-sslclientkeyarn
	//
	SslClientKeyArn *string `field:"optional" json:"sslClientKeyArn" yaml:"sslClientKeyArn"`
	// The password for the client private key used to securely connect to a Kafka target endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-sslclientkeypassword
	//
	SslClientKeyPassword *string `field:"optional" json:"sslClientKeyPassword" yaml:"sslClientKeyPassword"`
	// The topic to which you migrate the data.
	//
	// If you don't specify a topic, AWS DMS specifies `"kafka-default-topic"` as the migration topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kafkasettings.html#cfn-dms-endpoint-kafkasettings-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

