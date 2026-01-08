package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
)

// Options for DomainMapping.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var handler Function
//   var dn DomainName
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
type DomainMappingOptions struct {
	// The domain name for the mapping.
	DomainName interfacesawsapigatewayv2.IDomainNameRef `field:"required" json:"domainName" yaml:"domainName"`
	// The API mapping key.
	//
	// Leave it undefined for the root path mapping.
	// Default: - empty key for the root path mapping.
	//
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
}

