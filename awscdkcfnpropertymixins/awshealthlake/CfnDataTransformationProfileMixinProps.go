package awshealthlake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataTransformationProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDataTransformationProfileMixinProps := &CfnDataTransformationProfileMixinProps{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ProfileDescription: jsii.String("profileDescription"),
//   	ProfileName: jsii.String("profileName"),
//   	Source: &SourceProperty{
//   		ExistingVersionedProfileId: &ExistingVersionedProfileSourceProperty{
//   			ProfileId: jsii.String("profileId"),
//   			Version: jsii.Number(123),
//   		},
//   		ProfileMapping: &ProfileMappingSourceProperty{
//   			ProfileMapping: map[string]*string{
//   				"profileMappingKey": jsii.String("profileMapping"),
//   			},
//   		},
//   		StarterProfile: &StarterProfileSourceProperty{
//   			StarterProfileName: jsii.String("starterProfileName"),
//   		},
//   	},
//   	SourceFormat: jsii.String("sourceFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html
//
type CfnDataTransformationProfileMixinProps struct {
	// The identifier (key ID or ARN) of a customer-managed KMS key used to encrypt the profile's template content at rest.
	//
	// If omitted, an AWS owned key is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A human-readable description of the profile's purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-profiledescription
	//
	ProfileDescription *string `field:"optional" json:"profileDescription" yaml:"profileDescription"`
	// The human-readable name of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// The source from which to create the profile's initial template content.
	//
	// Exactly one of the members must be specified. Use StarterProfile (C-CDA only), ProfileMapping (C-CDA or CSV), or ExistingVersionedProfileId to clone an existing profile. Each produces a published profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The source format that this profile converts from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-sourceformat
	//
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// An array of key-value pairs to apply to this profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-datatransformationprofile.html#cfn-healthlake-datatransformationprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

