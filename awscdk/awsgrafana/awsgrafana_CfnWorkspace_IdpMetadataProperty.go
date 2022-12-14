package awsgrafana


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
	// `CfnWorkspace.IdpMetadataProperty.Url`.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// `CfnWorkspace.IdpMetadataProperty.Xml`.
	Xml *string `field:"optional" json:"xml" yaml:"xml"`
}

