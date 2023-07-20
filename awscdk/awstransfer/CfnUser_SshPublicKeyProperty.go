package awstransfer


// Provides information about the public Secure Shell (SSH) key that is associated with a Transfer Family user account for the specific file transfer protocol-enabled server (as identified by `ServerId` ).
//
// The information returned includes the date the key was imported, the public key contents, and the public key ID. A user can store more than one SSH public key associated with their user name on a specific server.
//
// *SshPublicKeyBody*
//
// Specifies the content of the SSH public key as specified by the `PublicKeyId` .
//
// AWS Transfer Family accepts RSA, ECDSA, and ED25519 keys.
//
// Type: String
//
// Length Constraints: Maximum length of 2048.
//
// Required: Yes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sshPublicKeyProperty := &SshPublicKeyProperty{
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-sshpublickey.html
//
type CfnUser_SshPublicKeyProperty struct {
}

