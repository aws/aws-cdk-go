package interfacesawsapigateway


// A reference to a BasePathMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basePathMappingReference := &BasePathMappingReference{
//   	BasePath: jsii.String("basePath"),
//   	DomainName: jsii.String("domainName"),
//   }
//
type BasePathMappingReference struct {
	// The BasePath of the BasePathMapping resource.
	BasePath *string `field:"required" json:"basePath" yaml:"basePath"`
	// The DomainName of the BasePathMapping resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

