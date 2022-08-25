package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties for {@link S3Artifacts}.
//
// Example:
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	artifacts: codebuild.artifacts.s3(&s3ArtifactsProps{
//   		bucket: bucket,
//   		includeBuildId: jsii.Boolean(false),
//   		packageZip: jsii.Boolean(true),
//   		path: jsii.String("another/path"),
//   		identifier: jsii.String("AddArtifact1"),
//   	}),
//   })
//
type S3ArtifactsProps struct {
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The name of the output bucket.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// If this is false, build output will not be encrypted.
	//
	// This is useful if the artifact to publish a static website or sharing content with others.
	Encryption *bool `field:"optional" json:"encryption" yaml:"encryption"`
	// Indicates if the build ID should be included in the path.
	//
	// If this is set to true,
	// then the build artifact will be stored in "<path>/<build-id>/<name>".
	IncludeBuildId *bool `field:"optional" json:"includeBuildId" yaml:"includeBuildId"`
	// The name of the build output ZIP file or folder inside the bucket.
	//
	// The full S3 object key will be "<path>/<build-id>/<name>" or
	// "<path>/<name>" depending on whether `includeBuildId` is set to true.
	//
	// If not set, `overrideArtifactName` will be set and the name from the
	// buildspec will be used instead.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If this is true, all build output will be packaged into a single .zip file. Otherwise, all files will be uploaded to <path>/<name>.
	PackageZip *bool `field:"optional" json:"packageZip" yaml:"packageZip"`
	// The path inside of the bucket for the build output .zip file or folder. If a value is not specified, then build output will be stored at the root of the bucket (or under the <build-id> directory if `includeBuildId` is set to true).
	Path *string `field:"optional" json:"path" yaml:"path"`
}

