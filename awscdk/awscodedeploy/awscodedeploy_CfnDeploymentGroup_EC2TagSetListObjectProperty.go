package awscodedeploy


// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same template as EC2TagFilters.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// `EC2TagSet` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagSetListObjectProperty := &eC2TagSetListObjectProperty{
//   	ec2TagGroup: []interface{}{
//   		&eC2TagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_EC2TagSetListObjectProperty struct {
	// A list that contains other lists of Amazon EC2 instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	Ec2TagGroup interface{} `field:"optional" json:"ec2TagGroup" yaml:"ec2TagGroup"`
}

