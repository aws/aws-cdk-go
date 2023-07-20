package awsapigateway


// The `EndpointConfiguration` property type specifies the endpoint types of a REST API.
//
// `EndpointConfiguration` is a property of the [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &EndpointConfigurationProperty{
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   	VpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
//
type CfnRestApi_EndpointConfigurationProperty struct {
	// A list of endpoint types of an API (RestApi) or its custom domain name (DomainName).
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is `"EDGE"` . For a regional API and its custom domain name, the endpoint type is `REGIONAL` . For a private API, the endpoint type is `PRIVATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes.
	//
	// It is only supported for `PRIVATE` endpoint type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-vpcendpointids
	//
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

