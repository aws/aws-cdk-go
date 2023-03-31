package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	apiGatewayProxy: &apiGatewayProxyInputProperty{
//   		endpointType: jsii.String("endpointType"),
//   		stageName: jsii.String("stageName"),
//   	},
//   	environmentIdentifier: jsii.String("environmentIdentifier"),
//   	name: jsii.String("name"),
//   	proxyType: jsii.String("proxyType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnApplicationProps struct {
	// The endpoint URL of the Amazon API Gateway proxy.
	ApiGatewayProxy interface{} `field:"optional" json:"apiGatewayProxy" yaml:"apiGatewayProxy"`
	// The unique identifier of the environment.
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the application.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The proxy type of the proxy created within the application.
	ProxyType *string `field:"optional" json:"proxyType" yaml:"proxyType"`
	// The tags assigned to the application.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

