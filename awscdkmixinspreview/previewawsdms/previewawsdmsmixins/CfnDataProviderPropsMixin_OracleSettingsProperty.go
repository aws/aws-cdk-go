package previewawsdmsmixins


// Provides information that defines an Oracle endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oracleSettingsProperty := &OracleSettingsProperty{
//   	AsmServer: jsii.String("asmServer"),
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Port: jsii.Number(123),
//   	SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   	SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   	SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   	SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html
//
type CfnDataProviderPropsMixin_OracleSettingsProperty struct {
	// For an Oracle source endpoint, your ASM server address.
	//
	// You can set this value from the `asm_server` value. You set `asm_server` as part of the extra connection attribute string to access an Oracle server with Binary Reader that uses ASM. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-asmserver
	//
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Database name for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Endpoint TCP port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Required only if your Oracle endpoint uses Automatic Storage Management (ASM).
	//
	// The full ARN of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the `SecretsManagerOracleAsmSecret` . This `SecretsManagerOracleAsmSecret` has the secret value that allows access to the Oracle ASM of the endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerOracleAsmSecretId` . Or you can specify clear-text values for `AsmUser` , `AsmPassword` , and `AsmServerName` . You can't specify both. For more information on creating this `SecretsManagerOracleAsmSecret` and the `SecretsManagerOracleAsmAccessRoleArn` and `SecretsManagerOracleAsmSecretId` required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-secretsmanageroracleasmaccessrolearn
	//
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Required only if your Oracle endpoint uses Automatic Storage Management (ASM).
	//
	// The full ARN, partial ARN, or friendly name of the `SecretsManagerOracleAsmSecret` that contains the Oracle ASM connection details for the Oracle endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-secretsmanageroracleasmsecretid
	//
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-secretsmanagersecuritydbencryptionaccessrolearn
	//
	SecretsManagerSecurityDbEncryptionAccessRoleArn *string `field:"optional" json:"secretsManagerSecurityDbEncryptionAccessRoleArn" yaml:"secretsManagerSecurityDbEncryptionAccessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-secretsmanagersecuritydbencryptionsecretid
	//
	SecretsManagerSecurityDbEncryptionSecretId *string `field:"optional" json:"secretsManagerSecurityDbEncryptionSecretId" yaml:"secretsManagerSecurityDbEncryptionSecretId"`
	// Fully qualified domain name of the endpoint.
	//
	// For an Amazon RDS Oracle instance, this is the output of [DescribeDBInstances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBInstances.html) , in the `[Endpoint](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Endpoint.html) .Address` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-oraclesettings.html#cfn-dms-dataprovider-oraclesettings-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

