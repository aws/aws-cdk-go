package awsgrafana


// A structure containing the identity provider (IdP) metadata used to integrate the identity provider with this workspace.
//
// You can specify the metadata either by providing a URL to its location in the `url` parameter, or by specifying the full metadata in XML format in the `xml` parameter. Specifying both will cause an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idpMetadataProperty := &idpMetadataProperty{
//   	url: jsii.String("url"),
//   	xml: jsii.String("xml"),
//   }
//
type CfnWorkspace_IdpMetadataProperty struct {
	// The URL of the location containing the IdP metadata.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The full IdP metadata, in XML format.
	Xml *string `field:"optional" json:"xml" yaml:"xml"`
}

