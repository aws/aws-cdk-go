package awsapigateway


// Container for defining throttling parameters to API stages or methods.
//
// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.root.addResource(jsii.String("v1"))
//   echo := v1.addResource(jsii.String("echo"))
//   echoMethod := echo.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	apiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.addUsagePlan(jsii.String("UsagePlan"), &usagePlanProps{
//   	name: jsii.String("Easy"),
//   	throttle: &throttleSettings{
//   		rateLimit: jsii.Number(10),
//   		burstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.addApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type ThrottleSettings struct {
	// The maximum API request rate limit over a time ranging from one to a few seconds.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API request steady-state rate limit (average requests per second over an extended period of time).
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

