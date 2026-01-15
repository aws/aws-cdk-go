package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
)

// Properties used to create the ApiMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiRef IApiRef
//   var domainNameRef IDomainNameRef
//   var stage IStage
//
//   apiMappingProps := &ApiMappingProps{
//   	Api: apiRef,
//   	DomainName: domainNameRef,
//
//   	// the properties below are optional
//   	ApiMappingKey: jsii.String("apiMappingKey"),
//   	Stage: stage,
//   }
//
type ApiMappingProps struct {
	// The Api to which this mapping is applied.
	Api interfacesawsapigatewayv2.IApiRef `field:"required" json:"api" yaml:"api"`
	// custom domain name of the mapping target.
	DomainName interfacesawsapigatewayv2.IDomainNameRef `field:"required" json:"domainName" yaml:"domainName"`
	// Api mapping key.
	//
	// The path where this stage should be mapped to on the domain.
	// Default: - undefined for the root path mapping.
	//
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
	// stage for the ApiMapping resource required for WebSocket API defaults to default stage of an HTTP API.
	// Default: - Default stage of the passed API for HTTP API, required for WebSocket API.
	//
	Stage IStage `field:"optional" json:"stage" yaml:"stage"`
}

