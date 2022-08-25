package awsappmesh


// The type returned from the `bind()` method in {@link HttpGatewayRoutePathMatch}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRoutePathMatchConfig := &httpGatewayRoutePathMatchConfig{
//   	prefixPathMatch: jsii.String("prefixPathMatch"),
//   	prefixPathRewrite: &httpGatewayRoutePrefixRewriteProperty{
//   		defaultPrefix: jsii.String("defaultPrefix"),
//   		value: jsii.String("value"),
//   	},
//   	wholePathMatch: &httpPathMatchProperty{
//   		exact: jsii.String("exact"),
//   		regex: jsii.String("regex"),
//   	},
//   	wholePathRewrite: &httpGatewayRoutePathRewriteProperty{
//   		exact: jsii.String("exact"),
//   	},
//   }
//
type HttpGatewayRoutePathMatchConfig struct {
	// Gateway route configuration for matching on the prefix of the URL path of the request.
	PrefixPathMatch *string `field:"optional" json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Gateway route configuration for rewriting the prefix of the URL path of the request.
	PrefixPathRewrite *CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty `field:"optional" json:"prefixPathRewrite" yaml:"prefixPathRewrite"`
	// Gateway route configuration for matching on the complete URL path of the request.
	WholePathMatch *CfnGatewayRoute_HttpPathMatchProperty `field:"optional" json:"wholePathMatch" yaml:"wholePathMatch"`
	// Gateway route configuration for rewriting the complete URL path of the request..
	WholePathRewrite *CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty `field:"optional" json:"wholePathRewrite" yaml:"wholePathRewrite"`
}

