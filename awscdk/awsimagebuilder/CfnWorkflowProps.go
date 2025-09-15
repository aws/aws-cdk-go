package awsimagebuilder


// Properties for defining a `CfnWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkflowProps := &CfnWorkflowProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	ChangeDescription: jsii.String("changeDescription"),
//   	Data: jsii.String("data"),
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html
//
type CfnWorkflowProps struct {
	// The name of the workflow to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The phase in the image build process for which the workflow resource is responsible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The semantic version of this workflow resource. The semantic version syntax adheres to the following rules.
	//
	// > The semantic version has four nodes: <major>.<minor>.<patch>/<build>. You can assign values for the first three, and can filter on all of them.
	// >
	// > *Assignment:* For the first three nodes you can assign any positive integer value, including zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the build number to the fourth node.
	// >
	// > *Patterns:* You can use any numeric pattern that adheres to the assignment requirements for the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// Describes what change has been made in this version of the workflow, or what makes this version different from other versions of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-changedescription
	//
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// Contains the UTF-8 encoded YAML document content for the workflow.
	//
	// Alternatively, you can specify the `uri` of a YAML document file stored in Amazon S3. However, you cannot specify both properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Describes the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this workflow resource.
	//
	// This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Tags that apply to the workflow resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The `uri` of a YAML component document file.
	//
	// This must be an S3 URL ( `s3://bucket/key` ), and the requester must have permission to access the S3 bucket it points to. If you use Amazon S3, you can specify component content up to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component `data` property. You cannot specify both properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

