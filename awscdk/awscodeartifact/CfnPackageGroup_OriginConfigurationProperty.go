package awscodeartifact


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originConfigurationProperty := &OriginConfigurationProperty{
//   	Restrictions: &RestrictionsProperty{
//   		ExternalUpstream: &RestrictionTypeProperty{
//   			RestrictionMode: jsii.String("restrictionMode"),
//
//   			// the properties below are optional
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   		},
//   		InternalUpstream: &RestrictionTypeProperty{
//   			RestrictionMode: jsii.String("restrictionMode"),
//
//   			// the properties below are optional
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   		},
//   		Publish: &RestrictionTypeProperty{
//   			RestrictionMode: jsii.String("restrictionMode"),
//
//   			// the properties below are optional
//   			Repositories: []*string{
//   				jsii.String("repositories"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html
//
type CfnPackageGroup_OriginConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html#cfn-codeartifact-packagegroup-originconfiguration-restrictions
	//
	Restrictions interface{} `field:"required" json:"restrictions" yaml:"restrictions"`
}

