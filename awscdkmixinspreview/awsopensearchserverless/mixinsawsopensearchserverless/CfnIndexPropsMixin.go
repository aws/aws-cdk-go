package mixinsawsopensearchserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsopensearchserverless/mixinsawsopensearchserverless/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// An OpenSearch Serverless index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var propertyMappingProperty_ PropertyMappingProperty
//
//   cfnIndexPropsMixin := awscdkmixinspreview.Mixins.NewCfnIndexPropsMixin(&CfnIndexMixinProps{
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	IndexName: jsii.String("indexName"),
//   	Mappings: &MappingsProperty{
//   		Properties: map[string]interface{}{
//   			"propertiesKey": &PropertyMappingProperty{
//   				"dimension": jsii.Number(123),
//   				"index": jsii.Boolean(false),
//   				"method": &MethodProperty{
//   					"engine": jsii.String("engine"),
//   					"name": jsii.String("name"),
//   					"parameters": &ParametersProperty{
//   						"efConstruction": jsii.Number(123),
//   						"m": jsii.Number(123),
//   					},
//   					"spaceType": jsii.String("spaceType"),
//   				},
//   				"properties": map[string]interface{}{
//   					"propertiesKey": propertyMappingProperty_,
//   				},
//   				"type": jsii.String("type"),
//   				"value": jsii.String("value"),
//   			},
//   		},
//   	},
//   	Settings: &IndexSettingsProperty{
//   		Index: &IndexProperty{
//   			Knn: jsii.Boolean(false),
//   			KnnAlgoParamEfSearch: jsii.Number(123),
//   			RefreshInterval: jsii.String("refreshInterval"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-index.html
//
type CfnIndexPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIndexMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIndexPropsMixin
type jsiiProxy_CfnIndexPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIndexPropsMixin) Props() *CfnIndexMixinProps {
	var returns *CfnIndexMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndexPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpenSearchServerless::Index`.
func NewCfnIndexPropsMixin(props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIndexPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIndexPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIndexPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpenSearchServerless::Index`.
func NewCfnIndexPropsMixin_Override(c CfnIndexPropsMixin, props *CfnIndexMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnIndexPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIndexPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIndexPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnIndexPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndexPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opensearchserverless.mixins.CfnIndexPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIndexPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIndexPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

