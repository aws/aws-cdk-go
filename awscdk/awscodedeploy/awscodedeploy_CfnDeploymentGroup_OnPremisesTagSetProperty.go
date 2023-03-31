package awscodedeploy


// The `OnPremisesTagSet` property type specifies a list containing other lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `OnPremisesTagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremisesTagSetProperty := &onPremisesTagSetProperty{
//   	onPremisesTagSetList: []interface{}{
//   		&onPremisesTagSetListObjectProperty{
//   			onPremisesTagGroup: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_OnPremisesTagSetProperty struct {
	// A list that contains other lists of on-premises instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	//
	// Duplicates are not allowed.
	OnPremisesTagSetList interface{} `field:"optional" json:"onPremisesTagSetList" yaml:"onPremisesTagSetList"`
}

