package awscdkapigatewayv2alpha


// Properties used to create the ApiMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var api iApi
//   var domainName domainName
//   var stage iStage
//
//   apiMappingProps := &ApiMappingProps{
//   	Api: api,
//   	DomainName: domainName,
//
//   	// the properties below are optional
//   	ApiMappingKey: jsii.String("apiMappingKey"),
//   	Stage: stage,
//   }
//
// Experimental.
type ApiMappingProps struct {
	// The Api to which this mapping is applied.
	// Experimental.
	Api IApi `field:"required" json:"api" yaml:"api"`
	// custom domain name of the mapping target.
	// Experimental.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// Api mapping key.
	//
	// The path where this stage should be mapped to on the domain.
	// Experimental.
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
	// stage for the ApiMapping resource required for WebSocket API defaults to default stage of an HTTP API.
	// Experimental.
	Stage IStage `field:"optional" json:"stage" yaml:"stage"`
}

