package awsapigateway


// Properties for controlling items output in JSON standard format.
//
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
type JsonWithStandardFieldProps struct {
	// If this flag is enabled, the principal identifier of the caller will be output to the log.
	Caller *bool `field:"required" json:"caller" yaml:"caller"`
	// If this flag is enabled, the http method will be output to the log.
	HttpMethod *bool `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// If this flag is enabled, the source IP of request will be output to the log.
	Ip *bool `field:"required" json:"ip" yaml:"ip"`
	// If this flag is enabled, the request protocol will be output to the log.
	Protocol *bool `field:"required" json:"protocol" yaml:"protocol"`
	// If this flag is enabled, the CLF-formatted request time((dd/MMM/yyyy:HH:mm:ss +-hhmm) will be output to the log.
	RequestTime *bool `field:"required" json:"requestTime" yaml:"requestTime"`
	// If this flag is enabled, the path to your resource will be output to the log.
	ResourcePath *bool `field:"required" json:"resourcePath" yaml:"resourcePath"`
	// If this flag is enabled, the response payload length will be output to the log.
	ResponseLength *bool `field:"required" json:"responseLength" yaml:"responseLength"`
	// If this flag is enabled, the method response status will be output to the log.
	Status *bool `field:"required" json:"status" yaml:"status"`
	// If this flag is enabled, the principal identifier of the user will be output to the log.
	User *bool `field:"required" json:"user" yaml:"user"`
}

