package previewawsimagebuildermixins


// The latest version references of the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   latestVersionProperty := &LatestVersionProperty{
//   	Arn: jsii.String("arn"),
//   	Major: jsii.String("major"),
//   	Minor: jsii.String("minor"),
//   	Patch: jsii.String("patch"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-latestversion.html
//
type CfnImagePropsMixin_LatestVersionProperty struct {
	// The Amazon Resource Name (ARN) of the image.
	//
	// > Semantic versioning is included in each object's Amazon Resource Name (ARN), at the level that applies to that object as follows:
	// >
	// > - Versionless ARNs and Name ARNs do not include specific values in any of the nodes. The nodes are either left off entirely, or they are specified as wildcards, for example: x.x.x.
	// > - Version ARNs have only the first three nodes: <major>.<minor>.<patch>
	// > - Build version ARNs have all four nodes, and point to a specific build for a specific version of an object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-latestversion.html#cfn-imagebuilder-image-latestversion-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The latest version ARN of the created image, with the same major version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-latestversion.html#cfn-imagebuilder-image-latestversion-major
	//
	Major *string `field:"optional" json:"major" yaml:"major"`
	// The latest version ARN of the created image, with the same minor version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-latestversion.html#cfn-imagebuilder-image-latestversion-minor
	//
	Minor *string `field:"optional" json:"minor" yaml:"minor"`
	// The latest version ARN of the created image, with the same patch version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-latestversion.html#cfn-imagebuilder-image-latestversion-patch
	//
	Patch *string `field:"optional" json:"patch" yaml:"patch"`
}

