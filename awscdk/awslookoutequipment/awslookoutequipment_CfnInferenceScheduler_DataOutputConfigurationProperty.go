package awslookoutequipment


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataOutputConfigurationProperty := &dataOutputConfigurationProperty{
//   	s3OutputConfiguration: &s3OutputConfigurationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnInferenceScheduler_DataOutputConfigurationProperty struct {
	// `CfnInferenceScheduler.DataOutputConfigurationProperty.S3OutputConfiguration`.
	S3OutputConfiguration interface{} `field:"required" json:"s3OutputConfiguration" yaml:"s3OutputConfiguration"`
	// `CfnInferenceScheduler.DataOutputConfigurationProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

