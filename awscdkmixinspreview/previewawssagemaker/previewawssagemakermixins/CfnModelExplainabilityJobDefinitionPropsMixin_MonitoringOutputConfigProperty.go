package previewawssagemakermixins


// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringOutputConfigProperty := &MonitoringOutputConfigProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MonitoringOutputs: []interface{}{
//   		&MonitoringOutputProperty{
//   			S3Output: &S3OutputProperty{
//   				LocalPath: jsii.String("localPath"),
//   				S3UploadMode: jsii.String("s3UploadMode"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-monitoringoutputconfig.html
//
type CfnModelExplainabilityJobDefinitionPropsMixin_MonitoringOutputConfigProperty struct {
	// The AWS Key Management Service ( AWS  ) key that Amazon SageMaker AI uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-monitoringoutputconfig.html#cfn-sagemaker-modelexplainabilityjobdefinition-monitoringoutputconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-monitoringoutputconfig.html#cfn-sagemaker-modelexplainabilityjobdefinition-monitoringoutputconfig-monitoringoutputs
	//
	MonitoringOutputs interface{} `field:"optional" json:"monitoringOutputs" yaml:"monitoringOutputs"`
}

