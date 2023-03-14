package awssagemaker


// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnModelExplainabilityJobDefinition_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

