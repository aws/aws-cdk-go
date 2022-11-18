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
	// `AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration`.
	IngressVpcConfiguration interface{} `field:"required" json:"ingressVpcConfiguration" yaml:"ingressVpcConfiguration"`
	// `AWS::AppRunner::VpcIngressConnection.ServiceArn`.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// `AWS::AppRunner::VpcIngressConnection.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::AppRunner::VpcIngressConnection.VpcIngressConnectionName`.
	VpcIngressConnectionName *string `field:"optional" json:"vpcIngressConnectionName" yaml:"vpcIngressConnectionName"`
}

