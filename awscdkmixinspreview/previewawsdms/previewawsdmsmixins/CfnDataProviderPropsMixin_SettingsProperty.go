package previewawsdmsmixins


// The property identifies the exact type of settings for the data provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   settingsProperty := &SettingsProperty{
//   	DocDbSettings: &DocDbSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	IbmDb2LuwSettings: &IbmDb2LuwSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	IbmDb2ZOsSettings: &IbmDb2zOsSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	MariaDbSettings: &MariaDbSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	MongoDbSettings: &MongoDbSettingsProperty{
//   		AuthMechanism: jsii.String("authMechanism"),
//   		AuthSource: jsii.String("authSource"),
//   		AuthType: jsii.String("authType"),
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	MySqlSettings: &MySqlSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	OracleSettings: &OracleSettingsProperty{
//   		AsmServer: jsii.String("asmServer"),
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   		SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   	RedshiftSettings: &RedshiftSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   	},
//   	SybaseAseSettings: &SybaseAseSettingsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		EncryptPassword: jsii.Boolean(false),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html
//
type CfnDataProviderPropsMixin_SettingsProperty struct {
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

