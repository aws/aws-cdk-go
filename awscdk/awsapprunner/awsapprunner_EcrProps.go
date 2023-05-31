package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
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
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image tag.
	// Deprecated: use `tagOrDigest`.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Image tag or digest (digests must start with `sha256:`).
	// Experimental.
	TagOrDigest *string `field:"optional" json:"tagOrDigest" yaml:"tagOrDigest"`
}

