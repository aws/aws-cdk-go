package awspipes


// This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaAccessConfigurationVpcProperty := &SelfManagedKafkaAccessConfigurationVpcProperty{
//   	SecurityGroup: []*string{
//   		jsii.String("securityGroup"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc.html
//
type CfnPipe_SelfManagedKafkaAccessConfigurationVpcProperty struct {
	// Specifies the security groups associated with the stream.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc.html#cfn-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc-securitygroup
	//
	SecurityGroup *[]*string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Specifies the subnets associated with the stream.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc.html#cfn-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

