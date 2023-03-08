package awsappmesh


// An object that represents the path to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRoutePathRewriteProperty := &httpGatewayRoutePathRewriteProperty{
//   	exact: jsii.String("exact"),
//   }
//
type CfnGatewayRoute_HttpGatewayRoutePathRewriteProperty struct {
	// The exact path to rewrite.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

