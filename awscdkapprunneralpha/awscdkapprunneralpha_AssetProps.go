// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
)

// Properties of the image repository for `Source.fromAsset()`.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromAsset(&assetProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		asset: imageAsset,
//   	}),
//   })
//
// Experimental.
type AssetProps struct {
	// Represents the docker image asset.
	// Experimental.
	Asset awsecrassets.DockerImageAsset `field:"required" json:"asset" yaml:"asset"`
	// The image configuration for the image built from the asset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

