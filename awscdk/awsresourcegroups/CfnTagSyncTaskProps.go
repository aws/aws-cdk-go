package awsresourcegroups


// Properties for defining a `CfnTagSyncTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagSyncTaskProps := &CfnTagSyncTaskProps{
//   	Group: jsii.String("group"),
//   	RoleArn: jsii.String("roleArn"),
//   	TagKey: jsii.String("tagKey"),
//   	TagValue: jsii.String("tagValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html
//
type CfnTagSyncTaskProps struct {
	// The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html#cfn-resourcegroups-tagsynctask-group
	//
	Group *string `field:"required" json:"group" yaml:"group"`
	// The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html#cfn-resourcegroups-tagsynctask-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html#cfn-resourcegroups-tagsynctask-tagkey
	//
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// The tag value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-tagsynctask.html#cfn-resourcegroups-tagsynctask-tagvalue
	//
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
}

