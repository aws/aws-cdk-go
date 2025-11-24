package mixinsawsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceMixinProps := &CfnServiceMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	Description: jsii.String("description"),
//   	EndpointType: jsii.String("endpointType"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	LambdaEndpoint: &LambdaEndpointInputProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UrlEndpoint: &UrlEndpointInputProperty{
//   		HealthUrl: jsii.String("healthUrl"),
//   		Url: jsii.String("url"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html
//
type CfnServiceMixinProps struct {
	// The unique identifier of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-applicationidentifier
	//
	ApplicationIdentifier *string `field:"optional" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// A description of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The endpoint type of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-endpointtype
	//
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The unique identifier of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// A summary of the configuration for the AWS Lambda endpoint type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-lambdaendpoint
	//
	LambdaEndpoint interface{} `field:"optional" json:"lambdaEndpoint" yaml:"lambdaEndpoint"`
	// The name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The summary of the configuration for the URL endpoint type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-urlendpoint
	//
	UrlEndpoint interface{} `field:"optional" json:"urlEndpoint" yaml:"urlEndpoint"`
	// The ID of the virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-refactorspaces-service.html#cfn-refactorspaces-service-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

