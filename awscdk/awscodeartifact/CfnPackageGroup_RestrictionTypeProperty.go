package awscodeartifact


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictionTypeProperty := &RestrictionTypeProperty{
//   	RestrictionMode: jsii.String("restrictionMode"),
//
//   	// the properties below are optional
//   	Repositories: []*string{
//   		jsii.String("repositories"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html
//
type CfnPackageGroup_RestrictionTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-restrictionmode
	//
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-repositories
	//
	Repositories *[]*string `field:"optional" json:"repositories" yaml:"repositories"`
}

