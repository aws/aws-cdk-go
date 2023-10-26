package awscustomerprofiles


// The process of matching duplicate profiles.
//
// If Matching = true, Amazon Connect Customer Profiles starts a weekly batch process called Identity Resolution Job. If you do not specify a date and time for Identity Resolution Job to run, by default it runs every Saturday at 12AM UTC to detect duplicate profiles in your domains. After the Identity Resolution Job completes, use the GetMatches API to return and review the results. Or, if you have configured ExportingConfig in the MatchingRequest, you can download the results from S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchingProperty := &MatchingProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AutoMerging: &AutoMergingProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ConflictResolution: &ConflictResolutionProperty{
//   			ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//
//   			// the properties below are optional
//   			SourceName: jsii.String("sourceName"),
//   		},
//   		Consolidation: &ConsolidationProperty{
//   			MatchingAttributesList: []interface{}{
//   				[]*string{
//   					jsii.String("matchingAttributesList"),
//   				},
//   			},
//   		},
//   		MinAllowedConfidenceScoreForMerging: jsii.Number(123),
//   	},
//   	ExportingConfig: &ExportingConfigProperty{
//   		S3Exporting: &S3ExportingConfigProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//
//   			// the properties below are optional
//   			S3KeyName: jsii.String("s3KeyName"),
//   		},
//   	},
//   	JobSchedule: &JobScheduleProperty{
//   		DayOfTheWeek: jsii.String("dayOfTheWeek"),
//   		Time: jsii.String("time"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html
//
type CfnDomain_MatchingProperty struct {
	// The flag that enables the matching process of duplicate profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Configuration information about the auto-merging process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-automerging
	//
	AutoMerging interface{} `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// Configuration information for exporting Identity Resolution results, for example, to an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-exportingconfig
	//
	ExportingConfig interface{} `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// The day and time when do you want to start the Identity Resolution Job every week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-jobschedule
	//
	JobSchedule interface{} `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}

