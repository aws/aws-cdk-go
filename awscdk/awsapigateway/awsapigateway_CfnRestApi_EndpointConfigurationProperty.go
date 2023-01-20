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
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	types: []*string{
//   		jsii.String("types"),
//   	},
//   	vpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
type CfnRestApi_EndpointConfigurationProperty struct {
	// A list of endpoint types of an API or its custom domain name. Valid values include:.
	//
	// - `EDGE` : For an edge-optimized API and its custom domain name.
	// - `REGIONAL` : For a regional API and its custom domain name.
	// - `PRIVATE` : For a private API.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// A list of VPC endpoint IDs of an API ( [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) ) against which to create Route53 ALIASes. It is only supported for `PRIVATE` endpoint type.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

