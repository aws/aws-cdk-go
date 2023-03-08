package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageScanningConfigurationProperty := &imageScanningConfigurationProperty{
//   	ecrConfiguration: &ecrConfigurationProperty{
//   		containerTags: []*string{
//   			jsii.String("containerTags"),
//   		},
//   		repositoryName: jsii.String("repositoryName"),
//   	},
//   	imageScanningEnabled: jsii.Boolean(false),
//   }
//
type CfnImagePipeline_ImageScanningConfigurationProperty struct {
	// `CfnImagePipeline.ImageScanningConfigurationProperty.EcrConfiguration`.
	EcrConfiguration interface{} `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// `CfnImagePipeline.ImageScanningConfigurationProperty.ImageScanningEnabled`.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

