package interfacesawsservicecatalogappregistry


// A reference to a AttributeGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeGroupReference := &AttributeGroupReference{
//   	AttributeGroupArn: jsii.String("attributeGroupArn"),
//   	AttributeGroupId: jsii.String("attributeGroupId"),
//   }
//
type AttributeGroupReference struct {
	// The ARN of the AttributeGroup resource.
	AttributeGroupArn *string `field:"required" json:"attributeGroupArn" yaml:"attributeGroupArn"`
	// The Id of the AttributeGroup resource.
	AttributeGroupId *string `field:"required" json:"attributeGroupId" yaml:"attributeGroupId"`
}

