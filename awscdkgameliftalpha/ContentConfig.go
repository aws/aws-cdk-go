package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Result of binding `Content` into a `Build`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   contentConfig := &ContentConfig{
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// Experimental.
type ContentConfig struct {
	// The location of the content in S3.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
}

