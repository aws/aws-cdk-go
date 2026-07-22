package interfacesawsoutposts


// A reference to a Site resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   siteReference := &SiteReference{
//   	SiteArn: jsii.String("siteArn"),
//   }
//
type SiteReference struct {
	// The SiteArn of the Site resource.
	SiteArn *string `field:"required" json:"siteArn" yaml:"siteArn"`
}

