package mixinsawsappmesh


// An object representing the query parameter to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpQueryParameterMatchProperty := &HttpQueryParameterMatchProperty{
//   	Exact: jsii.String("exact"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpqueryparametermatch.html
//
type CfnGatewayRoutePropsMixin_HttpQueryParameterMatchProperty struct {
	// The exact query parameter to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpqueryparametermatch.html#cfn-appmesh-gatewayroute-httpqueryparametermatch-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

