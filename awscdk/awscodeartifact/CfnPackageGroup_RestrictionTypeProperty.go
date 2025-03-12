package awscodeartifact


// The `RestrictionType` property type specifies the package group origin configuration restriction mode, and the repositories when the `RestrictionMode` is set to `ALLOW_SPECIFIC_REPOSITORIES` .
//
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
	// The package group origin restriction setting.
	//
	// When the value is `INHERIT` , the value is set to the value of the first parent package group which does not have a value of `INHERIT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-restrictionmode
	//
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// The repositories to add to the allowed repositories list.
	//
	// The allowed repositories list is used when the `RestrictionMode` is set to `ALLOW_SPECIFIC_REPOSITORIES` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-repositories
	//
	Repositories *[]*string `field:"optional" json:"repositories" yaml:"repositories"`
}

