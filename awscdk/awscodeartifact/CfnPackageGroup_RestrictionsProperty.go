package awscodeartifact


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-externalupstream
	//
	ExternalUpstream interface{} `field:"optional" json:"externalUpstream" yaml:"externalUpstream"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-internalupstream
	//
	InternalUpstream interface{} `field:"optional" json:"internalUpstream" yaml:"internalUpstream"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html#cfn-codeartifact-packagegroup-restrictions-publish
	//
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
}

