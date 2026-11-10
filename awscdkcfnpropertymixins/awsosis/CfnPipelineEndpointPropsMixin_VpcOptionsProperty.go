package awsosis


// Configuration settings for the VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcOptionsProperty := &VpcOptionsProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipelineendpoint-vpcoptions.html
//
type CfnPipelineEndpointPropsMixin_VpcOptionsProperty struct {
	// A list of security group IDs that control network access to the pipeline endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipelineendpoint-vpcoptions.html#cfn-osis-pipelineendpoint-vpcoptions-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs where the pipeline endpoint network interfaces are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipelineendpoint-vpcoptions.html#cfn-osis-pipelineendpoint-vpcoptions-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

