package awsgreengrassv2


// Contains information about an AWS IoT job configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rateIncreaseCriteria interface{}
//
//   deploymentIoTJobConfigurationProperty := &deploymentIoTJobConfigurationProperty{
//   	abortConfig: &ioTJobAbortConfigProperty{
//   		criteriaList: []interface{}{
//   			&ioTJobAbortCriteriaProperty{
//   				action: jsii.String("action"),
//   				failureType: jsii.String("failureType"),
//   				minNumberOfExecutedThings: jsii.Number(123),
//   				thresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	jobExecutionsRolloutConfig: &ioTJobExecutionsRolloutConfigProperty{
//   		exponentialRate: &ioTJobExponentialRolloutRateProperty{
//   			baseRatePerMinute: jsii.Number(123),
//   			incrementFactor: jsii.Number(123),
//   			rateIncreaseCriteria: rateIncreaseCriteria,
//   		},
//   		maximumPerMinute: jsii.Number(123),
//   	},
//   	timeoutConfig: &ioTJobTimeoutConfigProperty{
//   		inProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnDeployment_DeploymentIoTJobConfigurationProperty struct {
	// The stop configuration for the job.
	//
	// This configuration defines when and how to stop a job rollout.
	AbortConfig interface{} `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// The rollout configuration for the job.
	//
	// This configuration defines the rate at which the job rolls out to the fleet of target devices.
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// The timeout configuration for the job.
	//
	// This configuration defines the amount of time each device has to complete the job.
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

