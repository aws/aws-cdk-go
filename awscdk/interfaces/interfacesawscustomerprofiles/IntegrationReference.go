package interfacesawscustomerprofiles


// A reference to a Integration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationReference := &IntegrationReference{
//   	DomainName: jsii.String("domainName"),
//   	Uri: jsii.String("uri"),
//   }
//
type IntegrationReference struct {
	// The DomainName of the Integration resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Uri of the Integration resource.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

