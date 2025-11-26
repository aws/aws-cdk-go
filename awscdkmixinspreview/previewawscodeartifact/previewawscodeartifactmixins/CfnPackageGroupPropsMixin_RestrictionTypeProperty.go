package previewawscodeartifactmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   restrictionTypeProperty := &RestrictionTypeProperty{
//   	Repositories: []*string{
//   		jsii.String("repositories"),
//   	},
//   	RestrictionMode: jsii.String("restrictionMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html
//
type CfnPackageGroupPropsMixin_RestrictionTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-repositories
	//
	Repositories *[]*string `field:"optional" json:"repositories" yaml:"repositories"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-restrictionmode
	//
	RestrictionMode *string `field:"optional" json:"restrictionMode" yaml:"restrictionMode"`
}

