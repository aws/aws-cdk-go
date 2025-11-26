package previewawscustomerprofilesmixins


// The process of matching duplicate profiles.
//
// If `Matching = true` , Amazon Connect Customer Profiles starts a weekly batch process called *Identity Resolution Job* . If you do not specify a date and time for the *Identity Resolution Job* to run, by default it runs every Saturday at 12AM UTC to detect duplicate profiles in your domains. After the *Identity Resolution Job* completes, use the `GetMatches` API to return and review the results. Or, if you have configured `ExportingConfig` in the `MatchingRequest` , you can download the results from S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matchingProperty := &MatchingProperty{
//   	AutoMerging: &AutoMergingProperty{
//   		ConflictResolution: &ConflictResolutionProperty{
//   			ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//   			SourceName: jsii.String("sourceName"),
//   		},
//   		Consolidation: &ConsolidationProperty{
//   			MatchingAttributesList: []interface{}{
//   				[]*string{
//   					jsii.String("matchingAttributesList"),
//   				},
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		MinAllowedConfidenceScoreForMerging: jsii.Number(123),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	ExportingConfig: &ExportingConfigProperty{
//   		S3Exporting: &S3ExportingConfigProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
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
type CfnDomainPropsMixin_MatchingProperty struct {
	// Configuration information about the auto-merging process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-automerging
	//
	AutoMerging interface{} `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// The flag that enables the matching process of duplicate profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 location where Identity Resolution Jobs write result files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-exportingconfig
	//
	ExportingConfig interface{} `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// The day and time when do you want to start the Identity Resolution Job every week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matching.html#cfn-customerprofiles-domain-matching-jobschedule
	//
	JobSchedule interface{} `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}

