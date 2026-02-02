package previewawsdatabrewmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatabrew/previewawsdatabrewmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDatasetPropsMixin := awscdkmixinspreview.Mixins.NewCfnDatasetPropsMixin(&CfnDatasetMixinProps{
//   	Format: jsii.String("format"),
//   	FormatOptions: &FormatOptionsProperty{
//   		Csv: &CsvOptionsProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			HeaderRow: jsii.Boolean(false),
//   		},
//   		Excel: &ExcelOptionsProperty{
//   			HeaderRow: jsii.Boolean(false),
//   			SheetIndexes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SheetNames: []*string{
//   				jsii.String("sheetNames"),
//   			},
//   		},
//   		Json: &JsonOptionsProperty{
//   			MultiLine: jsii.Boolean(false),
//   		},
//   	},
//   	Input: &InputProperty{
//   		DatabaseInputDefinition: &DatabaseInputDefinitionProperty{
//   			DatabaseTableName: jsii.String("databaseTableName"),
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   			QueryString: jsii.String("queryString"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		Metadata: &MetadataProperty{
//   			SourceArn: jsii.String("sourceArn"),
//   		},
//   		S3InputDefinition: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PathOptions: &PathOptionsProperty{
//   		FilesLimit: &FilesLimitProperty{
//   			MaxFiles: jsii.Number(123),
//   			Order: jsii.String("order"),
//   			OrderedBy: jsii.String("orderedBy"),
//   		},
//   		LastModifiedDateCondition: &FilterExpressionProperty{
//   			Expression: jsii.String("expression"),
//   			ValuesMap: []interface{}{
//   				&FilterValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		Parameters: []interface{}{
//   			&PathParameterProperty{
//   				DatasetParameter: &DatasetParameterProperty{
//   					CreateColumn: jsii.Boolean(false),
//   					DatetimeOptions: &DatetimeOptionsProperty{
//   						Format: jsii.String("format"),
//   						LocaleCode: jsii.String("localeCode"),
//   						TimezoneOffset: jsii.String("timezoneOffset"),
//   					},
//   					Filter: &FilterExpressionProperty{
//   						Expression: jsii.String("expression"),
//   						ValuesMap: []interface{}{
//   							&FilterValueProperty{
//   								Value: jsii.String("value"),
//   								ValueReference: jsii.String("valueReference"),
//   							},
//   						},
//   					},
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   				PathParameterName: jsii.String("pathParameterName"),
//   			},
//   		},
//   	},
//   	Source: jsii.String("source"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html
//
type CfnDatasetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDatasetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDatasetPropsMixin
type jsiiProxy_CfnDatasetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Props() *CfnDatasetMixinProps {
	var returns *CfnDatasetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataBrew::Dataset`.
func NewCfnDatasetPropsMixin(props *CfnDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDatasetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDatasetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDatasetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataBrew::Dataset`.
func NewCfnDatasetPropsMixin_Override(c CfnDatasetPropsMixin, props *CfnDatasetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDatasetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDatasetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatasetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatasetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDatasetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

