package interfacesawsgreengrass


// A reference to a DeviceDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceDefinitionVersionReference := &DeviceDefinitionVersionReference{
//   	DeviceDefinitionVersionId: jsii.String("deviceDefinitionVersionId"),
//   }
//
type DeviceDefinitionVersionReference struct {
	// The Id of the DeviceDefinitionVersion resource.
	DeviceDefinitionVersionId *string `field:"required" json:"deviceDefinitionVersionId" yaml:"deviceDefinitionVersionId"`
}

