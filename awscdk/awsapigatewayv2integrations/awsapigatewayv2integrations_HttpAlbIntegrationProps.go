package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
)

// Properties to initialize `HttpAlbIntegration`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().appendHeader(jsii.String("header2"), apigwv2.mappingValue.requestHeader(jsii.String("header1"))).removeHeader(jsii.String("header1")),
//   	}),
//   })
//
// Experimental.
type HttpAlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

