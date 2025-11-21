package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an EC2 Image Builder Distribution Configuration.
//
// Example:
//   distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("DistributionConfiguration"), &DistributionConfigurationProps{
//   	DistributionConfigurationName: jsii.String("test-distribution-configuration"),
//   	Description: jsii.String("A Distribution Configuration"),
//   	AmiDistributions: []AmiDistribution{
//   		&AmiDistribution{
//   			// Distribute AMI to us-east-2 and publish the AMI ID to an SSM parameter
//   			Region: jsii.String("us-east-2"),
//   			SsmParameters: []SSMParameterConfigurations{
//   				&SSMParameterConfigurations{
//   					Parameter: ssm.StringParameter_FromStringParameterAttributes(this, jsii.String("CrossRegionParameter"), &StringParameterAttributes{
//   						ParameterName: jsii.String("/imagebuilder/ami"),
//   						ForceDynamicReference: jsii.Boolean(true),
//   					}),
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // For AMI-based image builds - add an AMI distribution in the current region
//   distributionConfiguration.AddAmiDistributions(&AmiDistribution{
//   	AmiName: jsii.String("imagebuilder-{{ imagebuilder:buildDate }}"),
//   	AmiDescription: jsii.String("Build AMI"),
//   	AmiKmsKey: kms.Key_FromLookup(this, jsii.String("ComponentKey"), &KeyLookupOptions{
//   		AliasName: jsii.String("alias/distribution-encryption-key"),
//   	}),
//   	// Copy the AMI to different accounts
//   	AmiTargetAccountIds: []*string{
//   		jsii.String("123456789012"),
//   		jsii.String("098765432109"),
//   	},
//   	// Add launch permissions on the AMI
//   	AmiLaunchPermission: &AmiLaunchPermission{
//   		OrganizationArns: []*string{
//   			this.FormatArn(&ArnComponents{
//   				Region: jsii.String(""),
//   				Service: jsii.String("organizations"),
//   				Resource: jsii.String("organization"),
//   				ResourceName: jsii.String("o-1234567abc"),
//   			}),
//   		},
//   		OrganizationalUnitArns: []*string{
//   			this.*FormatArn(&ArnComponents{
//   				Region: jsii.String(""),
//   				Service: jsii.String("organizations"),
//   				Resource: jsii.String("ou"),
//   				ResourceName: jsii.String("o-1234567abc/ou-a123-b4567890"),
//   			}),
//   		},
//   		IsPublicUserGroup: jsii.Boolean(true),
//   		AccountIds: []*string{
//   			jsii.String("234567890123"),
//   		},
//   	},
//   	// Attach tags to the AMI
//   	AmiTags: map[string]*string{
//   		"Environment": jsii.String("production"),
//   		"Version": jsii.String("{{ imagebuilder:buildVersion }}"),
//   	},
//   	// Optional - publish the distributed AMI ID to an SSM parameter
//   	SsmParameters: []SSMParameterConfigurations{
//   		&SSMParameterConfigurations{
//   			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("Parameter"), &StringParameterAttributes{
//   				ParameterName: jsii.String("/imagebuilder/ami"),
//   				ForceDynamicReference: jsii.Boolean(true),
//   			}),
//   		},
//   		&SSMParameterConfigurations{
//   			AmiAccount: jsii.String("098765432109"),
//   			DataType: ssm.ParameterDataType_TEXT,
//   			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("CrossAccountParameter"), &StringParameterAttributes{
//   				ParameterName: jsii.String("imagebuilder-prod-ami"),
//   				ForceDynamicReference: jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   	// Optional - create a new launch template version with the distributed AMI ID
//   	LaunchTemplates: []LaunchTemplateConfiguration{
//   		&LaunchTemplateConfiguration{
//   			LaunchTemplate: ec2.LaunchTemplate_FromLaunchTemplateAttributes(this, jsii.String("LaunchTemplate"), &LaunchTemplateAttributes{
//   				LaunchTemplateId: jsii.String("lt-1234"),
//   			}),
//   			SetDefaultVersion: jsii.Boolean(true),
//   		},
//   		&LaunchTemplateConfiguration{
//   			AccountId: jsii.String("123456789012"),
//   			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("CrossAccountLaunchTemplate"), &LaunchTemplateAttributes{
//   				LaunchTemplateId: jsii.String("lt-5678"),
//   			}),
//   			SetDefaultVersion: jsii.Boolean(true),
//   		},
//   	},
//   	// Optional - enable Fast Launch on an imported launch template
//   	FastLaunchConfigurations: []FastLaunchConfiguration{
//   		&FastLaunchConfiguration{
//   			Enabled: jsii.Boolean(true),
//   			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("FastLaunchLT"), &LaunchTemplateAttributes{
//   				LaunchTemplateName: jsii.String("fast-launch-lt"),
//   			}),
//   			MaxParallelLaunches: jsii.Number(10),
//   			TargetSnapshotCount: jsii.Number(2),
//   		},
//   	},
//   	// Optional - license configurations to apply to the AMI
//   	LicenseConfigurationArns: []*string{
//   		jsii.String("arn:aws:license-manager:us-west-2:123456789012:license-configuration:lic-abcdefghijklmnopqrstuvwxyz"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/manage-distribution-settings.html
//
// Experimental.
type DistributionConfiguration interface {
	awscdk.Resource
	IDistributionConfiguration
	// The ARN of the distribution configuration.
	// Experimental.
	DistributionConfigurationArn() *string
	// The name of the distribution configuration.
	// Experimental.
	DistributionConfigurationName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	// Experimental.
	Env() *interfaces.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds AMI distribution settings to the distribution configuration.
	// Experimental.
	AddAmiDistributions(amiDistributions ...*AmiDistribution)
	// Adds container distribution settings to the distribution configuration.
	// Experimental.
	AddContainerDistributions(containerDistributions ...*ContainerDistribution)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant custom actions to the given grantee for the distribution configuration.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the distribution configuration.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DistributionConfiguration
type jsiiProxy_DistributionConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_IDistributionConfiguration
}

