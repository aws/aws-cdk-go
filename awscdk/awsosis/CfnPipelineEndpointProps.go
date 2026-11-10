package awsosis


// Properties for defining a `CfnPipelineEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineEndpointProps := &CfnPipelineEndpointProps{
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
type CfnPipelineEndpointProps struct {
	// The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineendpoint.html#cfn-osis-pipelineendpoint-pipelinearn
	//
	PipelineArn *string `field:"required" json:"pipelineArn" yaml:"pipelineArn"`
	// Configuration settings for the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineendpoint.html#cfn-osis-pipelineendpoint-vpcoptions
	//
	VpcOptions interface{} `field:"required" json:"vpcOptions" yaml:"vpcOptions"`
}

