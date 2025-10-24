package awsapigateway


// Represents the API stages that a usage plan applies to.
//
// Example:
//   var plan UsagePlan
//   var api RestApi
//   var echoMethod Method
//
//
//   plan.AddApiStage(&UsagePlanPerApiStage{
//   	Stage: api.DeploymentStage,
//   	Throttle: []ThrottlingPerMethod{
//   		&ThrottlingPerMethod{
//   			Method: echoMethod,
//   			Throttle: &ThrottleSettings{
//   				RateLimit: jsii.Number(10),
//   				BurstLimit: jsii.Number(2),
//   			},
//   		},
//   	},
//   })
//
type UsagePlanPerApiStage struct {
	// Default: none.
	//
	Api IRestApi `field:"optional" json:"api" yaml:"api"`
	// [disable-awslint:ref-via-interface].
	// Default: none.
	//
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	// Default: none.
	//
	Throttle *[]*ThrottlingPerMethod `field:"optional" json:"throttle" yaml:"throttle"`
}

