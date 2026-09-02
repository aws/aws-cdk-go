package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPortalProductPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPortalProductMixinProps := &CfnPortalProductMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-portalproduct.html
//
type CfnPortalProductMixinProps struct {
	// A description of the portal product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-portalproduct.html#cfn-apigatewayv2-portalproduct-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the portal product as it appears in a published portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-portalproduct.html#cfn-apigatewayv2-portalproduct-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The collection of tags associated with the portal product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-portalproduct.html#cfn-apigatewayv2-portalproduct-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

