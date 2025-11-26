package previewawsiotmixins


// The properties of a virtual private cloud (VPC) destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcDestinationPropertiesProperty := &VpcDestinationPropertiesProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-vpcdestinationproperties.html
//
type CfnTopicRuleDestinationPropsMixin_VpcDestinationPropertiesProperty struct {
	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-vpcdestinationproperties.html#cfn-iot-topicruledestination-vpcdestinationproperties-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The security groups of the VPC destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-vpcdestinationproperties.html#cfn-iot-topicruledestination-vpcdestinationproperties-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnet IDs of the VPC destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-vpcdestinationproperties.html#cfn-iot-topicruledestination-vpcdestinationproperties-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-vpcdestinationproperties.html#cfn-iot-topicruledestination-vpcdestinationproperties-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

