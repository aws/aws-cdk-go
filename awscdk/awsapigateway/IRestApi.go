package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

type IRestApi interface {
	awscdk.IResource
	// Gets the "execute-api" ARN.
	//
	// Returns: The "execute-api" ARN.
	// Default: "*" returns the execute API ARN for all methods/resources in
	// this API.
	//
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	// API Gateway stage that points to the latest deployment (if defined).
	DeploymentStage() Stage
	SetDeploymentStage(d Stage)
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// `undefined` when no deployment is configured.
	LatestDeployment() Deployment
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// The name of this API Gateway RestApi.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource ("/") of this API. Use it to define the API model:.
	//
	// api.root.addMethod('ANY', redirectToHomePage); // "ANY /"
	//    api.root.addResource('friends').addMethod('GET', getFriendsHandler); // "GET /friends"
	Root() IResource
}

// The jsii proxy for IRestApi
type jsiiProxy_IRestApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi)SetDeploymentStage(val Stage) {
	if err := j.validateSetDeploymentStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

func (j *jsiiProxy_IRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

