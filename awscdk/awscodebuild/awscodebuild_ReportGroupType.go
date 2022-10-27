package awscodebuild


// The type of reports in the report group.
//
// Example:
//   var source source
//
//
//   // create a new ReportGroup
//   reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"), &reportGroupProps{
//   	type: codebuild.reportGroupType_CODE_COVERAGE,
//   })
//
//   project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	source: source,
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		// ...
//   		"reports": map[string]map[string]*string{
//   			reportGroup.reportGroupArn: map[string]*string{
//   				"files": jsii.String("**/*"),
//   				"base-directory": jsii.String("build/coverage-report.xml"),
//   				"file-format": jsii.String("JACOCOXML"),
//   			},
//   		},
//   	}),
//   })
//
type ReportGroupType string

const (
	// The report group contains test reports.
	ReportGroupType_TEST ReportGroupType = "TEST"
	// The report group contains code coverage reports.
	ReportGroupType_CODE_COVERAGE ReportGroupType = "CODE_COVERAGE"
)

