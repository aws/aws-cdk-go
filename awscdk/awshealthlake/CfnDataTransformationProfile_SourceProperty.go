package awshealthlake


// The source from which to create the profile's initial template content.
//
// Exactly one of the members must be specified. Use StarterProfile (C-CDA only), ProfileMapping (C-CDA or CSV), or ExistingVersionedProfileId to clone an existing profile. Each produces a published profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	ExistingVersionedProfileId: &ExistingVersionedProfileSourceProperty{
//   		ProfileId: jsii.String("profileId"),
//   		Version: jsii.Number(123),
//   	},
//   	ProfileMapping: &ProfileMappingSourceProperty{
//   		ProfileMapping: map[string]*string{
//   			"profileMappingKey": jsii.String("profileMapping"),
//   		},
//   	},
//   	StarterProfile: &StarterProfileSourceProperty{
//   		StarterProfileName: jsii.String("starterProfileName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-source.html
//
type CfnDataTransformationProfile_SourceProperty struct {
	// Create the profile by cloning a specific version of an existing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-source.html#cfn-healthlake-datatransformationprofile-source-existingversionedprofileid
	//
	ExistingVersionedProfileId interface{} `field:"optional" json:"existingVersionedProfileId" yaml:"existingVersionedProfileId"`
	// Create the profile from raw Velocity template mapping content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-source.html#cfn-healthlake-datatransformationprofile-source-profilemapping
	//
	ProfileMapping interface{} `field:"optional" json:"profileMapping" yaml:"profileMapping"`
	// Create the profile from a predefined starter profile of transformation templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-source.html#cfn-healthlake-datatransformationprofile-source-starterprofile
	//
	StarterProfile interface{} `field:"optional" json:"starterProfile" yaml:"starterProfile"`
}

