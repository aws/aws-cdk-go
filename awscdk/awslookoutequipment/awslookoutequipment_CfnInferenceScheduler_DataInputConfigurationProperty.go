package awslookoutequipment


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataInputConfigurationProperty := &DataInputConfigurationProperty{
//   	S3InputConfiguration: &S3InputConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	InferenceInputNameConfiguration: &InputNameConfigurationProperty{
//   		ComponentTimestampDelimiter: jsii.String("componentTimestampDelimiter"),
//   		TimestampFormat: jsii.String("timestampFormat"),
//   	},
//   	InputTimeZoneOffset: jsii.String("inputTimeZoneOffset"),
//   }
//
type CfnInferenceScheduler_DataInputConfigurationProperty struct {
	// `CfnInferenceScheduler.DataInputConfigurationProperty.S3InputConfiguration`.
	S3InputConfiguration interface{} `field:"required" json:"s3InputConfiguration" yaml:"s3InputConfiguration"`
	// `CfnInferenceScheduler.DataInputConfigurationProperty.InferenceInputNameConfiguration`.
	InferenceInputNameConfiguration interface{} `field:"optional" json:"inferenceInputNameConfiguration" yaml:"inferenceInputNameConfiguration"`
	// `CfnInferenceScheduler.DataInputConfigurationProperty.InputTimeZoneOffset`.
	InputTimeZoneOffset *string `field:"optional" json:"inputTimeZoneOffset" yaml:"inputTimeZoneOffset"`
}

