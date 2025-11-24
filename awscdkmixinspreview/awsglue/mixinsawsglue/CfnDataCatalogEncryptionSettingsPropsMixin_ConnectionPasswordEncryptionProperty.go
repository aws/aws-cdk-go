package mixinsawsglue


// The data structure used by the Data Catalog to encrypt the password as part of `CreateConnection` or `UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in the connection properties.
//
// You can enable catalog encryption or only password encryption.
//
// When a `CreationConnection` request arrives containing a password, the Data Catalog first encrypts the password using your AWS  key. It then encrypts the whole connection object again if catalog encryption is also enabled.
//
// This encryption requires that you set AWS  key permissions to enable or restrict access on the password key according to your security requirements. For example, you might want only administrators to have decrypt permission on the password key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectionPasswordEncryptionProperty := &ConnectionPasswordEncryptionProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ReturnConnectionPasswordEncrypted: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html
//
type CfnDataCatalogEncryptionSettingsPropsMixin_ConnectionPasswordEncryptionProperty struct {
	// An AWS  key that is used to encrypt the connection password.
	//
	// If connection password protection is enabled, the caller of `CreateConnection` and `UpdateConnection` needs at least `kms:Encrypt` permission on the specified AWS  key, to encrypt passwords before storing them in the Data Catalog. You can set the decrypt permission to enable or restrict access on the password key according to your security requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// When the `ReturnConnectionPasswordEncrypted` flag is set to "true", passwords remain encrypted in the responses of `GetConnection` and `GetConnections` .
	//
	// This encryption takes effect independently from catalog encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-returnconnectionpasswordencrypted
	//
	ReturnConnectionPasswordEncrypted interface{} `field:"optional" json:"returnConnectionPasswordEncrypted" yaml:"returnConnectionPasswordEncrypted"`
}

