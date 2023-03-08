package awsgroundstation


// Information about IAM roles, subnets, and security groups needed for this DataflowEndpointGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityDetailsProperty := &securityDetailsProperty{
//   	roleArn: jsii.String("roleArn"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDataflowEndpointGroup_SecurityDetailsProperty struct {
	// The ARN of a role which Ground Station has permission to assume, such as `arn:aws:iam::1234567890:role/DataDeliveryServiceRole` .
	//
	// Ground Station will assume this role and create an ENI in your VPC on the specified subnet upon creation of a dataflow endpoint group. This ENI is used as the ingress/egress point for data streamed during a satellite contact.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The security group Ids of the security role, such as `sg-1234567890abcdef0` .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet Ids of the security details, such as `subnet-12345678` .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

