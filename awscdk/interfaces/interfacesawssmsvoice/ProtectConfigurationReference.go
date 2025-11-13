package interfacesawssmsvoice


// A reference to a ProtectConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protectConfigurationReference := &ProtectConfigurationReference{
//   	ProtectConfigurationArn: jsii.String("protectConfigurationArn"),
//   	ProtectConfigurationId: jsii.String("protectConfigurationId"),
//   }
//
type ProtectConfigurationReference struct {
	// The ARN of the ProtectConfiguration resource.
	ProtectConfigurationArn *string `field:"required" json:"protectConfigurationArn" yaml:"protectConfigurationArn"`
	// The ProtectConfigurationId of the ProtectConfiguration resource.
	ProtectConfigurationId *string `field:"required" json:"protectConfigurationId" yaml:"protectConfigurationId"`
}

