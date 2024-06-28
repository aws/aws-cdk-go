package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::LakeFormation::PrincipalPermissions` resource represents the permissions that a principal has on a Data Catalog resource (such as AWS Glue databases or AWS Glue tables).
//
// When you create a `PrincipalPermissions` resource, the permissions are granted via the AWS Lake Formation `GrantPermissions` API operation. When you delete a `PrincipalPermissions` resource, the permissions on principal-resource pair are revoked via the AWS Lake Formation `RevokePermissions` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnPrincipalPermissions := awscdk.Aws_lakeformation.NewCfnPrincipalPermissions(this, jsii.String("MyCfnPrincipalPermissions"), &CfnPrincipalPermissionsProps{
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	PermissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	Principal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	Resource: &ResourceProperty{
//   		Catalog: catalog,
//   		Database: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		DataCellsFilter: &DataCellsFilterResourceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			TableCatalogId: jsii.String("tableCatalogId"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		DataLocation: &DataLocationResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		LfTag: &LFTagKeyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		LfTagPolicy: &LFTagPolicyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Expression: []interface{}{
//   				&LFTagProperty{
//   					TagKey: jsii.String("tagKey"),
//   					TagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   		Table: &TableResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   			TableWildcard: tableWildcard,
//   		},
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			ColumnWildcard: &ColumnWildcardProperty{
//   				ExcludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Catalog: jsii.String("catalog"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html
//
type CfnPrincipalPermissions interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Json encoding of the input principal.
	//
	// For example: `{"DataLakePrincipalIdentifier":"arn:aws:iam::123456789012:role/ExampleRole"}`.
	AttrPrincipalIdentifier() *string
	// Json encoding of the input resource.
	//
	// For example: `{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":null,"DataLocation":null,"DataCellsFilter":{"TableCatalogId":"123456789012","DatabaseName":"ExampleDatabase","TableName":"ExampleTable","Name":"ExampleFilter"},"LFTag":null,"LFTagPolicy":null}`.
	AttrResourceIdentifier() *string
	// The identifier for the Data Catalog .
	Catalog() *string
	SetCatalog(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The tree node.
	Node() constructs.Node
	// The permissions granted or revoked.
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	// The principal to be granted a permission.
	Principal() interface{}
	SetPrincipal(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The resource to be granted or revoked permissions.
	Resource() interface{}
	SetResource(val interface{})
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

// The jsii proxy struct for CfnPrincipalPermissions
type jsiiProxy_CfnPrincipalPermissions struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPrincipalPermissions) AttrPrincipalIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrincipalIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) AttrResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Principal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Resource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnPrincipalPermissions(scope constructs.Construct, id *string, props *CfnPrincipalPermissionsProps) CfnPrincipalPermissions {
	_init_.Initialize()

	if err := validateNewCfnPrincipalPermissionsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPrincipalPermissions{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnPrincipalPermissions_Override(c CfnPrincipalPermissions, scope constructs.Construct, id *string, props *CfnPrincipalPermissionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions)SetCatalog(val *string) {
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions)SetPermissions(val *[]*string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions)SetPermissionsWithGrantOption(val *[]*string) {
	if err := j.validateSetPermissionsWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions)SetPrincipal(val interface{}) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions)SetResource(val interface{}) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPrincipalPermissions_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrincipalPermissions_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnPrincipalPermissions_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrincipalPermissions_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
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
func CfnPrincipalPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrincipalPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrincipalPermissions_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lakeformation.CfnPrincipalPermissions",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnPrincipalPermissions) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnPrincipalPermissions) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnPrincipalPermissions) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

