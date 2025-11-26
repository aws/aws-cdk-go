package previewawsrefactorspacesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	ApiGatewayProxy: &ApiGatewayProxyInputProperty{
//   		EndpointType: jsii.String("endpointType"),
//   		StageName: jsii.String("stageName"),
//   	},
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	ProxyType: jsii.String("proxyType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html
//
type CfnApplicationMixinProps struct {
	// The endpoint URL of the Amazon API Gateway proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-apigatewayproxy
	//
	ApiGatewayProxy interface{} `field:"optional" json:"apiGatewayProxy" yaml:"apiGatewayProxy"`
	// The unique identifier of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The proxy type of the proxy created within the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-proxytype
	//
	ProxyType *string `field:"optional" json:"proxyType" yaml:"proxyType"`
	// The tags assigned to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-application.html#cfn-refactorspaces-application-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

