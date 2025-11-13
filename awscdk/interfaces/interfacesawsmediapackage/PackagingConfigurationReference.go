package interfacesawsmediapackage


// A reference to a PackagingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packagingConfigurationReference := &PackagingConfigurationReference{
//   	PackagingConfigurationArn: jsii.String("packagingConfigurationArn"),
//   	PackagingConfigurationId: jsii.String("packagingConfigurationId"),
//   }
//
type PackagingConfigurationReference struct {
	// The ARN of the PackagingConfiguration resource.
	PackagingConfigurationArn *string `field:"required" json:"packagingConfigurationArn" yaml:"packagingConfigurationArn"`
	// The Id of the PackagingConfiguration resource.
	PackagingConfigurationId *string `field:"required" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
}

