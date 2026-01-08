package previewawsdmsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataProviderMixinProps := &CfnDataProviderMixinProps{
//   	DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   	DataProviderName: jsii.String("dataProviderName"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	ExactSettings: jsii.Boolean(false),
//   	Settings: &SettingsProperty{
//   		DocDbSettings: &DocDbSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		IbmDb2LuwSettings: &IbmDb2LuwSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		IbmDb2ZOsSettings: &IbmDb2zOsSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MariaDbSettings: &MariaDbSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MongoDbSettings: &MongoDbSettingsProperty{
//   			AuthMechanism: jsii.String("authMechanism"),
//   			AuthSource: jsii.String("authSource"),
//   			AuthType: jsii.String("authType"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MySqlSettings: &MySqlSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		OracleSettings: &OracleSettingsProperty{
//   			AsmServer: jsii.String("asmServer"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   			SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   			SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   			SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		RedshiftSettings: &RedshiftSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   		},
//   		SybaseAseSettings: &SybaseAseSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			EncryptPassword: jsii.Boolean(false),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html
//
type CfnDataProviderMixinProps struct {
	// The identifier of the data provider.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-dataprovideridentifier
	//
	DataProviderIdentifier *string `field:"optional" json:"dataProviderIdentifier" yaml:"dataProviderIdentifier"`
	// The name of the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-dataprovidername
	//
	DataProviderName *string `field:"optional" json:"dataProviderName" yaml:"dataProviderName"`
	// A description of the data provider.
	//
	// Descriptions can have up to 31 characters. A description can contain only ASCII letters, digits, and hyphens ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of database engine for the data provider.
	//
	// Valid values include `"aurora"` , `"aurora-postgresql"` , `"mysql"` , `"oracle"` , `"postgres"` , `"sqlserver"` , `redshift` , `mariadb` , `mongodb` , `db2` , `db2-zos` , `docdb` , and `sybase` . A value of `"aurora"` represents Amazon Aurora MySQL-Compatible Edition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The property describes the exact settings which can be modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-exactsettings
	//
	// Default: - false.
	//
	ExactSettings interface{} `field:"optional" json:"exactSettings" yaml:"exactSettings"`
	// The settings in JSON format for a data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

