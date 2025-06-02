package awstransfer


// A structure that contains the parameters for an SFTP connector object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sftpConfigProperty := &SftpConfigProperty{
//   	TrustedHostKeys: []*string{
//   		jsii.String("trustedHostKeys"),
//   	},
//   	UserSecretId: jsii.String("userSecretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html
//
type CfnConnector_SftpConfigProperty struct {
	// The public portion of the host key, or keys, that are used to identify the external server to which you are connecting.
	//
	// You can use the `ssh-keyscan` command against the SFTP server to retrieve the necessary key.
	//
	// > `TrustedHostKeys` is optional for `CreateConnector` . If not provided, you can use `TestConnection` to retrieve the server host key during the initial connection attempt, and subsequently update the connector with the observed host key.
	//
	// The three standard SSH public key format elements are `<key type>` , `<body base64>` , and an optional `<comment>` , with spaces between each element. Specify only the `<key type>` and `<body base64>` : do not enter the `<comment>` portion of the key.
	//
	// For the trusted host key, AWS Transfer Family accepts RSA and ECDSA keys.
	//
	// - For RSA keys, the `<key type>` string is `ssh-rsa` .
	// - For ECDSA keys, the `<key type>` string is either `ecdsa-sha2-nistp256` , `ecdsa-sha2-nistp384` , or `ecdsa-sha2-nistp521` , depending on the size of the key you generated.
	//
	// Run this command to retrieve the SFTP server host key, where your SFTP server name is `ftp.host.com` .
	//
	// `ssh-keyscan ftp.host.com`
	//
	// This prints the public host key to standard output.
	//
	// `ftp.host.com ssh-rsa AAAAB3Nza...<long-string-for-public-key`
	//
	// Copy and paste this string into the `TrustedHostKeys` field for the `create-connector` command or into the *Trusted host keys* field in the console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html#cfn-transfer-connector-sftpconfig-trustedhostkeys
	//
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's private key, password, or both.
	//
	// The identifier must be the Amazon Resource Name (ARN) of the secret.
	//
	// > - Required when creating an SFTP connector
	// > - Optional when updating an existing SFTP connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html#cfn-transfer-connector-sftpconfig-usersecretid
	//
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

