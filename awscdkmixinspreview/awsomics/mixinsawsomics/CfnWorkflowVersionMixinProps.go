package mixinsawsomics


// Properties for CfnWorkflowVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkflowVersionMixinProps := &CfnWorkflowVersionMixinProps{
//   	Accelerators: jsii.String("accelerators"),
//   	ContainerRegistryMap: &ContainerRegistryMapProperty{
//   		ImageMappings: []interface{}{
//   			&ImageMappingProperty{
//   				DestinationImage: jsii.String("destinationImage"),
//   				SourceImage: jsii.String("sourceImage"),
//   			},
//   		},
//   		RegistryMappings: []interface{}{
//   			&RegistryMappingProperty{
//   				EcrAccountId: jsii.String("ecrAccountId"),
//   				EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   				UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   				UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   			},
//   		},
//   	},
//   	ContainerRegistryMapUri: jsii.String("containerRegistryMapUri"),
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
//   	VersionName: jsii.String("versionName"),
//   	WorkflowBucketOwnerId: jsii.String("workflowBucketOwnerId"),
//   	WorkflowId: jsii.String("workflowId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html
//
type CfnWorkflowVersionMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-accelerators
	//
	Accelerators *string `field:"optional" json:"accelerators" yaml:"accelerators"`
	// Use a container registry map to specify mappings between the ECR private repository and one or more upstream registries.
	//
	// For more information, see [Container images](https://docs.aws.amazon.com/omics/latest/dev/workflows-ecr.html) in the *AWS HealthOmics User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-containerregistrymap
	//
	ContainerRegistryMap interface{} `field:"optional" json:"containerRegistryMap" yaml:"containerRegistryMap"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-containerregistrymapuri
	//
	ContainerRegistryMapUri *string `field:"optional" json:"containerRegistryMapUri" yaml:"containerRegistryMapUri"`
	// Contains information about a source code repository that hosts the workflow definition files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-definitionrepository
	//
	DefinitionRepository interface{} `field:"optional" json:"definitionRepository" yaml:"definitionRepository"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-definitionuri
	//
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// The description of the workflow version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-main
	//
	Main *string `field:"optional" json:"main" yaml:"main"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-parametertemplate
	//
	ParameterTemplate interface{} `field:"optional" json:"parameterTemplate" yaml:"parameterTemplate"`
	// Path to the primary workflow parameter template JSON file inside the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-parametertemplatepath
	//
	ParameterTemplatePath *string `field:"optional" json:"parameterTemplatePath" yaml:"parameterTemplatePath"`
	// The markdown content for the workflow's README file.
	//
	// This provides documentation and usage information for users of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-readmemarkdown
	//
	ReadmeMarkdown *string `field:"optional" json:"readmeMarkdown" yaml:"readmeMarkdown"`
	// The path to the workflow README markdown file within the repository.
	//
	// This file provides documentation and usage information for the workflow. If not specified, the README.md file from the root directory of the repository will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-readmepath
	//
	ReadmePath *string `field:"optional" json:"readmePath" yaml:"readmePath"`
	// The S3 URI of the README file for the workflow.
	//
	// This file provides documentation and usage information for the workflow. The S3 URI must begin with s3://USER-OWNED-BUCKET/. The requester must have access to the S3 bucket and object. The max README content length is 500 KiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-readmeuri
	//
	ReadmeUri *string `field:"optional" json:"readmeUri" yaml:"readmeUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-storagecapacity
	//
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// A map of resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-workflowbucketownerid
	//
	WorkflowBucketOwnerId *string `field:"optional" json:"workflowBucketOwnerId" yaml:"workflowBucketOwnerId"`
	// The workflow's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-workflowid
	//
	WorkflowId *string `field:"optional" json:"workflowId" yaml:"workflowId"`
}

