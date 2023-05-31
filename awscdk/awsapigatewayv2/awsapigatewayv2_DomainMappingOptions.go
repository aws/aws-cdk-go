package awsapigatewayv2


// Options for DomainMapping.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var handler function
//   var dn domainName
//
//
//   apiDemo := apigwv2.NewHttpApi(this, jsii.String("DemoApi"), &HttpApiProps{
//   	DefaultIntegration: awscdk.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
//   	// https://${dn.domainName}/demo goes to apiDemo $default stage
//   	DefaultDomainMapping: &DomainMappingOptions{
//   		DomainName: dn,
//   		MappingKey: jsii.String("demo"),
//   	},
//   })
//
// Experimental.
type DomainMappingOptions struct {
	// The domain name for the mapping.
	// Experimental.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The API mapping key.
	//
	// Leave it undefined for the root path mapping.
	// Experimental.
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
}

