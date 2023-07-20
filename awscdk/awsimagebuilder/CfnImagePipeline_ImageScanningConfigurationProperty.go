package awsimagebuilder


// Determines if tests should run after building the image.
//
// Image Builder defaults to enable tests to run following the image build, before image distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageScanningConfigurationProperty := &ImageScanningConfigurationProperty{
//   	EcrConfiguration: &EcrConfigurationProperty{
//   		ContainerTags: []*string{
//   			jsii.String("containerTags"),
//   		},
//   		RepositoryName: jsii.String("repositoryName"),
//   	},
//   	ImageScanningEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html
//
type CfnImagePipeline_ImageScanningConfigurationProperty struct {
	// Settings for Image Builder to configure the ECR repository and output container images that are scanned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-ecrconfiguration
	//
	EcrConfiguration interface{} `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-imagescanningenabled
	//
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

