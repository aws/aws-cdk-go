package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for backend integrations for an API Gateway method.
//
// Use one of the concrete classes such as `MockIntegration`, `AwsIntegration`, `LambdaIntegration`
// or implement on your own by specifying the set of props.
//
// Example:
//   var books resource
//   var iamUser user
//
//
//   getBooks := books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizationType: apigateway.authorizationType_IAM,
//   })
//
//   iamUser.attachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowBooks"), &policyProps{
//   	statements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   			actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			effect: iam.effect_ALLOW,
//   			resources: []*string{
//   				getBooks.methodArn,
//   			},
//   		}),
//   	},
//   }))
//
type Integration interface {
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for Integration
type jsiiProxy_Integration struct {
	_ byte // padding
}

func NewIntegration(props *IntegrationProps) Integration {
	_init_.Initialize()

	j := jsiiProxy_Integration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Integration",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewIntegration_Override(i Integration, props *IntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.Integration",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_Integration) Bind(_method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

