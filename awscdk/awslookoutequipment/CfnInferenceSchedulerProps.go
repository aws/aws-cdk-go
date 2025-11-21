package awslookoutequipment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInferenceScheduler`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataInputConfiguration interface{}
//   var dataOutputConfiguration interface{}
//
//   cfnInferenceSchedulerProps := &CfnInferenceSchedulerProps{
//   	DataInputConfiguration: dataInputConfiguration,
//   	DataOutputConfiguration: dataOutputConfiguration,
//   	DataUploadFrequency: jsii.String("dataUploadFrequency"),
//   	ModelName: jsii.String("modelName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	DataDelayOffsetInMinutes: jsii.Number(123),
//   	InferenceSchedulerName: jsii.String("inferenceSchedulerName"),
//   	ServerSideKmsKeyId: jsii.String("serverSideKmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html
//
type CfnInferenceSchedulerProps struct {
	// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-datainputconfiguration
	//
	DataInputConfiguration interface{} `field:"required" json:"dataInputConfiguration" yaml:"dataInputConfiguration"`
	// Specifies configuration information for the output results for the inference scheduler, including the Amazon S3 location for the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-dataoutputconfiguration
	//
	DataOutputConfiguration interface{} `field:"required" json:"dataOutputConfiguration" yaml:"dataOutputConfiguration"`
	// How often data is uploaded to the source S3 bucket for the input data.
	//
	// This value is the length of time between data uploads. For instance, if you select 5 minutes, Amazon Lookout for Equipment will upload the real-time data to the source bucket once every 5 minutes. This frequency also determines how often Amazon Lookout for Equipment starts a scheduled inference on your data. In this example, it starts once every 5 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-datauploadfrequency
	//
	DataUploadFrequency *string `field:"required" json:"dataUploadFrequency" yaml:"dataUploadFrequency"`
	// The name of the machine learning model used for the inference scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-modelname
	//
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A period of time (in minutes) by which inference on the data is delayed after the data starts.
	//
	// For instance, if an offset delay time of five minutes was selected, inference will not begin on the data until the first data measurement after the five minute mark. For example, if five minutes is selected, the inference scheduler will wake up at the configured frequency with the additional five minute delay time to check the customer S3 bucket. The customer can upload data at the same frequency and they don't need to stop and restart the scheduler when uploading new data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-datadelayoffsetinminutes
	//
	DataDelayOffsetInMinutes *float64 `field:"optional" json:"dataDelayOffsetInMinutes" yaml:"dataDelayOffsetInMinutes"`
	// The name of the inference scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-inferenceschedulername
	//
	InferenceSchedulerName *string `field:"optional" json:"inferenceSchedulerName" yaml:"inferenceSchedulerName"`
	// Provides the identifier of the AWS KMS key used to encrypt inference scheduler data by  .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-serversidekmskeyid
	//
	ServerSideKmsKeyId *string `field:"optional" json:"serverSideKmsKeyId" yaml:"serverSideKmsKeyId"`
	// Any tags associated with the inference scheduler.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html#cfn-lookoutequipment-inferencescheduler-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

