package awscloudformation


// Properties for CfnChangeSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnChangeSetMixinProps := &CfnChangeSetMixinProps{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	ChangeSetName: jsii.String("changeSetName"),
//   	ChangeSetType: jsii.String("changeSetType"),
//   	DeploymentMode: jsii.String("deploymentMode"),
//   	Description: jsii.String("description"),
//   	ImportExistingResources: jsii.Boolean(false),
//   	IncludeNestedStacks: jsii.Boolean(false),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	OnStackFailure: jsii.String("onStackFailure"),
//   	RoleArn: jsii.String("roleArn"),
//   	StackName: jsii.String("stackName"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateUrl: jsii.String("templateUrl"),
//   	UsePreviousTemplate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html
//
type CfnChangeSetMixinProps struct {
	// The capabilities that are allowed in the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The name of the change set.
	//
	// Must be unique among all change sets associated with the specified stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-changesetname
	//
	ChangeSetName *string `field:"optional" json:"changeSetName" yaml:"changeSetName"`
	// The type of change set operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-changesettype
	//
	ChangeSetType *string `field:"optional" json:"changeSetType" yaml:"changeSetType"`
	// Determines how CloudFormation handles configuration drift during deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-deploymentmode
	//
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// A description to help you identify this change set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if the change set imports resources that already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-importexistingresources
	//
	ImportExistingResources interface{} `field:"optional" json:"importExistingResources" yaml:"importExistingResources"`
	// Creates a change set for all nested stacks specified in the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-includenestedstacks
	//
	IncludeNestedStacks interface{} `field:"optional" json:"includeNestedStacks" yaml:"includeNestedStacks"`
	// The ARNs of Amazon SNS topics that CloudFormation associates with the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-notificationarns
	//
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Determines what action will be taken if stack creation fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-onstackfailure
	//
	OnStackFailure *string `field:"optional" json:"onStackFailure" yaml:"onStackFailure"`
	// The ARN of an IAM role that CloudFormation assumes when executing the change set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name or unique ID of the stack for which you are creating a change set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-stackname
	//
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Key-value pairs to associate with the change set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-tags
	//
	Tags *[]*CfnChangeSetPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// A structure that contains the body of the revised template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// The URL of the file that contains the revised template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-templateurl
	//
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Whether to reuse the template associated with the stack to create the change set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-changeset.html#cfn-cloudformation-changeset-useprevioustemplate
	//
	UsePreviousTemplate interface{} `field:"optional" json:"usePreviousTemplate" yaml:"usePreviousTemplate"`
}

