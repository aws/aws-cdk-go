package previewawscodeartifactmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   restrictionsProperty := &RestrictionsProperty{
//   	ExternalUpstream: &RestrictionTypeProperty{
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   		RestrictionMode: jsii.String("restrictionMode"),
//   	},
//   	InternalUpstream: &RestrictionTypeProperty{
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   		RestrictionMode: jsii.String("restrictionMode"),
//   	},
//   	Publish: &RestrictionTypeProperty{
//   		Repositories: []*string{
//   			jsii.String("repositories"),
//   		},
//   		RestrictionMode: jsii.String("restrictionMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictions.html
//
type CfnPackageGroupPropsMixin_RestrictionsProperty struct {
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

