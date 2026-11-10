package awsimagebuilder


// Contains settings for Image Builder image resource and container image scans.
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
//   	ImageScanningTypes: []*string{
//   		jsii.String("imageScanningTypes"),
//   	},
//   	ScanResultCondition: &ScanResultConditionProperty{
//   		MaximumSeverityLabel: jsii.String("maximumSeverityLabel"),
//
//   		// the properties below are optional
//   		FixAvailable: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html
//
type CfnImagePipeline_ImageScanningConfigurationProperty struct {
	// Contains Amazon ECR settings for vulnerability scans.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-ecrconfiguration
	//
	EcrConfiguration interface{} `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// A setting that indicates whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-imagescanningenabled
	//
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
	// The type of vulnerability scan that Image Builder runs for AMI output when ImageScanningEnabled is true.
	//
	// Specify a single value. If you omit ImageScanningTypes, Image Builder uses TEST_INSTANCE_CVE. With TEST_INSTANCE_CVE, Amazon Inspector scans the test instance during image tests, and you must activate Amazon Inspector for Amazon EC2 scanning in your account. With ON_DEMAND_AMI_CVE, Amazon Inspector starts an on-demand scan of the output AMI during image tests. You must build from an image recipe, but you do not need to activate Amazon Inspector in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-imagescanningtypes
	//
	ImageScanningTypes *[]*string `field:"optional" json:"imageScanningTypes" yaml:"imageScanningTypes"`
	// The condition that the output AMI's scan results must meet for the image to pass.
	//
	// When you set it, the test phase waits for the scan to finish, and Image Builder fails the image if a counted finding exceeds MaximumSeverityLabel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html#cfn-imagebuilder-imagepipeline-imagescanningconfiguration-scanresultcondition
	//
	ScanResultCondition interface{} `field:"optional" json:"scanResultCondition" yaml:"scanResultCondition"`
}

