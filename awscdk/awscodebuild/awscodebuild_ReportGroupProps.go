package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties for {@link ReportGroup}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   reportGroupProps := &reportGroupProps{
//   	exportBucket: bucket,
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   	reportGroupName: jsii.String("reportGroupName"),
//   	zipExport: jsii.Boolean(false),
//   }
//
// Experimental.
type ReportGroupProps struct {
	// An optional S3 bucket to export the reports to.
	// Experimental.
	ExportBucket awss3.IBucket `field:"optional" json:"exportBucket" yaml:"exportBucket"`
	// What to do when this resource is deleted from a stack.
	//
	// As CodeBuild does not allow deleting a ResourceGroup that has reports inside of it,
	// this is set to retain the resource by default.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The physical name of the report group.
	// Experimental.
	ReportGroupName *string `field:"optional" json:"reportGroupName" yaml:"reportGroupName"`
	// Whether to output the report files into the export bucket as-is, or create a ZIP from them before doing the export.
	//
	// Ignored if {@link exportBucket} has not been provided.
	// Experimental.
	ZipExport *bool `field:"optional" json:"zipExport" yaml:"zipExport"`
}

