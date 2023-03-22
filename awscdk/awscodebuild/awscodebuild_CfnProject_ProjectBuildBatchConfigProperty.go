package awscodebuild


// Contains configuration information about a batch build project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectBuildBatchConfigProperty := &projectBuildBatchConfigProperty{
//   	batchReportMode: jsii.String("batchReportMode"),
//   	combineArtifacts: jsii.Boolean(false),
//   	restrictions: &batchRestrictionsProperty{
//   		computeTypesAllowed: []*string{
//   			jsii.String("computeTypesAllowed"),
//   		},
//   		maximumBuildsAllowed: jsii.Number(123),
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//   	timeoutInMins: jsii.Number(123),
//   }
//
type CfnProject_ProjectBuildBatchConfigProperty struct {
	// Specifies how build status reports are sent to the source provider for the batch build.
	//
	// This property is only used when the source provider for your project is Bitbucket, GitHub, or GitHub Enterprise, and your project is configured to report build statuses to the source provider.
	//
	// - **REPORT_AGGREGATED_BATCH** - (Default) Aggregate all of the build statuses into a single status report.
	// - **REPORT_INDIVIDUAL_BUILDS** - Send a separate status report for each individual build.
	BatchReportMode *string `field:"optional" json:"batchReportMode" yaml:"batchReportMode"`
	// Specifies if the build artifacts for the batch build should be combined into a single artifact location.
	CombineArtifacts interface{} `field:"optional" json:"combineArtifacts" yaml:"combineArtifacts"`
	// A `BatchRestrictions` object that specifies the restrictions for the batch build.
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Specifies the service role ARN for the batch build project.
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Specifies the maximum amount of time, in minutes, that the batch build must be completed in.
	TimeoutInMins *float64 `field:"optional" json:"timeoutInMins" yaml:"timeoutInMins"`
}

