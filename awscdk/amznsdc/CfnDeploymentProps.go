package amznsdc


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	ConfigName: jsii.String("configName"),
//   	Dimension: jsii.String("dimension"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   	Stage: jsii.String("stage"),
//
//   	// the properties below are optional
//   	PipelineId: jsii.String("pipelineId"),
//   	TargetRegionOverride: jsii.String("targetRegionOverride"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html
//
type CfnDeploymentProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-configname
	//
	ConfigName *string `field:"required" json:"configName" yaml:"configName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-dimension
	//
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-s3key
	//
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-stage
	//
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-pipelineid
	//
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sdc-deployment.html#cfn-sdc-deployment-targetregionoverride
	//
	TargetRegionOverride *string `field:"optional" json:"targetRegionOverride" yaml:"targetRegionOverride"`
}

