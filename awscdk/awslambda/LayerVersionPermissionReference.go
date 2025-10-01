package awslambda


// A reference to a LayerVersionPermission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layerVersionPermissionReference := &LayerVersionPermissionReference{
//   	LayerVersionPermissionId: jsii.String("layerVersionPermissionId"),
//   }
//
type LayerVersionPermissionReference struct {
	// The Id of the LayerVersionPermission resource.
	LayerVersionPermissionId *string `field:"required" json:"layerVersionPermissionId" yaml:"layerVersionPermissionId"`
}

