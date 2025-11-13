package interfacesawsapigateway


// A reference to a VpcLink resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcLinkReference := &VpcLinkReference{
//   	VpcLinkId: jsii.String("vpcLinkId"),
//   }
//
type VpcLinkReference struct {
	// The VpcLinkId of the VpcLink resource.
	VpcLinkId *string `field:"required" json:"vpcLinkId" yaml:"vpcLinkId"`
}

