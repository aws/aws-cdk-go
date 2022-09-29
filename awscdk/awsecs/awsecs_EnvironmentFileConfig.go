package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Configuration for the environment file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileConfig := &environmentFileConfig{
//   	fileType: awscdk.Aws_ecs.environmentFileType_S3,
//   	s3Location: &location{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Experimental.
type EnvironmentFileConfig struct {
	// The type of environment file.
	// Experimental.
	FileType EnvironmentFileType `field:"required" json:"fileType" yaml:"fileType"`
	// The location of the environment file in S3.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

