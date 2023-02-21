package awsimagebuilder


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
type CfnImage_ImageScanningConfigurationProperty struct {
	// `CfnImage.ImageScanningConfigurationProperty.EcrConfiguration`.
	EcrConfiguration interface{} `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// `CfnImage.ImageScanningConfigurationProperty.ImageScanningEnabled`.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}

