package interfacesawsrtbfabric


// A reference to a InboundExternalLink resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inboundExternalLinkReference := &InboundExternalLinkReference{
//   	InboundExternalLinkArn: jsii.String("inboundExternalLinkArn"),
//   }
//
type InboundExternalLinkReference struct {
	// The Arn of the InboundExternalLink resource.
	InboundExternalLinkArn *string `field:"required" json:"inboundExternalLinkArn" yaml:"inboundExternalLinkArn"`
}

