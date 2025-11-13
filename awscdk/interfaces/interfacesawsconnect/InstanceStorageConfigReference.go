package interfacesawsconnect


// A reference to a InstanceStorageConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceStorageConfigReference := &InstanceStorageConfigReference{
//   	AssociationId: jsii.String("associationId"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
type InstanceStorageConfigReference struct {
	// The AssociationId of the InstanceStorageConfig resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
	// The InstanceArn of the InstanceStorageConfig resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The ResourceType of the InstanceStorageConfig resource.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

