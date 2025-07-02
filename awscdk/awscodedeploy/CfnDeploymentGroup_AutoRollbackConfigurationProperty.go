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
//   autoRollbackConfigurationProperty := &AutoRollbackConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	Events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-autorollbackconfiguration.html
//
type CfnDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-autorollbackconfiguration.html#cfn-codedeploy-deploymentgroup-autorollbackconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The event type or types that trigger a rollback.
	//
	// Valid values are `DEPLOYMENT_FAILURE` , `DEPLOYMENT_STOP_ON_ALARM` , or `DEPLOYMENT_STOP_ON_REQUEST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-autorollbackconfiguration.html#cfn-codedeploy-deploymentgroup-autorollbackconfiguration-events
	//
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

