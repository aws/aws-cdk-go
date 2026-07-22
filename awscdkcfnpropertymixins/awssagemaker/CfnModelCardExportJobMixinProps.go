package awssagemaker


// Properties for CfnModelCardExportJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnModelCardExportJobMixinProps := &CfnModelCardExportJobMixinProps{
//   	ModelCardExportJobName: jsii.String("modelCardExportJobName"),
//   	ModelCardName: jsii.String("modelCardName"),
//   	ModelCardVersion: jsii.Number(123),
//   	OutputConfig: &ModelCardExportOutputConfigProperty{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html
//
type CfnModelCardExportJobMixinProps struct {
	// The name of the model card export job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardexportjobname
	//
	ModelCardExportJobName *string `field:"optional" json:"modelCardExportJobName" yaml:"modelCardExportJobName"`
	// The name or Amazon Resource Name (ARN) of the model card to export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardname
	//
	ModelCardName *string `field:"optional" json:"modelCardName" yaml:"modelCardName"`
	// The version of the model card to export.
	//
	// If a version is not provided, then the latest version of the model card is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-modelcardversion
	//
	ModelCardVersion *float64 `field:"optional" json:"modelCardVersion" yaml:"modelCardVersion"`
	// Configure the export output details for an Amazon SageMaker Model Card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelcardexportjob.html#cfn-sagemaker-modelcardexportjob-outputconfig
	//
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
}

