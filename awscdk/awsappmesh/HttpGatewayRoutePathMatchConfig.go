package awsappmesh


// The type returned from the `bind()` method in `HttpGatewayRoutePathMatch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRoutePathMatchConfig := &HttpGatewayRoutePathMatchConfig{
//   	PrefixPathMatch: jsii.String("prefixPathMatch"),
//   	PrefixPathRewrite: &HttpGatewayRoutePrefixRewriteProperty{
//   		DefaultPrefix: jsii.String("defaultPrefix"),
//   		Value: jsii.String("value"),
//   	},
//   	WholePathMatch: &HttpPathMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Regex: jsii.String("regex"),
//   	},
//   	WholePathRewrite: &HttpGatewayRoutePathRewriteProperty{
//   		Exact: jsii.String("exact"),
//   	},
//   }
//
type HttpGatewayRoutePathMatchConfig struct {
	// Gateway route configuration for matching on the prefix of the URL path of the request.
	// Default: - no matching will be performed on the prefix of the URL path.
	//
	PrefixPathMatch *string `field:"optional" json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Gateway route configuration for rewriting the prefix of the URL path of the request.
	// Default: - rewrites the request's URL path to '/'.
	//
	PrefixPathRewrite *CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty `field:"optional" json:"prefixPathRewrite" yaml:"prefixPathRewrite"`
	// Gateway route configuration for matching on the complete URL path of the request.
	// Default: - no matching will be performed on the complete URL path.
	//
	WholePathMatch *CfnGatewayRoute_HttpPathMatchProperty `field:"optional" json:"wholePathMatch" yaml:"wholePathMatch"`
	// Gateway route configuration for rewriting the complete URL path of the request..
	// Default: - no rewrite will be performed on the request's complete URL path.
	//
	WholePathRewrite *CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty `field:"optional" json:"wholePathRewrite" yaml:"wholePathRewrite"`
}

