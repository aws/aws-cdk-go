package awsroute53globalresolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccessSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessSourceProps := &CfnAccessSourceProps{
//   	DnsViewId: jsii.String("dnsViewId"),
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	Cidr: jsii.String("cidr"),
//   	ClientToken: jsii.String("clientToken"),
//   	CoreNetworkEndpointId: jsii.String("coreNetworkEndpointId"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Name: jsii.String("name"),
//   	Segment: jsii.String("segment"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html
//
type CfnAccessSourceProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-dnsviewid
	//
	DnsViewId *string `field:"required" json:"dnsViewId" yaml:"dnsViewId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-corenetworkendpointid
	//
	CoreNetworkEndpointId *string `field:"optional" json:"coreNetworkEndpointId" yaml:"coreNetworkEndpointId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-segment
	//
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesssource.html#cfn-route53globalresolver-accesssource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

