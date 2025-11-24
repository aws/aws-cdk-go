package mixinsawsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsglue/mixinsawsglue/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Crawler` resource specifies an AWS Glue crawler.
//
// For more information, see [Cataloging Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) and [Crawler Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-crawling.html#aws-glue-api-crawler-crawling-Crawler) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnCrawlerPropsMixin := awscdkmixinspreview.Mixins.NewCfnCrawlerPropsMixin(&CfnCrawlerMixinProps{
//   	Classifiers: []*string{
//   		jsii.String("classifiers"),
//   	},
//   	Configuration: jsii.String("configuration"),
//   	CrawlerSecurityConfiguration: jsii.String("crawlerSecurityConfiguration"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Description: jsii.String("description"),
//   	LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		UseLakeFormationCredentials: jsii.Boolean(false),
//   	},
//   	Name: jsii.String("name"),
//   	RecrawlPolicy: &RecrawlPolicyProperty{
//   		RecrawlBehavior: jsii.String("recrawlBehavior"),
//   	},
//   	Role: jsii.String("role"),
//   	Schedule: &ScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SchemaChangePolicy: &SchemaChangePolicyProperty{
//   		DeleteBehavior: jsii.String("deleteBehavior"),
//   		UpdateBehavior: jsii.String("updateBehavior"),
//   	},
//   	TablePrefix: jsii.String("tablePrefix"),
//   	Tags: tags,
//   	Targets: &TargetsProperty{
//   		CatalogTargets: []interface{}{
//   			&CatalogTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				EventQueueArn: jsii.String("eventQueueArn"),
//   				Tables: []*string{
//   					jsii.String("tables"),
//   				},
//   			},
//   		},
//   		DeltaTargets: []interface{}{
//   			&DeltaTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				CreateNativeDeltaTable: jsii.Boolean(false),
//   				DeltaTables: []*string{
//   					jsii.String("deltaTables"),
//   				},
//   				WriteManifest: jsii.Boolean(false),
//   			},
//   		},
//   		DynamoDbTargets: []interface{}{
//   			&DynamoDBTargetProperty{
//   				Path: jsii.String("path"),
//   				ScanAll: jsii.Boolean(false),
//   				ScanRate: jsii.Number(123),
//   			},
//   		},
//   		HudiTargets: []interface{}{
//   			&HudiTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				MaximumTraversalDepth: jsii.Number(123),
//   				Paths: []*string{
//   					jsii.String("paths"),
//   				},
//   			},
//   		},
//   		IcebergTargets: []interface{}{
//   			&IcebergTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				MaximumTraversalDepth: jsii.Number(123),
//   				Paths: []*string{
//   					jsii.String("paths"),
//   				},
//   			},
//   		},
//   		JdbcTargets: []interface{}{
//   			&JdbcTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				EnableAdditionalMetadata: []*string{
//   					jsii.String("enableAdditionalMetadata"),
//   				},
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		MongoDbTargets: []interface{}{
//   			&MongoDBTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		S3Targets: []interface{}{
//   			&S3TargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				EventQueueArn: jsii.String("eventQueueArn"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   				SampleSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html
//
type CfnCrawlerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCrawlerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCrawlerPropsMixin
type jsiiProxy_CfnCrawlerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCrawlerPropsMixin) Props() *CfnCrawlerMixinProps {
	var returns *CfnCrawlerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawlerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Crawler`.
func NewCfnCrawlerPropsMixin(props *CfnCrawlerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCrawlerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCrawlerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCrawlerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Crawler`.
func NewCfnCrawlerPropsMixin_Override(c CfnCrawlerPropsMixin, props *CfnCrawlerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCrawlerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCrawlerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCrawlerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnCrawlerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCrawlerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCrawlerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

