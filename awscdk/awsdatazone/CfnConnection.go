package awsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone"
	"github.com/aws/constructs-go/constructs/v10"
)

// In Amazon DataZone, a connection enables you to connect your resources (domains, projects, and environments) to external resources and services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnection := awscdk.Aws_datazone.NewCfnConnection(this, jsii.String("MyCfnConnection"), &CfnConnectionProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AwsLocation: &AwsLocationProperty{
//   		AccessRole: jsii.String("accessRole"),
//   		AwsAccountId: jsii.String("awsAccountId"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		IamConnectionId: jsii.String("iamConnectionId"),
//   	},
//   	Description: jsii.String("description"),
//   	EnableTrustedIdentityPropagation: jsii.Boolean(false),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	Props: &ConnectionPropertiesInputProperty{
//   		AmazonQProperties: &AmazonQPropertiesInputProperty{
//   			AuthMode: jsii.String("authMode"),
//   			IsEnabled: jsii.Boolean(false),
//   			ProfileArn: jsii.String("profileArn"),
//   		},
//   		AthenaProperties: &AthenaPropertiesInputProperty{
//   			WorkgroupName: jsii.String("workgroupName"),
//   		},
//   		GlueProperties: &GluePropertiesInputProperty{
//   			GlueConnectionInput: &GlueConnectionInputProperty{
//   				AthenaProperties: map[string]*string{
//   					"athenaPropertiesKey": jsii.String("athenaProperties"),
//   				},
//   				AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   					AuthenticationType: jsii.String("authenticationType"),
//   					BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   						Password: jsii.String("password"),
//   						UserName: jsii.String("userName"),
//   					},
//   					CustomAuthenticationCredentials: map[string]*string{
//   						"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   					},
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					OAuth2Properties: &OAuth2PropertiesProperty{
//   						AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   							AuthorizationCode: jsii.String("authorizationCode"),
//   							RedirectUri: jsii.String("redirectUri"),
//   						},
//   						OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   							AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   							UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   						},
//   						OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   							AccessToken: jsii.String("accessToken"),
//   							JwtToken: jsii.String("jwtToken"),
//   							RefreshToken: jsii.String("refreshToken"),
//   							UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   						},
//   						OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   						TokenUrl: jsii.String("tokenUrl"),
//   						TokenUrlParametersMap: map[string]*string{
//   							"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   						},
//   					},
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				ConnectionProperties: map[string]*string{
//   					"connectionPropertiesKey": jsii.String("connectionProperties"),
//   				},
//   				ConnectionType: jsii.String("connectionType"),
//   				Description: jsii.String("description"),
//   				MatchCriteria: jsii.String("matchCriteria"),
//   				Name: jsii.String("name"),
//   				PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					SecurityGroupIdList: []*string{
//   						jsii.String("securityGroupIdList"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   					SubnetIdList: []*string{
//   						jsii.String("subnetIdList"),
//   					},
//   				},
//   				PythonProperties: map[string]*string{
//   					"pythonPropertiesKey": jsii.String("pythonProperties"),
//   				},
//   				SparkProperties: map[string]*string{
//   					"sparkPropertiesKey": jsii.String("sparkProperties"),
//   				},
//   				ValidateCredentials: jsii.Boolean(false),
//   				ValidateForComputeEnvironments: []*string{
//   					jsii.String("validateForComputeEnvironments"),
//   				},
//   			},
//   		},
//   		HyperPodProperties: &HyperPodPropertiesInputProperty{
//   			ClusterName: jsii.String("clusterName"),
//   		},
//   		IamProperties: &IamPropertiesInputProperty{
//   			GlueLineageSyncEnabled: jsii.Boolean(false),
//   		},
//   		MlflowProperties: &MlflowPropertiesInputProperty{
//   			TrackingServerArn: jsii.String("trackingServerArn"),
//   		},
//   		RedshiftProperties: &RedshiftPropertiesInputProperty{
//   			Credentials: &RedshiftCredentialsProperty{
//   				SecretArn: jsii.String("secretArn"),
//   				UsernamePassword: &UsernamePasswordProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Host: jsii.String("host"),
//   			LineageSync: &RedshiftLineageSyncConfigurationInputProperty{
//   				Enabled: jsii.Boolean(false),
//   				Schedule: &LineageSyncScheduleProperty{
//   					Schedule: jsii.String("schedule"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			Storage: &RedshiftStoragePropertiesProperty{
//   				ClusterName: jsii.String("clusterName"),
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   		},
//   		S3Properties: &S3PropertiesInputProperty{
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			S3AccessGrantLocationId: jsii.String("s3AccessGrantLocationId"),
//   		},
//   		SparkEmrProperties: &SparkEmrPropertiesInputProperty{
//   			ComputeArn: jsii.String("computeArn"),
//   			InstanceProfileArn: jsii.String("instanceProfileArn"),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			LogUri: jsii.String("logUri"),
//   			ManagedEndpointArn: jsii.String("managedEndpointArn"),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			RuntimeRole: jsii.String("runtimeRole"),
//   			TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   		},
//   		SparkGlueProperties: &SparkGluePropertiesInputProperty{
//   			AdditionalArgs: &SparkGlueArgsProperty{
//   				Connection: jsii.String("connection"),
//   			},
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   			GlueVersion: jsii.String("glueVersion"),
//   			IdleTimeout: jsii.Number(123),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			NumberOfWorkers: jsii.Number(123),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			WorkerType: jsii.String("workerType"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html
//
type CfnConnection interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsdatazone.IConnectionRef
	// The ID of the connection.
	AttrConnectionId() *string
	// The domain ID of the connection.
	AttrDomainId() *string
	// The domain unit ID of the connection.
	AttrDomainUnitId() *string
	// The ID of the environment.
	AttrEnvironmentId() *string
	// The environment user role.
	AttrEnvironmentUserRole() *string
	// The ID of the project.
	AttrProjectId() *string
	// The type of the connection.
	AttrType() *string
	// The location where the connection is created.
	AwsLocation() interface{}
	SetAwsLocation(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a Connection resource.
	ConnectionRef() *interfacesawsdatazone.ConnectionReference
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Connection description.
	Description() *string
	SetDescription(val *string)
	// The ID of the domain where the connection is created.
	DomainIdentifier() *string
	SetDomainIdentifier(val *string)
	// Specifies whether the trusted identity propagation is enabled.
	EnableTrustedIdentityPropagation() interface{}
	SetEnableTrustedIdentityPropagation(val interface{})
	Env() *interfaces.ResourceEnvironment
	// The ID of the environment where the connection is created.
	EnvironmentIdentifier() *string
	SetEnvironmentIdentifier(val *string)
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
	// The name of the connection.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The identifier of the project in which the connection should be created.
	ProjectIdentifier() *string
	SetProjectIdentifier(val *string)
	// Connection props.
	Props() interface{}
	SetProps(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The scope of the connection.
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnConnection
type jsiiProxy_CfnConnection struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsdatazoneIConnectionRef
}

func (j *jsiiProxy_CfnConnection) AttrConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrDomainUnitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainUnitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrEnvironmentUserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentUserRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AwsLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) ConnectionRef() *interfacesawsdatazone.ConnectionReference {
	var returns *interfacesawsdatazone.ConnectionReference
	_jsii_.Get(
		j,
		"connectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) DomainIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) EnableTrustedIdentityPropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTrustedIdentityPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) EnvironmentIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) ProjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Props() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataZone::Connection`.
func NewCfnConnection(scope constructs.Construct, id *string, props *CfnConnectionProps) CfnConnection {
	_init_.Initialize()

	if err := validateNewCfnConnectionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnection{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataZone::Connection`.
func NewCfnConnection_Override(c CfnConnection, scope constructs.Construct, id *string, props *CfnConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnection)SetAwsLocation(val interface{}) {
	if err := j.validateSetAwsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsLocation",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetDomainIdentifier(val *string) {
	if err := j.validateSetDomainIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetEnableTrustedIdentityPropagation(val interface{}) {
	if err := j.validateSetEnableTrustedIdentityPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTrustedIdentityPropagation",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetEnvironmentIdentifier(val *string) {
	_jsii_.Set(
		j,
		"environmentIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetProjectIdentifier(val *string) {
	_jsii_.Set(
		j,
		"projectIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetProps(val interface{}) {
	if err := j.validateSetPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"props",
		val,
	)
}

func (j *jsiiProxy_CfnConnection)SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

// Checks whether the given object is a CfnConnection.
func CfnConnection_IsCfnConnection(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnection_IsCfnConnectionParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		"isCfnConnection",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnection_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnection_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnConnection_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnection_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		"isCfnResource",
		[]interface{}{x},
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
func CfnConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnection_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datazone.CfnConnection",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnection) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnection) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnection) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnection) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnection) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnection) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnection) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnection) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnConnection) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnection) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnection) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnection) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnConnection) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnConnection) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

