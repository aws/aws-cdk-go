package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDocument`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var content interface{}
//
//   cfnDocumentProps := &cfnDocumentProps{
//   	content: content,
//
//   	// the properties below are optional
//   	attachments: []interface{}{
//   		&attachmentsSourceProperty{
//   			key: jsii.String("key"),
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	documentFormat: jsii.String("documentFormat"),
//   	documentType: jsii.String("documentType"),
//   	name: jsii.String("name"),
//   	requires: []interface{}{
//   		&documentRequiresProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	updateMethod: jsii.String("updateMethod"),
//   	versionName: jsii.String("versionName"),
//   }
//
type CfnDocumentProps struct {
	// The content for the new SSM document in JSON or YAML.
	//
	// > This parameter also supports `String` data types.
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// A list of key-value pairs that describe attachments to a version of a document.
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// Specify the document format for the request.
	//
	// JSON is the default format.
	DocumentFormat *string `field:"optional" json:"documentFormat" yaml:"documentFormat"`
	// The type of document to create.
	//
	// *Allowed Values* : `ApplicationConfigurationSchema` | `Automation` | `Automation.ChangeTemplate` | `Command` | `DeploymentStrategy` | `Package` | `Policy` | `Session`
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// A name for the SSM document.
	//
	// > You can't use the following strings as document name prefixes. These are reserved by AWS for use as document name prefixes:
	// >
	// > - `aws-`
	// > - `amazon`
	// > - `amzn`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of SSM documents required by a document.
	//
	// This parameter is used exclusively by AWS AppConfig . When a user creates an AWS AppConfig configuration in an SSM document, the user must also specify a required document for validation purposes. In this case, an `ApplicationConfiguration` document requires an `ApplicationConfigurationSchema` document for validation purposes. For more information, see [What is AWS AppConfig ?](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html) in the *AWS AppConfig User Guide* .
	Requires interface{} `field:"optional" json:"requires" yaml:"requires"`
	// AWS CloudFormation resource tags to apply to the document.
	//
	// Use tags to help you identify and categorize resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specify a target type to define the kinds of resources the document can run on.
	//
	// For example, to run a document on EC2 instances, specify the following value: `/AWS::EC2::Instance` . If you specify a value of '/' the document can run on all types of resources. If you don't specify a value, the document can't run on any resources. For a list of valid resource types, see [AWS resource and property types reference](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) in the *AWS CloudFormation User Guide* .
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// If the document resource you specify in your template already exists, this parameter determines whether a new version of the existing document is created, or the existing document is replaced.
	//
	// `Replace` is the default method. If you specify `NewVersion` for the `UpdateMethod` parameter, and the `Name` of the document does not match an existing resource, a new document is created. When you specify `NewVersion` , the default version of the document is changed to the newly created version.
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
	// An optional field specifying the version of the artifact you are creating with the document.
	//
	// For example, `Release12.1` . This value is unique across all versions of a document, and can't be changed.
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

