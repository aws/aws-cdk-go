package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Configuration for the environment file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileConfig := &EnvironmentFileConfig{
//   	FileType: awscdk.Aws_ecs.EnvironmentFileType_S3,
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type EnvironmentFileConfig struct {
	// The type of environment file.
	FileType EnvironmentFileType `field:"required" json:"fileType" yaml:"fileType"`
	// The location of the environment file in S3.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

