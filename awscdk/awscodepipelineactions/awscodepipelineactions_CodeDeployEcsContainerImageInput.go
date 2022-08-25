package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Configuration for replacing a placeholder string in the ECS task definition template file with an image URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   codeDeployEcsContainerImageInput := &codeDeployEcsContainerImageInput{
//   	input: artifact,
//
//   	// the properties below are optional
//   	taskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   }
//
type CodeDeployEcsContainerImageInput struct {
	// The artifact that contains an `imageDetails.json` file with the image URI.
	//
	// The artifact's `imageDetails.json` file must be a JSON file containing an
	// `ImageURI` property.  For example:
	// `{ "ImageURI": "ACCOUNTID.dkr.ecr.us-west-2.amazonaws.com/dk-image-repo@sha256:example3" }`
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The placeholder string in the ECS task definition template file that will be replaced with the image URI.
	//
	// The placeholder string must be surrounded by angle brackets in the template file.
	// For example, if the task definition template file contains a placeholder like
	// `"image": "<PLACEHOLDER>"`, then the `taskDefinitionPlaceholder` value should
	// be `PLACEHOLDER`.
	TaskDefinitionPlaceholder *string `field:"optional" json:"taskDefinitionPlaceholder" yaml:"taskDefinitionPlaceholder"`
}

