package previewawscodeartifactmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originConfigurationProperty := &OriginConfigurationProperty{
//   	Restrictions: &RestrictionsProperty{
//   		ExternalUpstream: &RestrictionTypeProperty{
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   			RestrictionMode: jsii.String("restrictionMode"),
//   		},
//   		InternalUpstream: &RestrictionTypeProperty{
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   			RestrictionMode: jsii.String("restrictionMode"),
//   		},
//   		Publish: &RestrictionTypeProperty{
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   			RestrictionMode: jsii.String("restrictionMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html
//
type CfnPackageGroupPropsMixin_OriginConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html#cfn-codeartifact-packagegroup-originconfiguration-restrictions
	//
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

