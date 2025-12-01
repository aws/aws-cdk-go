package interfacesawsrtbfabric


// A reference to a OutboundExternalLink resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundExternalLinkReference := &OutboundExternalLinkReference{
//   	OutboundExternalLinkArn: jsii.String("outboundExternalLinkArn"),
//   }
//
type OutboundExternalLinkReference struct {
	// The Arn of the OutboundExternalLink resource.
	OutboundExternalLinkArn *string `field:"required" json:"outboundExternalLinkArn" yaml:"outboundExternalLinkArn"`
}

