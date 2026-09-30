package awshealthagent


// Encryption context for the Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionContextProperty := &EncryptionContextProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthagent-domain-encryptioncontext.html
//
type CfnDomainPropsMixin_EncryptionContextProperty struct {
	// The type of encryption key used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthagent-domain-encryptioncontext.html#cfn-healthagent-domain-encryptioncontext-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
}

