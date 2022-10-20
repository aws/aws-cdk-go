package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lambda::Function`.
//
// The `AWS::Lambda::Function` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html) . The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
//
// You set the package type to `Image` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) . For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties.
//
// You set the package type to `Zip` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip) . For a .zip file archive, the code property specifies the location of the .zip file. You must also specify the handler and runtime properties. For a Python example, see [Deploy Python Lambda functions with .zip file archives](https://docs.aws.amazon.com/lambda/latest/dg/python-package.html) .
//
// You can use [code signing](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html) if your deployment package is a .zip file archive. To enable code signing for this function, specify the ARN of a code-signing configuration. When a user attempts to deploy a code package with `UpdateFunctionCode` , Lambda checks that the code package has a valid signature from a trusted publisher. The code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
//
// Note that you configure [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) on a `AWS::Lambda::Version` or a `AWS::Lambda::Alias` .
//
// For a complete introduction to Lambda functions, see [What is Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/lambda-welcome.html) in the *Lambda developer guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunction := awscdk.Aws_lambda.NewCfnFunction(this, jsii.String("MyCfnFunction"), &cfnFunctionProps{
//   	code: &codeProperty{
//   		imageUri: jsii.String("imageUri"),
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		zipFile: jsii.String("zipFile"),
//   	},
//   	role: jsii.String("role"),
//
//   	// the properties below are optional
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	deadLetterConfig: &deadLetterConfigProperty{
//   		targetArn: jsii.String("targetArn"),
//   	},
//   	description: jsii.String("description"),
//   	environment: &environmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	ephemeralStorage: &ephemeralStorageProperty{
//   		size: jsii.Number(123),
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	runtime: jsii.String("runtime"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeout: jsii.Number(123),
//   	tracingConfig: &tracingConfigProperty{
//   		mode: jsii.String("mode"),
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The instruction set architecture that the function supports.
	//
	// Enter a string array with one of the valid values (arm64 or x86_64). The default value is `x86_64` .
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	// The Amazon Resource Name (ARN) of the function.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The code for the function.
	Code() interface{}
	SetCode(val interface{})
	// To enable code signing for this function, specify the ARN of a code-signing configuration.
	//
	// A code-signing configuration
	// includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
	//
	// For more information, see [Dead Letter Queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) .
	DeadLetterConfig() interface{}
	SetDeadLetterConfig(val interface{})
	// A description of the function.
	Description() *string
	SetDescription(val *string)
	// Environment variables that are accessible from function code during execution.
	Environment() interface{}
	SetEnvironment(val interface{})
	// The size of the function’s /tmp directory in MB.
	//
	// The default value is 512, but can be any whole number between 512 and 10240 MB.
	EphemeralStorage() interface{}
	SetEphemeralStorage(val interface{})
	// Connection settings for an Amazon EFS file system.
	//
	// To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a `DependsOn` attribute to ensure that the mount target is created or updated before the function.
	//
	// For more information about using the `DependsOn` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	FileSystemConfigs() interface{}
	SetFileSystemConfigs(val interface{})
	// The name of the Lambda function, up to 64 characters in length.
	//
	// If you don't specify a name, AWS CloudFormation generates one.
	//
	// If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	FunctionName() *string
	SetFunctionName(val *string)
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Programming Model](https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html) .
	Handler() *string
	SetHandler(val *string)
	// Configuration values that override the container image Dockerfile settings.
	//
	// See [Container settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
	ImageConfig() interface{}
	SetImageConfig(val interface{})
	// The ARN of the AWS Key Management Service ( AWS KMS ) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	// A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.
	Layers() *[]*string
	SetLayers(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-memory.html) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize() *float64
	SetMemorySize(val *float64)
	// The tree node.
	Node() constructs.Node
	// The type of deployment package.
	//
	// Set to `Image` for container image and set `Zip` for .zip file archive.
	PackageType() *string
	SetPackageType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The number of simultaneous executions to reserve for the function.
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	// The Amazon Resource Name (ARN) of the function's execution role.
	Role() *string
	SetRole(val *string)
	// The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) . Runtime is required if the deployment package is a .zip file archive.
	Runtime() *string
	SetRuntime(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.
	Tags() awscdk.TagManager
	// The amount of time (in seconds) that Lambda allows a function to run before stopping it.
	//
	// The default is 3 seconds. The maximum allowed value is 900 seconds. For additional information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html) .
	Timeout() *float64
	SetTimeout(val *float64)
	// Set `Mode` to `Active` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) .
	TracingConfig() interface{}
	SetTracingConfig(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// For network connectivity to AWS resources in a [VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-network.html) , specify a list of security groups and subnets in the VPC.
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFunction
type jsiiProxy_CfnFunction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeadLetterConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) EphemeralStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FileSystemConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ImageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) TracingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::Function`.
func NewCfnFunction(scope constructs.Construct, id *string, props *CfnFunctionProps) CfnFunction {
	_init_.Initialize()

	if err := validateNewCfnFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::Function`.
func NewCfnFunction_Override(c CfnFunction, scope constructs.Construct, id *string, props *CfnFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunction)SetArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetCode(val interface{}) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetCodeSigningConfigArn(val *string) {
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetDeadLetterConfig(val interface{}) {
	if err := j.validateSetDeadLetterConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadLetterConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetEnvironment(val interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetEphemeralStorage(val interface{}) {
	if err := j.validateSetEphemeralStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralStorage",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetFileSystemConfigs(val interface{}) {
	if err := j.validateSetFileSystemConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetHandler(val *string) {
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetImageConfig(val interface{}) {
	if err := j.validateSetImageConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetLayers(val *[]*string) {
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetMemorySize(val *float64) {
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetPackageType(val *string) {
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetReservedConcurrentExecutions(val *float64) {
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetTracingConfig(val interface{}) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction)SetVpcConfig(val interface{}) {
	if err := j.validateSetVpcConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFunction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunction_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFunction_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunction) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunction) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunction) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunction) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunction) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

