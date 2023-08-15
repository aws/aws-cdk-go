package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties for `ReportGroup`.
//
// Example:
//   var source source
//
//
//   // create a new ReportGroup
//   reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"), &ReportGroupProps{
//   	Type: codebuild.ReportGroupType_CODE_COVERAGE,
//   })
//
//   project := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Source: Source,
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
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
type ReportGroupProps struct {
	// An optional S3 bucket to export the reports to.
	// Default: - the reports will not be exported.
	//
	ExportBucket awss3.IBucket `field:"optional" json:"exportBucket" yaml:"exportBucket"`
	// What to do when this resource is deleted from a stack.
	//
	// As CodeBuild does not allow deleting a ResourceGroup that has reports inside of it,
	// this is set to retain the resource by default.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The physical name of the report group.
	// Default: - CloudFormation-generated name.
	//
	ReportGroupName *string `field:"optional" json:"reportGroupName" yaml:"reportGroupName"`
	// The type of report group. This can be one of the following values:.
	//
	// - **TEST** - The report group contains test reports.
	// - **CODE_COVERAGE** - The report group contains code coverage reports.
	// Default: TEST.
	//
	Type ReportGroupType `field:"optional" json:"type" yaml:"type"`
	// Whether to output the report files into the export bucket as-is, or create a ZIP from them before doing the export.
	//
	// Ignored if `exportBucket` has not been provided.
	// Default: - false (the files will not be ZIPped).
	//
	ZipExport *bool `field:"optional" json:"zipExport" yaml:"zipExport"`
}

