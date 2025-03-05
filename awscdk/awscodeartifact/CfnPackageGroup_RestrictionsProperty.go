package awscodeartifact


// Contains information about the configured restrictions of the origin controls of a package group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictionsProperty := &RestrictionsProperty{
//   	ExternalUpstream: &RestrictionTypeProperty{
//   		RestrictionMode: jsii.String("restrictionMode"),
//
//   		// the properties below are optional
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   	},
//   	InternalUpstream: &RestrictionTypeProperty{
//   		RestrictionMode: jsii.String("restrictionMode"),
//
//   		// the properties below are optional
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   	},
//   	Publish: &RestrictionTypeProperty{
//   		RestrictionMode: jsii.String("restrictionMode"),
//
//   		// the properties below are optional
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html
//
type CfnPackageGroup_RestrictionsProperty struct {
	// The package group origin restriction setting for external, upstream repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-externalupstream
	//
	ExternalUpstream interface{} `field:"optional" json:"externalUpstream" yaml:"externalUpstream"`
	// The package group origin restriction setting for internal, upstream repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-internalupstream
	//
	InternalUpstream interface{} `field:"optional" json:"internalUpstream" yaml:"internalUpstream"`
	// The package group origin restriction setting for publishing packages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-publish
	//
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
}

