package awssagemaker


// Configure the export output details for an Amazon SageMaker Model Card.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelCardExportOutputConfigProperty := &ModelCardExportOutputConfigProperty{
//   	S3OutputPath: jsii.String("s3OutputPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcardexportjob-modelcardexportoutputconfig.html
//
type CfnModelCardExportJob_ModelCardExportOutputConfigProperty struct {
	// The Amazon S3 output path to export your model card PDF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcardexportjob-modelcardexportoutputconfig.html#cfn-sagemaker-modelcardexportjob-modelcardexportoutputconfig-s3outputpath
	//
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
}

