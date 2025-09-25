package awsdevicefarm


// A reference to a VPCEConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEConfigurationReference := &VPCEConfigurationReference{
//   	VpceConfigurationArn: jsii.String("vpceConfigurationArn"),
//   }
//
type VPCEConfigurationReference struct {
	// The Arn of the VPCEConfiguration resource.
	VpceConfigurationArn *string `field:"required" json:"vpceConfigurationArn" yaml:"vpceConfigurationArn"`
}

