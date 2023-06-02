package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	ProxyType: jsii.String("proxyType"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	ApiGatewayProxy: &ApiGatewayProxyInputProperty{
//   		EndpointType: jsii.String("endpointType"),
//   		StageName: jsii.String("stageName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// The unique identifier of the environment.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The proxy type of the proxy created within the application.
	ProxyType *string `field:"required" json:"proxyType" yaml:"proxyType"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The endpoint URL of the Amazon API Gateway proxy.
	ApiGatewayProxy interface{} `field:"optional" json:"apiGatewayProxy" yaml:"apiGatewayProxy"`
	// The tags assigned to the application.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

