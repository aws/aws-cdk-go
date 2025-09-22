package awsgreengrass


// A reference to a DeviceDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceDefinitionReference := &DeviceDefinitionReference{
//   	DeviceDefinitionArn: jsii.String("deviceDefinitionArn"),
//   	DeviceDefinitionId: jsii.String("deviceDefinitionId"),
//   }
//
type DeviceDefinitionReference struct {
	// The ARN of the DeviceDefinition resource.
	DeviceDefinitionArn *string `field:"required" json:"deviceDefinitionArn" yaml:"deviceDefinitionArn"`
	// The Id of the DeviceDefinition resource.
	DeviceDefinitionId *string `field:"required" json:"deviceDefinitionId" yaml:"deviceDefinitionId"`
}

