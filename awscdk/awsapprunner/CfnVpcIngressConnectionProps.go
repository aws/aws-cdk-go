package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVpcIngressConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcIngressConnectionProps := &CfnVpcIngressConnectionProps{
//   	IngressVpcConfiguration: &IngressVpcConfigurationProperty{
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	ServiceArn: jsii.String("serviceArn"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcIngressConnectionName: jsii.String("vpcIngressConnectionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcingressconnection.html
//
type CfnVpcIngressConnectionProps struct {
	// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcingressconnection.html#cfn-apprunner-vpcingressconnection-ingressvpcconfiguration
	//
	IngressVpcConfiguration interface{} `field:"required" json:"ingressVpcConfiguration" yaml:"ingressVpcConfiguration"`
	// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcingressconnection.html#cfn-apprunner-vpcingressconnection-servicearn
	//
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// An optional list of metadata items that you can associate with the VPC Ingress Connection resource.
	//
	// A tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcingressconnection.html#cfn-apprunner-vpcingressconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The customer-provided VPC Ingress Connection name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcingressconnection.html#cfn-apprunner-vpcingressconnection-vpcingressconnectionname
	//
	VpcIngressConnectionName *string `field:"optional" json:"vpcIngressConnectionName" yaml:"vpcIngressConnectionName"`
}

