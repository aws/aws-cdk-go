package awsapigatewayv2


// Options for DomainMapping.
//
// Example:
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var handler function
//   var dn domainName
//
//
//   apiDemo := apigwv2.NewHttpApi(this, jsii.String("DemoApi"), &HttpApiProps{
//   	DefaultIntegration: awscdklibawsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
//   	// https://${dn.domainName}/demo goes to apiDemo $default stage
//   	DefaultDomainMapping: &DomainMappingOptions{
//   		DomainName: dn,
//   		MappingKey: jsii.String("demo"),
//   	},
//   })
//
type DomainMappingOptions struct {
	// The domain name for the mapping.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The API mapping key.
	//
	// Leave it undefined for the root path mapping.
	// Default: - empty key for the root path mapping.
	//
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
}

