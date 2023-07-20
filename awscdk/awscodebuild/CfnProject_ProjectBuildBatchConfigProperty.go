package awscodebuild


// Contains configuration information about a batch build project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectBuildBatchConfigProperty := &ProjectBuildBatchConfigProperty{
//   	BatchReportMode: jsii.String("batchReportMode"),
//   	CombineArtifacts: jsii.Boolean(false),
//   	Restrictions: &BatchRestrictionsProperty{
//   		ComputeTypesAllowed: []*string{
//   			jsii.String("computeTypesAllowed"),
//   		},
//   		MaximumBuildsAllowed: jsii.Number(123),
//   	},
//   	ServiceRole: jsii.String("serviceRole"),
//   	TimeoutInMins: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html
//
type CfnProject_ProjectBuildBatchConfigProperty struct {
	// Specifies how build status reports are sent to the source provider for the batch build.
	//
	// This property is only used when the source provider for your project is Bitbucket, GitHub, or GitHub Enterprise, and your project is configured to report build statuses to the source provider.
	//
	// - **REPORT_AGGREGATED_BATCH** - (Default) Aggregate all of the build statuses into a single status report.
	// - **REPORT_INDIVIDUAL_BUILDS** - Send a separate status report for each individual build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html#cfn-codebuild-project-projectbuildbatchconfig-batchreportmode
	//
	BatchReportMode *string `field:"optional" json:"batchReportMode" yaml:"batchReportMode"`
	// Specifies if the build artifacts for the batch build should be combined into a single artifact location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html#cfn-codebuild-project-projectbuildbatchconfig-combineartifacts
	//
	CombineArtifacts interface{} `field:"optional" json:"combineArtifacts" yaml:"combineArtifacts"`
	// A `BatchRestrictions` object that specifies the restrictions for the batch build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html#cfn-codebuild-project-projectbuildbatchconfig-restrictions
	//
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Specifies the service role ARN for the batch build project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html#cfn-codebuild-project-projectbuildbatchconfig-servicerole
	//
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Specifies the maximum amount of time, in minutes, that the batch build must be completed in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectbuildbatchconfig.html#cfn-codebuild-project-projectbuildbatchconfig-timeoutinmins
	//
	TimeoutInMins *float64 `field:"optional" json:"timeoutInMins" yaml:"timeoutInMins"`
}

