package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a knowledge base as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - Name – Specify a name for the knowledge base.
// - RoleArn – Specify the Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base. For more information, see [Create a service role for Knowledge base for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/kb-permissions.html) .
// - KnowledgeBaseConfiguration – Specify the embeddings configuration of the knowledge base. The following sub-properties are required:
//
// - Type – Specify the value `VECTOR` .
// - StorageConfiguration – Specify information about the vector store in which the data source is stored. The following sub-properties are required:
//
// - Type – Specify the vector store service that you are using.
//
// > Redis Enterprise Cloud vector stores are currently unsupported in CloudFormation .
//
// For more information about using knowledge bases in Amazon Bedrock , see [Knowledge base for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnKnowledgeBaseLogsMixin := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html
//
type CfnKnowledgeBaseLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKnowledgeBaseLogsMixin
type jsiiProxy_CfnKnowledgeBaseLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnKnowledgeBaseLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBaseLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Bedrock::KnowledgeBase`.
func NewCfnKnowledgeBaseLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnKnowledgeBaseLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnKnowledgeBaseLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKnowledgeBaseLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Bedrock::KnowledgeBase`.
func NewCfnKnowledgeBaseLogsMixin_Override(c CfnKnowledgeBaseLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnKnowledgeBaseLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKnowledgeBaseLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKnowledgeBaseLogsMixin_APPLICATION_LOGS() CfnKnowledgeBaseApplicationLogs {
	_init_.Initialize()
	var returns CfnKnowledgeBaseApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnKnowledgeBaseLogsMixin_RUNTIME_LOGS() CfnKnowledgeBaseRuntimeLogs {
	_init_.Initialize()
	var returns CfnKnowledgeBaseRuntimeLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		"RUNTIME_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKnowledgeBaseLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKnowledgeBaseLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

