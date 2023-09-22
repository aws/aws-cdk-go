package awssagemaker


// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &MonitoringOutputConfigProperty{
//   	MonitoringOutputs: []interface{}{
//   		&MonitoringOutputProperty{
//   			S3Output: &S3OutputProperty{
//   				LocalPath: jsii.String("localPath"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				S3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html
//
type CfnMonitoringSchedule_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html#cfn-sagemaker-monitoringschedule-monitoringoutputconfig-monitoringoutputs
	//
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS ) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html#cfn-sagemaker-monitoringschedule-monitoringoutputconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

