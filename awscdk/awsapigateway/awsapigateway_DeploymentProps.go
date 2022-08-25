package awsapigateway


// Example:
//   // production stage
//   prdLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(prdLogGroup),
//   		accessLogFormat: apigateway.accessLogFormat.jsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   	api: api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &stageProps{
//   	deployment: deployment,
//   	accessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	accessLogFormat: apigateway.*accessLogFormat.jsonWithStandardFields(&jsonWithStandardFieldProps{
//   		caller: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(true),
//   		ip: jsii.Boolean(true),
//   		protocol: jsii.Boolean(true),
//   		requestTime: jsii.Boolean(true),
//   		resourcePath: jsii.Boolean(true),
//   		responseLength: jsii.Boolean(true),
//   		status: jsii.Boolean(true),
//   		user: jsii.Boolean(true),
//   	}),
//   })
//
type DeploymentProps struct {
	// The Rest API to deploy.
	Api IRestApi `field:"required" json:"api" yaml:"api"`
	// A description of the purpose of the API Gateway deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When an API Gateway model is updated, a new deployment will automatically be created.
	//
	// If this is true, the old API Gateway Deployment resource will not be deleted.
	// This will allow manually reverting back to a previous deployment in case for example.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
}

