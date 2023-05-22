package awsapigateway


// Represents the API stages that a usage plan applies to.
//
// Example:
//   var plan usagePlan
//   var api restApi
//   var echoMethod method
//
//
//   plan.AddApiStage(&UsagePlanPerApiStage{
//   	Stage: api.DeploymentStage,
//   	Throttle: []throttlingPerMethod{
//   		&throttlingPerMethod{
//   			Method: echoMethod,
//   			Throttle: &ThrottleSettings{
//   				RateLimit: jsii.Number(10),
//   				BurstLimit: jsii.Number(2),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type UsagePlanPerApiStage struct {
	// Experimental.
	Api IRestApi `field:"optional" json:"api" yaml:"api"`
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	// Experimental.
	Throttle *[]*ThrottlingPerMethod `field:"optional" json:"throttle" yaml:"throttle"`
}

