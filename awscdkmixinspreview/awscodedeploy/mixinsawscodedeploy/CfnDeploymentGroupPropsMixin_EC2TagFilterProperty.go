package mixinsawscodedeploy


// Information about an Amazon EC2 tag filter.
//
// For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment Groups in AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-tagging.html) in the *AWS CodeDeploy User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2TagFilterProperty := &EC2TagFilterProperty{
//   	Key: jsii.String("key"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilter.html
//
type CfnDeploymentGroupPropsMixin_EC2TagFilterProperty struct {
	// The tag filter key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilter.html#cfn-codedeploy-deploymentgroup-ec2tagfilter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag filter type:.
	//
	// - `KEY_ONLY` : Key only.
	// - `VALUE_ONLY` : Value only.
	// - `KEY_AND_VALUE` : Key and value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilter.html#cfn-codedeploy-deploymentgroup-ec2tagfilter-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The tag filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilter.html#cfn-codedeploy-deploymentgroup-ec2tagfilter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

