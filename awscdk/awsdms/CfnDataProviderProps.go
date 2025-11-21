package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataProviderProps := &CfnDataProviderProps{
//   	Engine: jsii.String("engine"),
//
//   	// the properties below are optional
//   	DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   	DataProviderName: jsii.String("dataProviderName"),
//   	Description: jsii.String("description"),
//   	ExactSettings: jsii.Boolean(false),
//   	Settings: &SettingsProperty{
//   		DocDbSettings: &DocDbSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		IbmDb2LuwSettings: &IbmDb2LuwSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		IbmDb2ZOsSettings: &IbmDb2zOsSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		MariaDbSettings: &MariaDbSettingsProperty{
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		MongoDbSettings: &MongoDbSettingsProperty{
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//
//   			// the properties below are optional
//   			AuthMechanism: jsii.String("authMechanism"),
//   			AuthSource: jsii.String("authSource"),
//   			AuthType: jsii.String("authType"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MySqlSettings: &MySqlSettingsProperty{
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		OracleSettings: &OracleSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			AsmServer: jsii.String("asmServer"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   			SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   			SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   			SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   		},
//   		PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   		RedshiftSettings: &RedshiftSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
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
type CfnDataProviderProps struct {
	// The type of database engine for the data provider.
	//
	// Valid values include `"aurora"` , `"aurora-postgresql"` , `"mysql"` , `"oracle"` , `"postgres"` , `"sqlserver"` , `redshift` , `mariadb` , `mongodb` , `db2` , `db2-zos` , `docdb` , and `sybase` . A value of `"aurora"` represents Amazon Aurora MySQL-Compatible Edition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-engine
	//
	Engine *string `field:"required" json:"engine" yaml:"engine"`
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

