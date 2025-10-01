package awsapigatewayv2


// A reference to a ApiMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiMappingReference := &ApiMappingReference{
//   	ApiMappingId: jsii.String("apiMappingId"),
//   	DomainName: jsii.String("domainName"),
//   }
//
type ApiMappingReference struct {
	// The ApiMappingId of the ApiMapping resource.
	ApiMappingId *string `field:"required" json:"apiMappingId" yaml:"apiMappingId"`
	// The DomainName of the ApiMapping resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

