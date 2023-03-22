package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
)

// The result of adding actions to the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//   codePipelineActionFactoryResult := &codePipelineActionFactoryResult{
//   	runOrdersConsumed: jsii.Number(123),
//
//   	// the properties below are optional
//   	project: project,
//   }
//
// Experimental.
type CodePipelineActionFactoryResult struct {
	// How many RunOrders were consumed.
	//
	// If you add 1 action, return the value 1 here.
	// Experimental.
	RunOrdersConsumed *float64 `field:"required" json:"runOrdersConsumed" yaml:"runOrdersConsumed"`
	// If a CodeBuild project got created, the project.
	// Experimental.
	Project awscodebuild.IProject `field:"optional" json:"project" yaml:"project"`
}

