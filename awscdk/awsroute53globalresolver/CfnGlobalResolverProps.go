package awsroute53globalresolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGlobalResolver`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGlobalResolverProps := &CfnGlobalResolverProps{
//   	Name: jsii.String("name"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	ObservabilityRegion: jsii.String("observabilityRegion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html
//
type CfnGlobalResolverProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of regions the Global Resolver will exist in.
	//
	// This list cannot be updated and will stay fixed for the duration of the Global Resolver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-regions
	//
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-observabilityregion
	//
	ObservabilityRegion *string `field:"optional" json:"observabilityRegion" yaml:"observabilityRegion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html#cfn-route53globalresolver-globalresolver-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

