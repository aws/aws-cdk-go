package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &VpcConfigProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-vpcconfig.html
//
type CfnEndpointConfig_VpcConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-vpcconfig.html#cfn-sagemaker-endpointconfig-vpcconfig-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-vpcconfig.html#cfn-sagemaker-endpointconfig-vpcconfig-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

