package awsapigateway


// The `EndpointConfiguration` property type specifies the endpoint types of an Amazon API Gateway domain name.
//
// `EndpointConfiguration` is a property of the [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
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
//   }
//
type CfnDomainName_EndpointConfigurationProperty struct {
	// A list of endpoint types of an API or its custom domain name.
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is `EDGE` . For a regional API and its custom domain name, the endpoint type is `REGIONAL` .
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

