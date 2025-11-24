package mixinsawsrtbfabric


// Describes the configuration of an auto scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoScalingGroupsConfigurationProperty := &AutoScalingGroupsConfigurationProperty{
//   	AutoScalingGroupNameList: []*string{
//   		jsii.String("autoScalingGroupNameList"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html
//
type CfnResponderGatewayPropsMixin_AutoScalingGroupsConfigurationProperty struct {
	// The names of the auto scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html#cfn-rtbfabric-respondergateway-autoscalinggroupsconfiguration-autoscalinggroupnamelist
	//
	AutoScalingGroupNameList *[]*string `field:"optional" json:"autoScalingGroupNameList" yaml:"autoScalingGroupNameList"`
	// The role ARN of the auto scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html#cfn-rtbfabric-respondergateway-autoscalinggroupsconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

