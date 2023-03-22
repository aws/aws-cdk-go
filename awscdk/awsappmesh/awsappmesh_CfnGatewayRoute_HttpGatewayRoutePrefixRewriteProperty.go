package awsappmesh


// An object representing the beginning characters of the route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRoutePrefixRewriteProperty := &httpGatewayRoutePrefixRewriteProperty{
//   	defaultPrefix: jsii.String("defaultPrefix"),
//   	value: jsii.String("value"),
//   }
//
type CfnGatewayRoute_HttpGatewayRoutePrefixRewriteProperty struct {
	// The default prefix used to replace the incoming route prefix when rewritten.
	DefaultPrefix *string `field:"optional" json:"defaultPrefix" yaml:"defaultPrefix"`
	// The value used to replace the incoming route prefix when rewritten.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

