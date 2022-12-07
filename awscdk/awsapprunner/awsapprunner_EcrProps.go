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
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcr(&ecrProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(80),
//   		},
//   		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		tagOrDigest: jsii.String("latest"),
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

