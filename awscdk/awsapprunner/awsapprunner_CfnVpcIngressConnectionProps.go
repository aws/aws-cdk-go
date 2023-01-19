package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVpcIngressConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcIngressConnectionProps := &cfnVpcIngressConnectionProps{
//   	ingressVpcConfiguration: &ingressVpcConfigurationProperty{
//   		vpcEndpointId: jsii.String("vpcEndpointId"),
//   		vpcId: jsii.String("vpcId"),
//   	},
//   	serviceArn: jsii.String("serviceArn"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcIngressConnectionName: jsii.String("vpcIngressConnectionName"),
//   }
//
type CfnVpcIngressConnectionProps struct {
	// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource.
	IngressVpcConfiguration interface{} `field:"required" json:"ingressVpcConfiguration" yaml:"ingressVpcConfiguration"`
	// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// An optional list of metadata items that you can associate with the VPC Ingress Connection resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The customer-provided VPC Ingress Connection name.
	VpcIngressConnectionName *string `field:"optional" json:"vpcIngressConnectionName" yaml:"vpcIngressConnectionName"`
}

