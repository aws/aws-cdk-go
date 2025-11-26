package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a data source connector for an Amazon Q Business application.
//
// `CreateDataSource` is a synchronous operation. The operation returns 200 if the data source was successfully created. Otherwise, an exception is raised.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configuration interface{}
//
//   cfnDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataSourcePropsMixin(&CfnDataSourceMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Configuration: configuration,
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DocumentEnrichmentConfiguration: &DocumentEnrichmentConfigurationProperty{
//   		InlineConfigurations: []interface{}{
//   			&InlineDocumentEnrichmentConfigurationProperty{
//   				Condition: &DocumentAttributeConditionProperty{
//   					Key: jsii.String("key"),
//   					Operator: jsii.String("operator"),
//   					Value: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   				DocumentContentOperator: jsii.String("documentContentOperator"),
//   				Target: &DocumentAttributeTargetProperty{
//   					AttributeValueOperator: jsii.String("attributeValueOperator"),
//   					Key: jsii.String("key"),
//   					Value: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   		},
//   		PostExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   		PreExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   	},
//   	IndexId: jsii.String("indexId"),
//   	MediaExtractionConfiguration: &MediaExtractionConfigurationProperty{
//   		AudioExtractionConfiguration: &AudioExtractionConfigurationProperty{
//   			AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   		},
//   		ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   			ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   		},
//   		VideoExtractionConfiguration: &VideoExtractionConfigurationProperty{
//   			VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SyncSchedule: jsii.String("syncSchedule"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html
//
type CfnDataSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataSourcePropsMixin
type jsiiProxy_CfnDataSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Props() *CfnDataSourceMixinProps {
	var returns *CfnDataSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::DataSource`.
func NewCfnDataSourcePropsMixin(props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::DataSource`.
func NewCfnDataSourcePropsMixin_Override(c CfnDataSourcePropsMixin, props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnDataSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

