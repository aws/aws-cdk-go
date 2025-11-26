package previewawskinesisfirehosemixins


// The top level object for configuring streams with database as a source.
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   databaseSourceConfigurationProperty := &DatabaseSourceConfigurationProperty{
//   	Columns: &DatabaseColumnsProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
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
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	DatabaseSourceVpcConfiguration: &DatabaseSourceVPCConfigurationProperty{
//   		VpcEndpointServiceName: jsii.String("vpcEndpointServiceName"),
//   	},
//   	Digest: jsii.String("digest"),
//   	Endpoint: jsii.String("endpoint"),
//   	Port: jsii.Number(123),
//   	PublicCertificate: jsii.String("publicCertificate"),
//   	SnapshotWatermarkTable: jsii.String("snapshotWatermarkTable"),
//   	SslMode: jsii.String("sslMode"),
//   	SurrogateKeys: []*string{
//   		jsii.String("surrogateKeys"),
//   	},
//   	Tables: &DatabaseTablesProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html
//
type CfnDeliveryStreamPropsMixin_DatabaseSourceConfigurationProperty struct {
	// The list of column patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The list of database patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databases
	//
	Databases interface{} `field:"optional" json:"databases" yaml:"databases"`
	// The structure to configure the authentication methods for Firehose to connect to source database endpoint.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourceauthenticationconfiguration
	//
	DatabaseSourceAuthenticationConfiguration interface{} `field:"optional" json:"databaseSourceAuthenticationConfiguration" yaml:"databaseSourceAuthenticationConfiguration"`
	// The details of the VPC Endpoint Service which Firehose uses to create a PrivateLink to the database.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourcevpcconfiguration
	//
	DatabaseSourceVpcConfiguration interface{} `field:"optional" json:"databaseSourceVpcConfiguration" yaml:"databaseSourceVpcConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-digest
	//
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// The endpoint of the database server.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The port of the database. This can be one of the following values.
	//
	// - 3306 for MySQL database type
	// - 5432 for PostgreSQL database type
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-publiccertificate
	//
	PublicCertificate *string `field:"optional" json:"publicCertificate" yaml:"publicCertificate"`
	// The fully qualified name of the table in source database endpoint that Firehose uses to track snapshot progress.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-snapshotwatermarktable
	//
	SnapshotWatermarkTable *string `field:"optional" json:"snapshotWatermarkTable" yaml:"snapshotWatermarkTable"`
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
	// The list of table patterns in source database endpoint for Firehose to read from.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-tables
	//
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// The type of database engine. This can be one of the following values.
	//
	// - MySQL
	// - PostgreSQL
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

