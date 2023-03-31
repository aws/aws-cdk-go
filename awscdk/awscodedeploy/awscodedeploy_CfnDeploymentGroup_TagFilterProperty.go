package awscodedeploy


// `TagFilter` is a property type of the [AWS::CodeDeploy::DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies which on-premises instances to associate with the deployment group. To register on-premise instances with AWS CodeDeploy , see [Configure Existing On-Premises Instances by Using AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* .
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnDeploymentGroup_TagFilterProperty struct {
	// The on-premises instance tag filter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The on-premises instance tag filter type:.
	//
	// - KEY_ONLY: Key only.
	// - VALUE_ONLY: Value only.
	// - KEY_AND_VALUE: Key and value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The on-premises instance tag filter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

