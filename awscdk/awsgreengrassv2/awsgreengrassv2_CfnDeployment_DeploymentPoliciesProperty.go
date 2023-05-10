package awsgreengrassv2


// Contains information about policies that define how a deployment updates components and handles failure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentPoliciesProperty := &DeploymentPoliciesProperty{
//   	ComponentUpdatePolicy: &DeploymentComponentUpdatePolicyProperty{
//   		Action: jsii.String("action"),
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	ConfigurationValidationPolicy: &DeploymentConfigurationValidationPolicyProperty{
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	FailureHandlingPolicy: jsii.String("failureHandlingPolicy"),
//   }
//
type CfnDeployment_DeploymentPoliciesProperty struct {
	// The component update policy for the configuration deployment.
	//
	// This policy defines when it's safe to deploy the configuration to devices.
	ComponentUpdatePolicy interface{} `field:"optional" json:"componentUpdatePolicy" yaml:"componentUpdatePolicy"`
	// The configuration validation policy for the configuration deployment.
	//
	// This policy defines how long each component has to validate its configure updates.
	ConfigurationValidationPolicy interface{} `field:"optional" json:"configurationValidationPolicy" yaml:"configurationValidationPolicy"`
	// The failure handling policy for the configuration deployment. This policy defines what to do if the deployment fails.
	//
	// Default: `ROLLBACK`.
	FailureHandlingPolicy *string `field:"optional" json:"failureHandlingPolicy" yaml:"failureHandlingPolicy"`
}

