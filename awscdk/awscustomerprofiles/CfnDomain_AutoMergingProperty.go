package awscustomerprofiles


// Configuration information about the auto-merging process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoMergingProperty := &AutoMergingProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	ConflictResolution: &ConflictResolutionProperty{
//   		ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//
//   		// the properties below are optional
//   		SourceName: jsii.String("sourceName"),
//   	},
//   	Consolidation: &ConsolidationProperty{
//   		MatchingAttributesList: []interface{}{
//   			[]*string{
//   				jsii.String("matchingAttributesList"),
//   			},
//   		},
//   	},
//   	MinAllowedConfidenceScoreForMerging: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-automerging.html
//
type CfnDomain_AutoMergingProperty struct {
	// The flag that enables the auto-merging of duplicate profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-automerging.html#cfn-customerprofiles-domain-automerging-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Determines how the auto-merging process should resolve conflicts between different profiles.
	//
	// For example, if Profile A and Profile B have the same `FirstName` and `LastName` , `ConflictResolution` specifies which `EmailAddress` should be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-automerging.html#cfn-customerprofiles-domain-automerging-conflictresolution
	//
	ConflictResolution interface{} `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// A list of matching attributes that represent matching criteria.
	//
	// If two profiles meet at least one of the requirements in the matching attributes list, they will be merged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-automerging.html#cfn-customerprofiles-domain-automerging-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// A number between 0 and 1 that represents the minimum confidence score required for profiles within a matching group to be merged during the auto-merge process.
	//
	// A higher score means that a higher similarity is required to merge profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-automerging.html#cfn-customerprofiles-domain-automerging-minallowedconfidencescoreformerging
	//
	MinAllowedConfidenceScoreForMerging *float64 `field:"optional" json:"minAllowedConfidenceScoreForMerging" yaml:"minAllowedConfidenceScoreForMerging"`
}

