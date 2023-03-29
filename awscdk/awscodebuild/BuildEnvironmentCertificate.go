package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Location of a PEM certificate on S3.
//
// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.WindowsBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.WindowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
//   			ObjectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
type BuildEnvironmentCertificate struct {
	// The bucket where the certificate is.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The full path and name of the key file.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
}

