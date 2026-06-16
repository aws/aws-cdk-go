package awsresiliencehubv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsresiliencehubv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a resilience-managed service with associated systems, input sources, assertions, and service functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnServicePropsMixin := awscdkcfnpropertymixins.Aws_resiliencehubv2.NewCfnServicePropsMixin(&CfnServiceMixinProps{
//   	Assertions: []interface{}{
//   		&AssertionDefinitionProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	AssociatedSystems: []interface{}{
//   		&AssociatedSystemProperty{
//   			SystemArn: jsii.String("systemArn"),
//   			UserJourneyIds: []*string{
//   				jsii.String("userJourneyIds"),
//   			},
//   		},
//   	},
//   	DependencyDiscovery: jsii.String("dependencyDiscovery"),
//   	Description: jsii.String("description"),
//   	InputSources: []interface{}{
//   		&InputSourceDefinitionProperty{
//   			ResourceConfiguration: &ResourceConfigurationProperty{
//   				CfnStackArn: jsii.String("cfnStackArn"),
//   				DesignFileS3Url: jsii.String("designFileS3Url"),
//   				Eks: &EksSourceProperty{
//   					ClusterArn: jsii.String("clusterArn"),
//   					Namespaces: []*string{
//   						jsii.String("namespaces"),
//   					},
//   				},
//   				ResourceTags: []interface{}{
//   					&ResourceTagProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				TfStateFileUrl: jsii.String("tfStateFileUrl"),
//   			},
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	PermissionModel: &PermissionModelProperty{
//   		CrossAccountRoleArns: []interface{}{
//   			&CrossAccountRoleConfigurationProperty{
//   				CrossAccountRoleArn: jsii.String("crossAccountRoleArn"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		InvokerRoleName: jsii.String("invokerRoleName"),
//   	},
//   	PolicyArn: jsii.String("policyArn"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	ReportConfiguration: &ServiceReportConfigurationProperty{
//   		ReportOutput: []interface{}{
//   			&ReportOutputConfigurationProperty{
//   				S3: &S3ReportOutputConfigurationProperty{
//   					BucketOwner: jsii.String("bucketOwner"),
//   					BucketPath: jsii.String("bucketPath"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html
//
type CfnServicePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServiceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServicePropsMixin
type jsiiProxy_CfnServicePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServicePropsMixin) Props() *CfnServiceMixinProps {
	var returns *CfnServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServicePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResilienceHubV2::Service`.
func NewCfnServicePropsMixin(props *CfnServiceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResilienceHubV2::Service`.
func NewCfnServicePropsMixin_Override(c CfnServicePropsMixin, props *CfnServiceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

