package awsapigateway


// Options for creating an api mapping.
//
// Example:
//   var acmCertificateForExampleCom interface{}
//   var restApi restApi
//   var secondRestApi restApi
//
//
//   domain := apigateway.NewDomainName(this, jsii.String("custom-domain"), &DomainNameProps{
//   	DomainName: jsii.String("example.com"),
//   	Certificate: acmCertificateForExampleCom,
//   	Mapping: restApi,
//   })
//
//   domain.AddApiMapping(secondRestApi.DeploymentStage, &ApiMappingOptions{
//   	BasePath: jsii.String("orders/v2/api"),
//   })
//
type ApiMappingOptions struct {
	// The api path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	//
	// If this is undefined, a mapping will be added for the empty path. Any request
	// that does not match a mapping will get sent to the API that has been mapped
	// to the empty path.
	// Default: - map requests from the domain root (e.g. `example.com`).
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
}

