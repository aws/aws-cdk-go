package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// Properties of the image repository for `Source.fromEcr()`.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcr(&EcrProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(80),
//   		},
//   		Repository: ecr.Repository_FromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		TagOrDigest: jsii.String("latest"),
//   	}),
//   })
//
// Experimental.
type EcrProps struct {
	// Represents the ECR repository.
	// Experimental.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The image configuration for the image from ECR.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Default: - no image configuration will be passed. The default `port` will be 8080.
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image tag.
	// Default: - 'latest'.
	//
	// Deprecated: use `tagOrDigest`.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Image tag or digest (digests must start with `sha256:`).
	// Default: - 'latest'.
	//
	// Experimental.
	TagOrDigest *string `field:"optional" json:"tagOrDigest" yaml:"tagOrDigest"`
}

