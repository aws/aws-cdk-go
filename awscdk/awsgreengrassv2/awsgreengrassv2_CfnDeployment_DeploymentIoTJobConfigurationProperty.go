package awsgreengrassv2


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
	// `CfnDeployment.DeploymentIoTJobConfigurationProperty.AbortConfig`.
	AbortConfig interface{} `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// `CfnDeployment.DeploymentIoTJobConfigurationProperty.JobExecutionsRolloutConfig`.
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// `CfnDeployment.DeploymentIoTJobConfigurationProperty.TimeoutConfig`.
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

