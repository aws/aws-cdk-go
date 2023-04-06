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
//   deploymentIoTJobConfigurationProperty := &DeploymentIoTJobConfigurationProperty{
//   	AbortConfig: &IoTJobAbortConfigProperty{
//   		CriteriaList: []interface{}{
//   			&IoTJobAbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: &IoTJobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &IoTJobExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: rateIncreaseCriteria,
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	TimeoutConfig: &IoTJobTimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
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

