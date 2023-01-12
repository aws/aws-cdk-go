package awsgreengrassv2


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rateIncreaseCriteria interface{}
//
//   cfnDeploymentProps := &cfnDeploymentProps{
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	components: map[string]interface{}{
//   		"componentsKey": &ComponentDeploymentSpecificationProperty{
//   			"componentVersion": jsii.String("componentVersion"),
//   			"configurationUpdate": &ComponentConfigurationUpdateProperty{
//   				"merge": jsii.String("merge"),
//   				"reset": []*string{
//   					jsii.String("reset"),
//   				},
//   			},
//   			"runWith": &ComponentRunWithProperty{
//   				"posixUser": jsii.String("posixUser"),
//   				"systemResourceLimits": &SystemResourceLimitsProperty{
//   					"cpus": jsii.Number(123),
//   					"memory": jsii.Number(123),
//   				},
//   				"windowsUser": jsii.String("windowsUser"),
//   			},
//   		},
//   	},
//   	deploymentName: jsii.String("deploymentName"),
//   	deploymentPolicies: &deploymentPoliciesProperty{
//   		componentUpdatePolicy: &deploymentComponentUpdatePolicyProperty{
//   			action: jsii.String("action"),
//   			timeoutInSeconds: jsii.Number(123),
//   		},
//   		configurationValidationPolicy: &deploymentConfigurationValidationPolicyProperty{
//   			timeoutInSeconds: jsii.Number(123),
//   		},
//   		failureHandlingPolicy: jsii.String("failureHandlingPolicy"),
//   	},
//   	iotJobConfiguration: &deploymentIoTJobConfigurationProperty{
//   		abortConfig: &ioTJobAbortConfigProperty{
//   			criteriaList: []interface{}{
//   				&ioTJobAbortCriteriaProperty{
//   					action: jsii.String("action"),
//   					failureType: jsii.String("failureType"),
//   					minNumberOfExecutedThings: jsii.Number(123),
//   					thresholdPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   		jobExecutionsRolloutConfig: &ioTJobExecutionsRolloutConfigProperty{
//   			exponentialRate: &ioTJobExponentialRolloutRateProperty{
//   				baseRatePerMinute: jsii.Number(123),
//   				incrementFactor: jsii.Number(123),
//   				rateIncreaseCriteria: rateIncreaseCriteria,
//   			},
//   			maximumPerMinute: jsii.Number(123),
//   		},
//   		timeoutConfig: &ioTJobTimeoutConfigProperty{
//   			inProgressTimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnDeploymentProps struct {
	// `AWS::GreengrassV2::Deployment.TargetArn`.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// `AWS::GreengrassV2::Deployment.Components`.
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// `AWS::GreengrassV2::Deployment.DeploymentName`.
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// `AWS::GreengrassV2::Deployment.DeploymentPolicies`.
	DeploymentPolicies interface{} `field:"optional" json:"deploymentPolicies" yaml:"deploymentPolicies"`
	// `AWS::GreengrassV2::Deployment.IotJobConfiguration`.
	IotJobConfiguration interface{} `field:"optional" json:"iotJobConfiguration" yaml:"iotJobConfiguration"`
	// `AWS::GreengrassV2::Deployment.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

