package awslookoutequipment


// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html
//
type CfnInferenceScheduler_DataInputConfigurationProperty struct {
	// Specifies configuration information for the input data for the inference, including input data S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-s3inputconfiguration
	//
	S3InputConfiguration interface{} `field:"required" json:"s3InputConfiguration" yaml:"s3InputConfiguration"`
	// Specifies configuration information for the input data for the inference, including timestamp format and delimiter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-inferenceinputnameconfiguration
	//
	InferenceInputNameConfiguration interface{} `field:"optional" json:"inferenceInputNameConfiguration" yaml:"inferenceInputNameConfiguration"`
	// Indicates the difference between your time zone and Greenwich Mean Time (GMT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration-inputtimezoneoffset
	//
	InputTimeZoneOffset *string `field:"optional" json:"inputTimeZoneOffset" yaml:"inputTimeZoneOffset"`
}