func (j *jsiiProxy_DistributionConfiguration) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributionConfiguration) DistributionConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributionConfiguration) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributionConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributionConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributionConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDistributionConfiguration(scope constructs.Construct, id *string, props *DistributionConfigurationProps) DistributionConfiguration {
	_init_.Initialize()

	if err := validateNewDistributionConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DistributionConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDistributionConfiguration_Override(d DistributionConfiguration, scope constructs.Construct, id *string, props *DistributionConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing distribution configuration given its ARN.
// Experimental.
func DistributionConfiguration_FromDistributionConfigurationArn(scope constructs.Construct, id *string, distributionConfigurationArn *string) IDistributionConfiguration {
	_init_.Initialize()

	if err := validateDistributionConfiguration_FromDistributionConfigurationArnParameters(scope, id, distributionConfigurationArn); err != nil {
		panic(err)
	}
	var returns IDistributionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"fromDistributionConfigurationArn",
		[]interface{}{scope, id, distributionConfigurationArn},
		&returns,
	)

	return returns
}

// Import an existing distribution configuration given its name.
//
// The provided name must be normalized by converting
// all alphabetical characters to lowercase, and replacing all spaces and underscores with hyphens.
// Experimental.
func DistributionConfiguration_FromDistributionConfigurationName(scope constructs.Construct, id *string, distributionConfigurationName *string) IDistributionConfiguration {
	_init_.Initialize()

	if err := validateDistributionConfiguration_FromDistributionConfigurationNameParameters(scope, id, distributionConfigurationName); err != nil {
		panic(err)
	}
	var returns IDistributionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"fromDistributionConfigurationName",
		[]interface{}{scope, id, distributionConfigurationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DistributionConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributionConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a DistributionConfiguration.
// Experimental.
func DistributionConfiguration_IsDistributionConfiguration(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributionConfiguration_IsDistributionConfigurationParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"isDistributionConfiguration",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func DistributionConfiguration_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDistributionConfiguration_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DistributionConfiguration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDistributionConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func DistributionConfiguration_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DistributionConfiguration) AddAmiDistributions(amiDistributions ...*AmiDistribution) {
	if err := d.validateAddAmiDistributionsParameters(&amiDistributions); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range amiDistributions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addAmiDistributions",
		args,
	)
}

func (d *jsiiProxy_DistributionConfiguration) AddContainerDistributions(containerDistributions ...*ContainerDistribution) {
	if err := d.validateAddContainerDistributionsParameters(&containerDistributions); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range containerDistributions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addContainerDistributions",
		args,
	)
}

func (d *jsiiProxy_DistributionConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DistributionConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionConfiguration) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := d.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionConfiguration) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

