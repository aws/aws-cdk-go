package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &CfnServiceProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EndpointType: jsii.String("endpointType"),
//   	LambdaEndpoint: &LambdaEndpointInputProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UrlEndpoint: &UrlEndpointInputProperty{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		HealthUrl: jsii.String("healthUrl"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
type CfnServiceProps struct {
	// The unique identifier of the application.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The unique identifier of the environment.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// A description of the service.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The endpoint type of the service.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// A summary of the configuration for the AWS Lambda endpoint type.
	LambdaEndpoint interface{} `field:"optional" json:"lambdaEndpoint" yaml:"lambdaEndpoint"`
	// The name of the service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the service.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The summary of the configuration for the URL endpoint type.
	UrlEndpoint interface{} `field:"optional" json:"urlEndpoint" yaml:"urlEndpoint"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

