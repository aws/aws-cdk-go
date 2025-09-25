package awsiot


// A reference to a ThingPrincipalAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingPrincipalAttachmentReference := &ThingPrincipalAttachmentReference{
//   	ThingPrincipalAttachmentId: jsii.String("thingPrincipalAttachmentId"),
//   }
//
type ThingPrincipalAttachmentReference struct {
	// The Id of the ThingPrincipalAttachment resource.
	ThingPrincipalAttachmentId *string `field:"required" json:"thingPrincipalAttachmentId" yaml:"thingPrincipalAttachmentId"`
}

