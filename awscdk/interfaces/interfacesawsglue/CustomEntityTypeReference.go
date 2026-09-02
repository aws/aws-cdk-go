package interfacesawsglue


// A reference to a CustomEntityType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customEntityTypeReference := &CustomEntityTypeReference{
//   	CustomEntityTypeName: jsii.String("customEntityTypeName"),
//   }
//
type CustomEntityTypeReference struct {
	// The Name of the CustomEntityType resource.
	CustomEntityTypeName *string `field:"required" json:"customEntityTypeName" yaml:"customEntityTypeName"`
}

