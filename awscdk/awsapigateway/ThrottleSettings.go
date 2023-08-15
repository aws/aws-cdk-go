package awsapigateway


// Container for defining throttling parameters to API stages or methods.
//
// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.Root.AddResource(jsii.String("v1"))
//   echo := v1.AddResource(jsii.String("echo"))
//   echoMethod := echo.AddMethod(jsii.String("GET"), integration, &MethodOptions{
//   	ApiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.AddUsagePlan(jsii.String("UsagePlan"), &UsagePlanProps{
//   	Name: jsii.String("Easy"),
//   	Throttle: &ThrottleSettings{
//   		RateLimit: jsii.Number(10),
//   		BurstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.AddApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type ThrottleSettings struct {
	// The maximum API request rate limit over a time ranging from one to a few seconds.
	// Default: none.
	//
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API request steady-state rate limit (average requests per second over an extended period of time).
	// Default: none.
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

