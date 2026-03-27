package awsroute53globalresolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAccessTokenPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAccessTokenMixinProps := &CfnAccessTokenMixinProps{
//   	ClientToken: jsii.String("clientToken"),
//   	DnsViewId: jsii.String("dnsViewId"),
//   	ExpiresAt: jsii.String("expiresAt"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html
//
type CfnAccessTokenMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html#cfn-route53globalresolver-accesstoken-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html#cfn-route53globalresolver-accesstoken-dnsviewid
	//
	DnsViewId *string `field:"optional" json:"dnsViewId" yaml:"dnsViewId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html#cfn-route53globalresolver-accesstoken-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html#cfn-route53globalresolver-accesstoken-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-accesstoken.html#cfn-route53globalresolver-accesstoken-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

