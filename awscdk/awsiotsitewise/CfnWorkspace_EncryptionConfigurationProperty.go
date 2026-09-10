package awsiotsitewise


// The encryption configuration for the workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-workspace-encryptionconfiguration.html
//
type CfnWorkspace_EncryptionConfigurationProperty struct {
	// The type of encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-workspace-encryptionconfiguration.html#cfn-iotsitewise-workspace-encryptionconfiguration-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
}

