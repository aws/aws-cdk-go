package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamProps := &CfnStreamProps{
//   	Files: []interface{}{
//   		&StreamFileProperty{
//   			FileId: jsii.Number(123),
//   			S3Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StreamId: jsii.String("streamId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html
//
type CfnStreamProps struct {
	// The files to stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-files
	//
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// An IAM role that allows the IoT service principal to access your S3 files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The stream ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-streamid
	//
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
	// The description of the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata which can be used to manage streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

