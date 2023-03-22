package awscodedeploy


// Information about an Amazon EC2 tag filter.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2TagFilterProperty := &eC2TagFilterProperty{
//   	key: jsii.String("key"),
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnDeploymentGroup_EC2TagFilterProperty struct {
	// The tag filter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag filter type:.
	//
	// - `KEY_ONLY` : Key only.
	// - `VALUE_ONLY` : Value only.
	// - `KEY_AND_VALUE` : Key and value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The tag filter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

