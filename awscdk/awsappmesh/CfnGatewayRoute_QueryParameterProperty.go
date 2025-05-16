package awsappmesh


// An object that represents the query parameter in the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryParameterProperty := &QueryParameterProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Match: &HttpQueryParameterMatchProperty{
//   		Exact: jsii.String("exact"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html
//
type CfnGatewayRoute_QueryParameterProperty struct {
	// A name for the query parameter that will be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query parameter to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

