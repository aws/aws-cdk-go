package awsecs


// Configuration for a credential specification (CredSpec) used for a ECS container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialSpecConfig := &CredentialSpecConfig{
//   	Location: jsii.String("location"),
//   	TypePrefix: jsii.String("typePrefix"),
//   }
//
type CredentialSpecConfig struct {
	// Location of the CredSpec file.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Prefix used for the CredSpec string.
	TypePrefix *string `field:"required" json:"typePrefix" yaml:"typePrefix"`
}

