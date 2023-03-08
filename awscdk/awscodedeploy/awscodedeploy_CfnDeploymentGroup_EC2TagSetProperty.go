package awscodedeploy


// The `EC2TagSet` property type specifies information about groups of tags applied to Amazon EC2 instances.
//
// The deployment group includes only Amazon EC2 instances identified by all the tag groups. `EC2TagSet` cannot be used in the same template as `EC2TagFilter` .
//
// For information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagSetProperty := &eC2TagSetProperty{
//   	ec2TagSetList: []interface{}{
//   		&eC2TagSetListObjectProperty{
//   			ec2TagGroup: []interface{}{
//   				&eC2TagFilterProperty{
//   					key: jsii.String("key"),
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_EC2TagSetProperty struct {
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group.
	//
	// Duplicates are not allowed.
	Ec2TagSetList interface{} `field:"optional" json:"ec2TagSetList" yaml:"ec2TagSetList"`
}

