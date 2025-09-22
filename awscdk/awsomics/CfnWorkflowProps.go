package awsomics


// Properties for defining a `CfnWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkflowProps := &CfnWorkflowProps{
//   	Accelerators: jsii.String("accelerators"),
//   	DefinitionRepository: &DefinitionRepositoryProperty{
//   		ConnectionArn: jsii.String("connectionArn"),
//   		ExcludeFilePatterns: []*string{
//   			jsii.String("excludeFilePatterns"),
//   		},
//   		FullRepositoryId: jsii.String("fullRepositoryId"),
//   		SourceReference: &SourceReferenceProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	Main: jsii.String("main"),
//   	Name: jsii.String("name"),
//   	ParameterTemplate: map[string]interface{}{
//   		"parameterTemplateKey": &WorkflowParameterProperty{
//   			"description": jsii.String("description"),
//   			"optional": jsii.Boolean(false),
//   		},
//   	},
//   	ParameterTemplatePath: jsii.String("parameterTemplatePath"),
//   	ReadmeMarkdown: jsii.String("readmeMarkdown"),
//   	ReadmePath: jsii.String("readmePath"),
//   	ReadmeUri: jsii.String("readmeUri"),
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkflowBucketOwnerId: jsii.String("workflowBucketOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html
//
type CfnWorkflowProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-accelerators
	//
	Accelerators *string `field:"optional" json:"accelerators" yaml:"accelerators"`
	// Contains information about a source code repository that hosts the workflow definition files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-definitionrepository
	//
	DefinitionRepository interface{} `field:"optional" json:"definitionRepository" yaml:"definitionRepository"`
	// The URI of a definition for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-definitionuri
	//
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// The parameter's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An engine for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The path of the main definition file for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-main
	//
	Main *string `field:"optional" json:"main" yaml:"main"`
	// The workflow's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The workflow's parameter template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-parametertemplate
	//
	ParameterTemplate interface{} `field:"optional" json:"parameterTemplate" yaml:"parameterTemplate"`
	// Path to the primary workflow parameter template JSON file inside the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-parametertemplatepath
	//
	ParameterTemplatePath *string `field:"optional" json:"parameterTemplatePath" yaml:"parameterTemplatePath"`
	// The markdown content for the workflow's README file.
	//
	// This provides documentation and usage information for users of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-readmemarkdown
	//
	ReadmeMarkdown *string `field:"optional" json:"readmeMarkdown" yaml:"readmeMarkdown"`
	// The path to the workflow README markdown file within the repository.
	//
	// This file provides documentation and usage information for the workflow. If not specified, the README.md file from the root directory of the repository will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-readmepath
	//
	ReadmePath *string `field:"optional" json:"readmePath" yaml:"readmePath"`
	// The S3 URI of the README file for the workflow.
	//
	// This file provides documentation and usage information for the workflow. The S3 URI must begin with s3://USER-OWNED-BUCKET/. The requester must have access to the S3 bucket and object. The max README content length is 500 KiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-readmeuri
	//
	ReadmeUri *string `field:"optional" json:"readmeUri" yaml:"readmeUri"`
	// The default static storage capacity (in gibibytes) for runs that use this workflow or workflow version.
	//
	// The `storageCapacity` can be overwritten at run time. The storage capacity is not required for runs with a `DYNAMIC` storage type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-storagecapacity
	//
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Tags for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional workflow bucket owner ID to verify the workflow bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html#cfn-omics-workflow-workflowbucketownerid
	//
	WorkflowBucketOwnerId *string `field:"optional" json:"workflowBucketOwnerId" yaml:"workflowBucketOwnerId"`
}

