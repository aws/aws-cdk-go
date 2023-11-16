package awsapigateway


// Example:
//   // production stage
//   prodLogGroup := logs.NewLogGroup(this, jsii.String("PrdLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: apigateway.NewLogGroupLogDestination(prodLogGroup),
//   		AccessLogFormat: apigateway.AccessLogFormat_JsonWithStandardFields(),
//   	},
//   })
//   deployment := apigateway.NewDeployment(this, jsii.String("Deployment"), &DeploymentProps{
//   	Api: Api,
//   })
//
//   // development stage
//   devLogGroup := logs.NewLogGroup(this, jsii.String("DevLogs"))
//   apigateway.NewStage(this, jsii.String("dev"), &StageProps{
//   	Deployment: Deployment,
//   	AccessLogDestination: apigateway.NewLogGroupLogDestination(devLogGroup),
//   	AccessLogFormat: apigateway.AccessLogFormat_*JsonWithStandardFields(&JsonWithStandardFieldProps{
//   		Caller: jsii.Boolean(false),
//   		HttpMethod: jsii.Boolean(true),
//   		Ip: jsii.Boolean(true),
//   		Protocol: jsii.Boolean(true),
//   		RequestTime: jsii.Boolean(true),
//   		ResourcePath: jsii.Boolean(true),
//   		ResponseLength: jsii.Boolean(true),
//   		Status: jsii.Boolean(true),
//   		User: jsii.Boolean(true),
//   	}),
//   })
//
type DeploymentProps struct {
	// The Rest API to deploy.
	Api IRestApi `field:"required" json:"api" yaml:"api"`
	// A description of the purpose of the API Gateway deployment.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When an API Gateway model is updated, a new deployment will automatically be created.
	//
	// If this is true, the old API Gateway Deployment resource will not be deleted.
	// This will allow manually reverting back to a previous deployment in case for example.
	// Default: false.
	//
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
}

