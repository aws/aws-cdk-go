package awsrefactorspaces


// A wrapper object holding the Amazon API Gateway endpoint input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiGatewayProxyInputProperty := &ApiGatewayProxyInputProperty{
//   	EndpointType: jsii.String("endpointType"),
//   	StageName: jsii.String("stageName"),
//   }
//
type CfnApplication_ApiGatewayProxyInputProperty struct {
	// The type of endpoint to use for the API Gateway proxy.
	//
	// If no value is specified in the request, the value is set to `REGIONAL` by default.
	//
	// If the value is set to `PRIVATE` in the request, this creates a private API endpoint that is isolated from the public internet. The private endpoint can only be accessed by using Amazon Virtual Private Cloud ( Amazon VPC ) interface endpoints for the Amazon API Gateway that has been granted access. For more information about creating a private connection with Refactor Spaces and interface endpoint ( AWS PrivateLink ) availability, see [Access Refactor Spaces using an interface endpoint ( AWS PrivateLink )](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/userguide/vpc-interface-endpoints.html) .
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The name of the API Gateway stage.
	//
	// The name defaults to `prod` .
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

