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
	// The name of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the workflow denotes whether the workflow is used to build, test, or distribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The version of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-changedescription
	//
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The data of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The tags associated with the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The uri of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflow.html#cfn-imagebuilder-workflow-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

