package mixinsawsappmesh


// An object representing the beginning characters of the route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpGatewayRoutePrefixRewriteProperty := &HttpGatewayRoutePrefixRewriteProperty{
//   	DefaultPrefix: jsii.String("defaultPrefix"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteprefixrewrite.html
//
type CfnGatewayRoutePropsMixin_HttpGatewayRoutePrefixRewriteProperty struct {
	// The default prefix used to replace the incoming route prefix when rewritten.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteprefixrewrite.html#cfn-appmesh-gatewayroute-httpgatewayrouteprefixrewrite-defaultprefix
	//
	DefaultPrefix *string `field:"optional" json:"defaultPrefix" yaml:"defaultPrefix"`
	// The value used to replace the incoming route prefix when rewritten.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteprefixrewrite.html#cfn-appmesh-gatewayroute-httpgatewayrouteprefixrewrite-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

