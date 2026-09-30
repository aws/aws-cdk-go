package interfacesawsiot


// A reference to a ThingPrincipalAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingPrincipalAttachmentReference := &ThingPrincipalAttachmentReference{
//   	Principal: jsii.String("principal"),
//   	ThingName: jsii.String("thingName"),
//   }
//
type ThingPrincipalAttachmentReference struct {
	// The Principal of the ThingPrincipalAttachment resource.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The ThingName of the ThingPrincipalAttachment resource.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
}

