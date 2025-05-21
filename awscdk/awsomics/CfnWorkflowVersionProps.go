package awsomics


// Properties for defining a `CfnWorkflowVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkflowVersionProps := &CfnWorkflowVersionProps{
//   	VersionName: jsii.String("versionName"),
//   	WorkflowId: jsii.String("workflowId"),
//
//   	// the properties below are optional
//   	Accelerators: jsii.String("accelerators"),
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
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkflowBucketOwnerId: jsii.String("workflowBucketOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html
//
type CfnWorkflowVersionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-versionname
	//
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-workflowid
	//
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-accelerators
	//
	Accelerators *string `field:"optional" json:"accelerators" yaml:"accelerators"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-definitionuri
	//
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html#cfn-omics-workflowversion-workflowbucketownerid
	//
	WorkflowBucketOwnerId *string `field:"optional" json:"workflowBucketOwnerId" yaml:"workflowBucketOwnerId"`
}

