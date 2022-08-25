package awscodedeploy


// The `AutoRollbackConfiguration` property type configures automatic rollback for an AWS CodeDeploy deployment group when a deployment is not completed successfully.
//
// For more information, see [Automatic Rollbacks](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployments-rollback-and-redeploy.html#deployments-rollback-and-redeploy-automatic-rollbacks) in the *AWS CodeDeploy User Guide* .
//
// `AutoRollbackConfiguration` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRollbackConfigurationProperty := &autoRollbackConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
type CfnDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The event type or types that trigger a rollback.
	//
	// Valid values are `DEPLOYMENT_FAILURE` , `DEPLOYMENT_STOP_ON_ALARM` , or `DEPLOYMENT_STOP_ON_REQUEST` .
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

