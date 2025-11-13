package interfacesawsinspectorv2


// A reference to a CodeSecurityScanConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeSecurityScanConfigurationReference := &CodeSecurityScanConfigurationReference{
//   	CodeSecurityScanConfigurationArn: jsii.String("codeSecurityScanConfigurationArn"),
//   }
//
type CodeSecurityScanConfigurationReference struct {
	// The Arn of the CodeSecurityScanConfiguration resource.
	CodeSecurityScanConfigurationArn *string `field:"required" json:"codeSecurityScanConfigurationArn" yaml:"codeSecurityScanConfigurationArn"`
}

