package awsiot


// The properties of a virtual private cloud (VPC) destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcDestinationPropertiesProperty := &vpcDestinationPropertiesProperty{
//   	roleArn: jsii.String("roleArn"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnTopicRuleDestination_VpcDestinationPropertiesProperty struct {
	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The security groups of the VPC destination.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnet IDs of the VPC destination.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

