package previewawsappmeshmixins


// An object that represents the path to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpGatewayRoutePathRewriteProperty := &HttpGatewayRoutePathRewriteProperty{
//   	Exact: jsii.String("exact"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutepathrewrite.html
//
type CfnGatewayRoutePropsMixin_HttpGatewayRoutePathRewriteProperty struct {
	// The exact path to rewrite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutepathrewrite.html#cfn-appmesh-gatewayroute-httpgatewayroutepathrewrite-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

