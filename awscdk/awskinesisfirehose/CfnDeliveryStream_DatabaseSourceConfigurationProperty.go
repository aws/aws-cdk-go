package awskinesisfirehose


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databases
	//
	Databases interface{} `field:"required" json:"databases" yaml:"databases"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourceauthenticationconfiguration
	//
	DatabaseSourceAuthenticationConfiguration interface{} `field:"required" json:"databaseSourceAuthenticationConfiguration" yaml:"databaseSourceAuthenticationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-databasesourcevpcconfiguration
	//
	DatabaseSourceVpcConfiguration interface{} `field:"required" json:"databaseSourceVpcConfiguration" yaml:"databaseSourceVpcConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-snapshotwatermarktable
	//
	SnapshotWatermarkTable *string `field:"required" json:"snapshotWatermarkTable" yaml:"snapshotWatermarkTable"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-tables
	//
	Tables interface{} `field:"required" json:"tables" yaml:"tables"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-digest
	//
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-publiccertificate
	//
	PublicCertificate *string `field:"optional" json:"publicCertificate" yaml:"publicCertificate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceconfiguration-surrogatekeys
	//
	SurrogateKeys *[]*string `field:"optional" json:"surrogateKeys" yaml:"surrogateKeys"`
}

