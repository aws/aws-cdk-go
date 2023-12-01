package awscdkapigatewayv2integrationsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

// Properties to initialize `HttpServiceDiscoveryIntegration`.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
//   	Vpc: Vpc,
//   })
//   namespace := servicediscovery.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
//   	Name: jsii.String("boobar.com"),
//   	Vpc: Vpc,
//   })
//   service := namespace.CreateService(jsii.String("Service"))
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &HttpServiceDiscoveryIntegrationProps{
//   		VpcLink: *VpcLink,
//   	}),
//   })
//
// Deprecated.
type HttpServiceDiscoveryIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Default: HttpMethod.ANY
	//
	// Deprecated.
	Method awscdkapigatewayv2alpha.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	// Deprecated.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	// Deprecated.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Default: - a new VpcLink is created.
	//
	// Deprecated.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

