package previewawsimagebuildermixins


// The latest version references of the component.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-component-latestversion.html
//
type CfnComponentPropsMixin_LatestVersionProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-component-latestversion.html#cfn-imagebuilder-component-latestversion-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The latest version ARN of the created component, with the same major version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-component-latestversion.html#cfn-imagebuilder-component-latestversion-major
	//
	Major *string `field:"optional" json:"major" yaml:"major"`
	// The latest version ARN of the created component, with the same minor version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-component-latestversion.html#cfn-imagebuilder-component-latestversion-minor
	//
	Minor *string `field:"optional" json:"minor" yaml:"minor"`
	// The latest version ARN of the created component, with the same patch version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-component-latestversion.html#cfn-imagebuilder-component-latestversion-patch
	//
	Patch *string `field:"optional" json:"patch" yaml:"patch"`
}

