package mixinsawsappmesh


// An object that represents the query parameter in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryParameterProperty := &QueryParameterProperty{
//   	Match: &HttpQueryParameterMatchProperty{
//   		Exact: jsii.String("exact"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html
//
type CfnGatewayRoutePropsMixin_QueryParameterProperty struct {
	// The query parameter to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// A name for the query parameter that will be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

