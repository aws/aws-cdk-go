package awsdms


// The property identifies the exact type of settings for the data provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   settingsProperty := &SettingsProperty{
//   	DocDbSettings: &DocDbSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	IbmDb2LuwSettings: &IbmDb2LuwSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	IbmDb2ZOsSettings: &IbmDb2zOsSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	MariaDbSettings: &MariaDbSettingsProperty{
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	MongoDbSettings: &MongoDbSettingsProperty{
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//
//   		// the properties below are optional
//   		AuthMechanism: jsii.String("authMechanism"),
//   		AuthSource: jsii.String("authSource"),
//   		AuthType: jsii.String("authType"),
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	MySqlSettings: &MySqlSettingsProperty{
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	OracleSettings: &OracleSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		AsmServer: jsii.String("asmServer"),
//   		CertificateArn: jsii.String("certificateArn"),
//   		SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   		SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   	},
//   	PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	RedshiftSettings: &RedshiftSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   	},
//   	SybaseAseSettings: &SybaseAseSettingsProperty{
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		EncryptPassword: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html
//
type CfnDataProvider_SettingsProperty struct {
	// DocDbSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-docdbsettings
	//
	DocDbSettings interface{} `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// IbmDb2LuwSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-ibmdb2luwsettings
	//
	IbmDb2LuwSettings interface{} `field:"optional" json:"ibmDb2LuwSettings" yaml:"ibmDb2LuwSettings"`
	// IbmDb2zOsSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-ibmdb2zossettings
	//
	IbmDb2ZOsSettings interface{} `field:"optional" json:"ibmDb2ZOsSettings" yaml:"ibmDb2ZOsSettings"`
	// MariaDbSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-mariadbsettings
	//
	MariaDbSettings interface{} `field:"optional" json:"mariaDbSettings" yaml:"mariaDbSettings"`
	// MicrosoftSqlServerSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-microsoftsqlserversettings
	//
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// MongoDbSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-mongodbsettings
	//
	MongoDbSettings interface{} `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// MySqlSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-mysqlsettings
	//
	MySqlSettings interface{} `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// OracleSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-oraclesettings
	//
	OracleSettings interface{} `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// PostgreSqlSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-postgresqlsettings
	//
	PostgreSqlSettings interface{} `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// RedshiftSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-redshiftsettings
	//
	RedshiftSettings interface{} `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// SybaseAseSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-sybaseasesettings
	//
	SybaseAseSettings interface{} `field:"optional" json:"sybaseAseSettings" yaml:"sybaseAseSettings"`
}

