package awsnetworkmanager


// A reference to a Site resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   siteReference := &SiteReference{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	SiteArn: jsii.String("siteArn"),
//   	SiteId: jsii.String("siteId"),
//   }
//
type SiteReference struct {
	// The GlobalNetworkId of the Site resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ARN of the Site resource.
	SiteArn *string `field:"required" json:"siteArn" yaml:"siteArn"`
	// The SiteId of the Site resource.
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
}

