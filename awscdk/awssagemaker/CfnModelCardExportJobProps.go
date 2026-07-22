package awssagemaker


// Properties for defining a `CfnModelCardExportJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelCardExportJobProps := &CfnModelCardExportJobProps{
//   	ModelCardExportJobName: jsii.String("modelCardExportJobName"),
//   	ModelCardName: jsii.String("modelCardName"),
//   	OutputConfig: &ModelCardExportOutputConfigProperty{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//   	},
//
//   	// the properties below are optional
//   	ModelCardVersion: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html
//
type CfnModelCardExportJobProps struct {
	// The name of the model card export job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardexportjobname
	//
	ModelCardExportJobName *string `field:"required" json:"modelCardExportJobName" yaml:"modelCardExportJobName"`
	// The name or Amazon Resource Name (ARN) of the model card to export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardname
	//
	ModelCardName *string `field:"required" json:"modelCardName" yaml:"modelCardName"`
	// Configure the export output details for an Amazon SageMaker Model Card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-outputconfig
	//
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// The version of the model card to export.
	//
	// If a version is not provided, then the latest version of the model card is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardversion
	//
	ModelCardVersion *float64 `field:"optional" json:"modelCardVersion" yaml:"modelCardVersion"`
}

