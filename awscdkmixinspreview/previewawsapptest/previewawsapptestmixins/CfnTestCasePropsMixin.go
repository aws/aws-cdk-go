package previewawsapptestmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapptest/previewawsapptestmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a test case for an application.
//
// For more information about test cases, see [Test cases](https://docs.aws.amazon.com/m2/latest/userguide/testing-test-cases.html) and [Application Testing concepts](https://docs.aws.amazon.com/m2/latest/userguide/concepts-apptest.html) in the *AWS Mainframe Modernization User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTestCasePropsMixin := awscdkmixinspreview.Mixins.NewCfnTestCasePropsMixin(&CfnTestCaseMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Steps: []interface{}{
//   		&StepProperty{
//   			Action: &StepActionProperty{
//   				CompareAction: &CompareActionProperty{
//   					Input: &InputProperty{
//   						File: &InputFileProperty{
//   							FileMetadata: &FileMetadataProperty{
//   								DatabaseCdc: &DatabaseCDCProperty{
//   									SourceMetadata: &SourceDatabaseMetadataProperty{
//   										CaptureTool: jsii.String("captureTool"),
//   										Type: jsii.String("type"),
//   									},
//   									TargetMetadata: &TargetDatabaseMetadataProperty{
//   										CaptureTool: jsii.String("captureTool"),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   								DataSets: []interface{}{
//   									&DataSetProperty{
//   										Ccsid: jsii.String("ccsid"),
//   										Format: jsii.String("format"),
//   										Length: jsii.Number(123),
//   										Name: jsii.String("name"),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							SourceLocation: jsii.String("sourceLocation"),
//   							TargetLocation: jsii.String("targetLocation"),
//   						},
//   					},
//   					Output: &OutputProperty{
//   						File: &OutputFileProperty{
//   							FileLocation: jsii.String("fileLocation"),
//   						},
//   					},
//   				},
//   				MainframeAction: &MainframeActionProperty{
//   					ActionType: &MainframeActionTypeProperty{
//   						Batch: &BatchProperty{
//   							BatchJobName: jsii.String("batchJobName"),
//   							BatchJobParameters: map[string]*string{
//   								"batchJobParametersKey": jsii.String("batchJobParameters"),
//   							},
//   							ExportDataSetNames: []*string{
//   								jsii.String("exportDataSetNames"),
//   							},
//   						},
//   						Tn3270: &TN3270Property{
//   							ExportDataSetNames: []*string{
//   								jsii.String("exportDataSetNames"),
//   							},
//   							Script: &ScriptProperty{
//   								ScriptLocation: jsii.String("scriptLocation"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					Properties: &MainframeActionPropertiesProperty{
//   						DmsTaskArn: jsii.String("dmsTaskArn"),
//   					},
//   					Resource: jsii.String("resource"),
//   				},
//   				ResourceAction: &ResourceActionProperty{
//   					CloudFormationAction: &CloudFormationActionProperty{
//   						ActionType: jsii.String("actionType"),
//   						Resource: jsii.String("resource"),
//   					},
//   					M2ManagedApplicationAction: &M2ManagedApplicationActionProperty{
//   						ActionType: jsii.String("actionType"),
//   						Properties: &M2ManagedActionPropertiesProperty{
//   							ForceStop: jsii.Boolean(false),
//   							ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   						},
//   						Resource: jsii.String("resource"),
//   					},
//   					M2NonManagedApplicationAction: &M2NonManagedApplicationActionProperty{
//   						ActionType: jsii.String("actionType"),
//   						Resource: jsii.String("resource"),
//   					},
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html
//
type CfnTestCasePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTestCaseMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTestCasePropsMixin
type jsiiProxy_CfnTestCasePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTestCasePropsMixin) Props() *CfnTestCaseMixinProps {
	var returns *CfnTestCaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTestCasePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppTest::TestCase`.
func NewCfnTestCasePropsMixin(props *CfnTestCaseMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTestCasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTestCasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTestCasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppTest::TestCase`.
func NewCfnTestCasePropsMixin_Override(c CfnTestCasePropsMixin, props *CfnTestCaseMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTestCasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTestCasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTestCasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTestCasePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTestCasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

