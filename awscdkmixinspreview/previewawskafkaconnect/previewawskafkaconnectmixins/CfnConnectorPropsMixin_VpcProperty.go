package previewawskafkaconnectmixins


// Information about the VPC in which the connector resides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcProperty := &VpcProperty{
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-vpc.html
//
type CfnConnectorPropsMixin_VpcProperty struct {
	// The security group IDs for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-vpc.html#cfn-kafkaconnect-connector-vpc-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-vpc.html#cfn-kafkaconnect-connector-vpc-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

