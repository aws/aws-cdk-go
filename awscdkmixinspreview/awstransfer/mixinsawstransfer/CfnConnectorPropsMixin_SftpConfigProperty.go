package mixinsawstransfer


// A structure that contains the parameters for an SFTP connector object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sftpConfigProperty := &SftpConfigProperty{
//   	MaxConcurrentConnections: jsii.Number(123),
//   	TrustedHostKeys: []*string{
//   		jsii.String("trustedHostKeys"),
//   	},
//   	UserSecretId: jsii.String("userSecretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html
//
type CfnConnectorPropsMixin_SftpConfigProperty struct {
	// Specify the number of concurrent connections that your connector creates to the remote server.
	//
	// The default value is `1` . The maximum values is `5` .
	//
	// > If you are using the AWS Management Console , the default value is `5` .
	//
	// This parameter specifies the number of active connections that your connector can establish with the remote server at the same time. Increasing this value can enhance connector performance when transferring large file batches by enabling parallel operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html#cfn-transfer-connector-sftpconfig-maxconcurrentconnections
	//
	// Default: - 1.
	//
	MaxConcurrentConnections *float64 `field:"optional" json:"maxConcurrentConnections" yaml:"maxConcurrentConnections"`
	// The public portion of the host key, or keys, that are used to identify the external server to which you are connecting.
	//
	// You can use the `ssh-keyscan` command against the SFTP server to retrieve the necessary key.
	//
	// > `TrustedHostKeys` is optional for `CreateConnector` . If not provided, you can use `TestConnection` to retrieve the server host key during the initial connection attempt, and subsequently update the connector with the observed host key.
	//
	// When creating connectors with egress config (VPC_LATTICE type connectors), since host name is not something we can verify, the only accepted trusted host key format is `key-type key-body` without the host name. For example: `ssh-rsa AAAAB3Nza...<long-string-for-public-key>`
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
	// `ftp.host.com ssh-rsa AAAAB3Nza...<long-string-for-public-key>`
	//
	// Copy and paste this string into the `TrustedHostKeys` field for the `create-connector` command or into the *Trusted host keys* field in the console.
	//
	// For VPC Lattice type connectors (VPC_LATTICE), remove the hostname from the key and use only the `key-type key-body` format. In this example, it should be: `ssh-rsa AAAAB3Nza...<long-string-for-public-key>`
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

