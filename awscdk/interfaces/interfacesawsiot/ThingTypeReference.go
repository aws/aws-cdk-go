package interfacesawsiot


// A reference to a ThingType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingTypeReference := &ThingTypeReference{
//   	ThingTypeArn: jsii.String("thingTypeArn"),
//   	ThingTypeName: jsii.String("thingTypeName"),
//   }
//
type ThingTypeReference struct {
	// The ARN of the ThingType resource.
	ThingTypeArn *string `field:"required" json:"thingTypeArn" yaml:"thingTypeArn"`
	// The ThingTypeName of the ThingType resource.
	ThingTypeName *string `field:"required" json:"thingTypeName" yaml:"thingTypeName"`
}

