package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Properties to initialize `HttpServiceDiscoveryIntegration`.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	DefaultIntegration: awscdk.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &HttpServiceDiscoveryIntegrationProps{
//   		VpcLink: *VpcLink,
//   	}),
//   })
//
type HttpServiceDiscoveryIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Default: HttpMethod.ANY
	//
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Default: - a new VpcLink is created.
	//
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

