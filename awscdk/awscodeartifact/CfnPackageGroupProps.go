package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPackageGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageGroupProps := &CfnPackageGroupProps{
//   	DomainName: jsii.String("domainName"),
//   	Pattern: jsii.String("pattern"),
//
//   	// the properties below are optional
//   	ContactInfo: jsii.String("contactInfo"),
//   	Description: jsii.String("description"),
//   	DomainOwner: jsii.String("domainOwner"),
//   	OriginConfiguration: &OriginConfigurationProperty{
//   		Restrictions: &RestrictionsProperty{
//   			ExternalUpstream: &RestrictionTypeProperty{
//   				RestrictionMode: jsii.String("restrictionMode"),
//
//   				// the properties below are optional
//   				Repositories: []*string{
//   					jsii.String("repositories"),
//   				},
//   			},
//   			InternalUpstream: &RestrictionTypeProperty{
//   				RestrictionMode: jsii.String("restrictionMode"),
//
//   				// the properties below are optional
//   				Repositories: []*string{
//   					jsii.String("repositories"),
//   				},
//   			},
//   			Publish: &RestrictionTypeProperty{
//   				RestrictionMode: jsii.String("restrictionMode"),
//
//   				// the properties below are optional
//   				Repositories: []*string{
//   					jsii.String("repositories"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html
//
type CfnPackageGroupProps struct {
	// The domain that contains the package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The pattern of the package group.
	//
	// The pattern determines which packages are associated with the package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-pattern
	//
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The contact information of the package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-contactinfo
	//
	ContactInfo *string `field:"optional" json:"contactInfo" yaml:"contactInfo"`
	// The description of the package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The 12-digit account number of the AWS account that owns the domain.
	//
	// It does not include dashes or spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-domainowner
	//
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// Details about the package origin configuration of a package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-originconfiguration
	//
	OriginConfiguration interface{} `field:"optional" json:"originConfiguration" yaml:"originConfiguration"`
	// An array of key-value pairs to apply to the package group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-packagegroup.html#cfn-codeartifact-packagegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

