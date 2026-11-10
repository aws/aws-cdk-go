package awsosis


// Properties for CfnPipelineEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPipelineEndpointMixinProps := &CfnPipelineEndpointMixinProps{
//   	PipelineArn: jsii.String("pipelineArn"),
//   	VpcOptions: &VpcOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineendpoint.html
//
type CfnPipelineEndpointMixinProps struct {
	// The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineendpoint.html#cfn-osis-pipelineendpoint-pipelinearn
	//
	PipelineArn *string `field:"optional" json:"pipelineArn" yaml:"pipelineArn"`
	// Configuration settings for the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineendpoint.html#cfn-osis-pipelineendpoint-vpcoptions
	//
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

