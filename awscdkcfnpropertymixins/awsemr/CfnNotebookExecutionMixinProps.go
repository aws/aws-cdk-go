package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNotebookExecutionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNotebookExecutionMixinProps := &CfnNotebookExecutionMixinProps{
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	ExecutionEngine: &ExecutionEngineConfigProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	NotebookExecutionName: jsii.String("notebookExecutionName"),
//   	NotebookParams: jsii.String("notebookParams"),
//   	NotebookS3Location: &NotebookS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	OutputNotebookFormat: jsii.String("outputNotebookFormat"),
//   	OutputNotebookS3Location: &OutputNotebookS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html
//
type CfnNotebookExecutionMixinProps struct {
	// The environment variables associated with the notebook execution.
	//
	// Keys must be prefixed with KERNEL_ (except LOG_CONTEXT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Specifies the execution engine (cluster) to run the notebook and perform the notebook execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-executionengine
	//
	ExecutionEngine interface{} `field:"optional" json:"executionEngine" yaml:"executionEngine"`
	// An optional name for the notebook execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-notebookexecutionname
	//
	NotebookExecutionName *string `field:"optional" json:"notebookExecutionName" yaml:"notebookExecutionName"`
	// Input parameters in JSON format passed to the Amazon EMR Notebook at runtime for execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-notebookparams
	//
	NotebookParams *string `field:"optional" json:"notebookParams" yaml:"notebookParams"`
	// The Amazon S3 location that stores the notebook execution input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-notebooks3location
	//
	NotebookS3Location interface{} `field:"optional" json:"notebookS3Location" yaml:"notebookS3Location"`
	// The output format for the notebook execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-outputnotebookformat
	//
	OutputNotebookFormat *string `field:"optional" json:"outputNotebookFormat" yaml:"outputNotebookFormat"`
	// The Amazon S3 location for the notebook execution output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-outputnotebooks3location
	//
	OutputNotebookS3Location interface{} `field:"optional" json:"outputNotebookS3Location" yaml:"outputNotebookS3Location"`
	// A list of tags associated with a notebook execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html#cfn-emr-notebookexecution-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

