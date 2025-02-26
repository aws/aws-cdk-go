package awskinesisfirehose


// The top level object for configuring streams with database as a source.
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseSourceConfigurationProperty := &DatabaseSourceConfigurationProperty{
//   	Databases: &DatabasesProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	DatabaseSourceAuthenticationConfiguration: &DatabaseSourceAuthenticationConfigurationProperty{
//   		SecretsManagerConfiguration: &SecretsManagerConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	DatabaseSourceVpcConfiguration: &DatabaseSourceVPCConfigurationProperty{
//   		VpcEndpointServiceName: jsii.String("vpcEndpointServiceName"),
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	Port: jsii.Number(123),
//   	SnapshotWatermarkTable: jsii.String("snapshotWatermarkTable"),
//   	Tables: &DatabaseTablesProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Columns: &DatabaseColumnsProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	Digest: jsii.String("digest"),
//   	PublicCertificate: jsii.String("publicCertificate"),
//   	SslMode: jsii.String("sslMode"),
//   	SurrogateKeys: []*string{
//   		jsii.String("surrogateKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html
//
type CfnDeliveryStream_DatabaseSourceConfigurationProperty struct {
	// The list of database patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databases
	//
	Databases interface{} `field:"required" json:"databases" yaml:"databases"`
	// The structure to configure the authentication methods for Firehose to connect to source database endpoint.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourceauthenticationconfiguration
	//
	DatabaseSourceAuthenticationConfiguration interface{} `field:"required" json:"databaseSourceAuthenticationConfiguration" yaml:"databaseSourceAuthenticationConfiguration"`
	// The details of the VPC Endpoint Service which Firehose uses to create a PrivateLink to the database.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourcevpcconfiguration
	//
	DatabaseSourceVpcConfiguration interface{} `field:"required" json:"databaseSourceVpcConfiguration" yaml:"databaseSourceVpcConfiguration"`
	// The endpoint of the database server.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The port of the database. This can be one of the following values.
	//
	// - 3306 for MySQL database type
	// - 5432 for PostgreSQL database type
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The fully qualified name of the table in source database endpoint that Firehose uses to track snapshot progress.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-snapshotwatermarktable
	//
	SnapshotWatermarkTable *string `field:"required" json:"snapshotWatermarkTable" yaml:"snapshotWatermarkTable"`
	// The list of table patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-tables
	//
	Tables interface{} `field:"required" json:"tables" yaml:"tables"`
	// The type of database engine. This can be one of the following values.
	//
	// - MySQL
	// - PostgreSQL
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The list of column patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-digest
	//
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-publiccertificate
	//
	PublicCertificate *string `field:"optional" json:"publicCertificate" yaml:"publicCertificate"`
	// The mode to enable or disable SSL when Firehose connects to the database endpoint.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// The optional list of table and column names used as unique key columns when taking snapshot if the tables don’t have primary keys configured.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-surrogatekeys
	//
	SurrogateKeys *[]*string `field:"optional" json:"surrogateKeys" yaml:"surrogateKeys"`
}

