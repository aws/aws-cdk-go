package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStreamPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnStreamMixinProps := &CfnStreamMixinProps{
//   	Description: jsii.String("description"),
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
type CfnStreamMixinProps struct {
	// The description of the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The files to stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-files
	//
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// An IAM role that allows the IoT service principal to access your S3 files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The stream ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// Metadata which can be used to manage streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-stream.html#cfn-iot-stream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

