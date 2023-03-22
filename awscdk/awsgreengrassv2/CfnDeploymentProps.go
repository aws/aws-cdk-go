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
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Components: map[string]interface{}{
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
//   	DeploymentName: jsii.String("deploymentName"),
//   	DeploymentPolicies: &DeploymentPoliciesProperty{
//   		ComponentUpdatePolicy: &DeploymentComponentUpdatePolicyProperty{
//   			Action: jsii.String("action"),
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		ConfigurationValidationPolicy: &DeploymentConfigurationValidationPolicyProperty{
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		FailureHandlingPolicy: jsii.String("failureHandlingPolicy"),
//   	},
//   	IotJobConfiguration: &DeploymentIoTJobConfigurationProperty{
//   		AbortConfig: &IoTJobAbortConfigProperty{
//   			CriteriaList: []interface{}{
//   				&IoTJobAbortCriteriaProperty{
//   					Action: jsii.String("action"),
//   					FailureType: jsii.String("failureType"),
//   					MinNumberOfExecutedThings: jsii.Number(123),
//   					ThresholdPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   		JobExecutionsRolloutConfig: &IoTJobExecutionsRolloutConfigProperty{
//   			ExponentialRate: &IoTJobExponentialRolloutRateProperty{
//   				BaseRatePerMinute: jsii.Number(123),
//   				IncrementFactor: jsii.Number(123),
//   				RateIncreaseCriteria: rateIncreaseCriteria,
//   			},
//   			MaximumPerMinute: jsii.Number(123),
//   		},
//   		TimeoutConfig: &IoTJobTimeoutConfigProperty{
//   			InProgressTimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	ParentTargetArn: jsii.String("parentTargetArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnDeploymentProps struct {
	// The ARN of the target AWS IoT thing or thing group.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The components to deploy.
	//
	// This is a dictionary, where each key is the name of a component, and each key's value is the version and configuration to deploy for that component.
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The name of the deployment.
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// The deployment policies for the deployment.
	//
	// These policies define how the deployment updates components and handles failure.
	DeploymentPolicies interface{} `field:"optional" json:"deploymentPolicies" yaml:"deploymentPolicies"`
	// The job configuration for the deployment configuration.
	//
	// The job configuration specifies the rollout, timeout, and stop configurations for the deployment configuration.
	IotJobConfiguration interface{} `field:"optional" json:"iotJobConfiguration" yaml:"iotJobConfiguration"`
	// The parent deployment's [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) for a subdeployment.
	ParentTargetArn *string `field:"optional" json:"parentTargetArn" yaml:"parentTargetArn"`
	// Application-specific metadata to attach to the deployment.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

