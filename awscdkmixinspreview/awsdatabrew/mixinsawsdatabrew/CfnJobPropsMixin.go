package mixinsawsdatabrew

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdatabrew/mixinsawsdatabrew/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new DataBrew job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnJobPropsMixin := awscdkmixinspreview.Mixins.NewCfnJobPropsMixin(&CfnJobMixinProps{
//   	DatabaseOutputs: []interface{}{
//   		&DatabaseOutputProperty{
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			DatabaseOutputMode: jsii.String("databaseOutputMode"),
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   		},
//   	},
//   	DataCatalogOutputs: []interface{}{
//   		&DataCatalogOutputProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			Overwrite: jsii.Boolean(false),
//   			S3Options: &S3TableOutputOptionsProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//   	DatasetName: jsii.String("datasetName"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EncryptionMode: jsii.String("encryptionMode"),
//   	JobSample: &JobSampleProperty{
//   		Mode: jsii.String("mode"),
//   		Size: jsii.Number(123),
//   	},
//   	LogSubscription: jsii.String("logSubscription"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	OutputLocation: &OutputLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   	Outputs: []interface{}{
//   		&OutputProperty{
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			Format: jsii.String("format"),
//   			FormatOptions: &OutputFormatOptionsProperty{
//   				Csv: &CsvOutputOptionsProperty{
//   					Delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   			MaxOutputFiles: jsii.Number(123),
//   			Overwrite: jsii.Boolean(false),
//   			PartitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	ProfileConfiguration: &ProfileConfigurationProperty{
//   		ColumnStatisticsConfigurations: []interface{}{
//   			&ColumnStatisticsConfigurationProperty{
//   				Selectors: []interface{}{
//   					&ColumnSelectorProperty{
//   						Name: jsii.String("name"),
//   						Regex: jsii.String("regex"),
//   					},
//   				},
//   				Statistics: &StatisticsConfigurationProperty{
//   					IncludedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					Overrides: []interface{}{
//   						&StatisticOverrideProperty{
//   							Parameters: map[string]*string{
//   								"parametersKey": jsii.String("parameters"),
//   							},
//   							Statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		DatasetStatisticsConfiguration: &StatisticsConfigurationProperty{
//   			IncludedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			Overrides: []interface{}{
//   				&StatisticOverrideProperty{
//   					Parameters: map[string]*string{
//   						"parametersKey": jsii.String("parameters"),
//   					},
//   					Statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		EntityDetectorConfiguration: &EntityDetectorConfigurationProperty{
//   			AllowedStatistics: &AllowedStatisticsProperty{
//   				Statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   			EntityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//   		},
//   		ProfileColumns: []interface{}{
//   			&ColumnSelectorProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	ProjectName: jsii.String("projectName"),
//   	Recipe: &RecipeProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	Type: jsii.String("type"),
//   	ValidationConfigurations: []interface{}{
//   		&ValidationConfigurationProperty{
//   			RulesetArn: jsii.String("rulesetArn"),
//   			ValidationMode: jsii.String("validationMode"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html
//
type CfnJobPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnJobMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobPropsMixin
type jsiiProxy_CfnJobPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnJobPropsMixin) Props() *CfnJobMixinProps {
	var returns *CfnJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataBrew::Job`.
func NewCfnJobPropsMixin(props *CfnJobMixinProps, options *mixins.CfnPropertyMixinOptions) CfnJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataBrew::Job`.
func NewCfnJobPropsMixin_Override(c CfnJobPropsMixin, props *CfnJobMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

