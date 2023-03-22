package awsdms


// Provides information that defines an IBMDB2 endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ibmDb2SettingsProperty := &ibmDb2SettingsProperty{
//   	currentLsn: jsii.String("currentLsn"),
//   	maxKBytesPerRead: jsii.Number(123),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	setDataCaptureChanges: jsii.Boolean(false),
//   }
//
type CfnEndpoint_IbmDb2SettingsProperty struct {
	// For ongoing replication (CDC), use CurrentLSN to specify a log sequence number (LSN) where you want the replication to start.
	CurrentLsn *string `field:"optional" json:"currentLsn" yaml:"currentLsn"`
	// Maximum number of bytes per read, as a NUMBER value.
	//
	// The default is 64 KB.
	MaxKBytesPerRead *float64 `field:"optional" json:"maxKBytesPerRead" yaml:"maxKBytesPerRead"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value ofthe AWS Secrets Manager secret that allows access to the Db2 LUW endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the IBMDB2 endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Enables ongoing replication (CDC) as a BOOLEAN value.
	//
	// The default is true.
	SetDataCaptureChanges interface{} `field:"optional" json:"setDataCaptureChanges" yaml:"setDataCaptureChanges"`
}

