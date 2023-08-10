package awstransfer


// Configuration for an SFTP connector.
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
	// List of public host keys, for the external server to which you are connecting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html#cfn-transfer-connector-sftpconfig-trustedhostkeys
	//
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-sftpconfig.html#cfn-transfer-connector-sftpconfig-usersecretid
	//
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

