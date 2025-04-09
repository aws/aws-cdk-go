package awslambda


// Properties for defining a `CfnLayerVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionProps := &CfnLayerVersionProps{
//   	Content: &ContentProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//
//   		// the properties below are optional
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//
//   	// the properties below are optional
//   	CompatibleArchitectures: []*string{
//   		jsii.String("compatibleArchitectures"),
//   	},
//   	CompatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	Description: jsii.String("description"),
//   	LayerName: jsii.String("layerName"),
//   	LicenseInfo: jsii.String("licenseInfo"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html
//
type CfnLayerVersionProps struct {
	// The function layer archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// A list of compatible [instruction set architectures](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatiblearchitectures
	//
	CompatibleArchitectures *[]*string `field:"optional" json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// A list of compatible [function runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) . Used for filtering with [ListLayers](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayers.html) and [ListLayerVersions](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayerVersions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-compatibleruntimes
	//
	CompatibleRuntimes *[]*string `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// The description of the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or Amazon Resource Name (ARN) of the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-layername
	//
	LayerName *string `field:"optional" json:"layerName" yaml:"layerName"`
	// The layer's software license. It can be any of the following:.
	//
	// - An [SPDX license identifier](https://docs.aws.amazon.com/https://spdx.org/licenses/) . For example, `MIT` .
	// - The URL of a license hosted on the internet. For example, `https://opensource.org/licenses/MIT` .
	// - The full text of the license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversion.html#cfn-lambda-layerversion-licenseinfo
	//
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
}

