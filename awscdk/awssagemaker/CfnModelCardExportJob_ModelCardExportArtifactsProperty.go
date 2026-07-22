package awssagemaker


// The artifacts of the model card export job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelCardExportArtifactsProperty := &ModelCardExportArtifactsProperty{
//   	S3ExportArtifacts: jsii.String("s3ExportArtifacts"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcardexportjob-modelcardexportartifacts.html
//
type CfnModelCardExportJob_ModelCardExportArtifactsProperty struct {
	// The Amazon S3 URI of the exported model artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcardexportjob-modelcardexportartifacts.html#cfn-sagemaker-modelcardexportjob-modelcardexportartifacts-s3exportartifacts
	//
	S3ExportArtifacts *string `field:"required" json:"s3ExportArtifacts" yaml:"s3ExportArtifacts"`
}

