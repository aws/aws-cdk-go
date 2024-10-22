package awsdms


// The property identifies the exact type of settings for the data provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   settingsProperty := &SettingsProperty{
//   	MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslMode: jsii.String("sslMode"),
//
//   		// the properties below are optional
//   		CertificateArn: jsii.String("certificateArn"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html
//
type CfnDataProvider_SettingsProperty struct {
	// MicrosoftSqlServerSettings property identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-microsoftsqlserversettings
	//
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
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
}

