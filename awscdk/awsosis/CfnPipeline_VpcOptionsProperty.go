package awsosis


// Options that specify the subnets and security groups for an OpenSearch Ingestion VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOptionsProperty := &VpcOptionsProperty{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	VpcAttachmentOptions: &VpcAttachmentOptionsProperty{
//   		AttachToVpc: jsii.Boolean(false),
//   		CidrBlock: jsii.String("cidrBlock"),
//   	},
//   	VpcEndpointManagement: jsii.String("vpcEndpointManagement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcoptions.html
//
type CfnPipeline_VpcOptionsProperty struct {
	// A list of subnet IDs associated with the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcoptions.html#cfn-osis-pipeline-vpcoptions-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// A list of security groups associated with the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcoptions.html#cfn-osis-pipeline-vpcoptions-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Options for attaching a VPC to a pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcoptions.html#cfn-osis-pipeline-vpcoptions-vpcattachmentoptions
	//
	VpcAttachmentOptions interface{} `field:"optional" json:"vpcAttachmentOptions" yaml:"vpcAttachmentOptions"`
	// Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcoptions.html#cfn-osis-pipeline-vpcoptions-vpcendpointmanagement
	//
	VpcEndpointManagement *string `field:"optional" json:"vpcEndpointManagement" yaml:"vpcEndpointManagement"`
}

