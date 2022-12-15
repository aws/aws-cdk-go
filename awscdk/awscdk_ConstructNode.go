// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents the construct node in the scope tree.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//
//   constructNode := monocdk.NewConstructNode(construct, construct, jsii.String("id"))
//
// Experimental.
type ConstructNode interface {
	// Returns an opaque tree-unique address for this construct.
	//
	// Addresses are 42 characters hexadecimal strings. They begin with "c8"
	// followed by 40 lowercase hexadecimal characters (0-9a-f).
	//
	// Addresses are calculated using a SHA-1 of the components of the construct
	// path.
	//
	// To enable refactorings of construct trees, constructs with the ID `Default`
	// will be excluded from the calculation. In those cases constructs in the
	// same tree may have the same addreess.
	//
	// Example value: `c83a2846e506bcc5f10682b564084bca2d275709ee`.
	// Experimental.
	Addr() *string
	// All direct children of this construct.
	// Experimental.
	Children() *[]IConstruct
	// Returns the child construct that has the id `Default` or `Resource"`.
	//
	// This is usually the construct that provides the bulk of the underlying functionality.
	// Useful for modifications of the underlying construct that are not available at the higher levels.
	// Override the defaultChild property.
	//
	// This should only be used in the cases where the correct
	// default child is not named 'Resource' or 'Default' as it
	// should be.
	//
	// If you set this to undefined, the default behavior of finding
	// the child named 'Resource' or 'Default' will be used.
	//
	// Returns: a construct or undefined if there is no default child.
	// Experimental.
	DefaultChild() IConstruct
	// Experimental.
	SetDefaultChild(val IConstruct)
	// Return all dependencies registered on this node or any of its children.
	// Experimental.
	Dependencies() *[]*Dependency
	// The id of this construct within the current scope.
	//
	// This is a a scope-unique id. To obtain an app-unique id for this construct, use `uniqueId`.
	// Experimental.
	Id() *string
	// Returns true if this construct or the scopes in which it is defined are locked.
	// Experimental.
	Locked() *bool
	// DEPRECATED.
	// Deprecated: use `metadataEntry`.
	Metadata() *[]*cxapi.MetadataEntry
	// An immutable array of metadata objects associated with this construct.
	//
	// This can be used, for example, to implement support for deprecation notices, source mapping, etc.
	// Experimental.
	MetadataEntry() *[]*constructs.MetadataEntry
	// The full, absolute path of this construct in the tree.
	//
	// Components are separated by '/'.
	// Experimental.
	Path() *string
	// Returns: The root of the construct tree.
	// Experimental.
	Root() IConstruct
	// Returns the scope in which this construct is defined.
	//
	// The value is `undefined` at the root of the construct scope tree.
	// Experimental.
	Scope() IConstruct
	// All parent scopes of this construct.
	//
	// Returns: a list of parent scopes. The last element in the list will always
	// be the current construct and the first element will be the root of the
	// tree.
	// Experimental.
	Scopes() *[]IConstruct
	// A tree-global unique alphanumeric identifier for this construct.
	//
	// Includes
	// all components of the tree.
	// Deprecated: use `node.addr` to obtain a consistent 42 character address for
	// this node (see https://github.com/aws/constructs/pull/314).
	// Alternatively, to get a CloudFormation-compatible unique identifier, use
	// `Names.uniqueId()`.
	UniqueId() *string
	// Add an ordering dependency on another Construct.
	//
	// All constructs in the dependency's scope will be deployed before any
	// construct in this construct's scope.
	// Experimental.
	AddDependency(dependencies ...IDependable)
	// DEPRECATED: Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail synthesis when errors are reported.
	// Deprecated: use `Annotations.of(construct).addError()`
	AddError(message *string)
	// DEPRECATED: Adds a { "info": <message> } metadata entry to this construct.
	//
	// The toolkit will display the info message when apps are synthesized.
	// Deprecated: use `Annotations.of(construct).addInfo()`
	AddInfo(message *string)
	// Adds a metadata entry to this construct.
	//
	// Entries are arbitrary values and will also include a stack trace to allow tracing back to
	// the code location for when the entry was added. It can be used, for example, to include source
	// mapping in CloudFormation templates to improve diagnostics.
	// Experimental.
	AddMetadata(type_ *string, data interface{}, fromFunction interface{})
	// Add a validator to this construct Node.
	// Experimental.
	AddValidation(validation constructs.IValidation)
	// DEPRECATED: Adds a { "warning": <message> } metadata entry to this construct.
	//
	// The toolkit will display the warning when an app is synthesized, or fail
	// if run in --strict mode.
	// Deprecated: use `Annotations.of(construct).addWarning()`
	AddWarning(message *string)
	// DEPRECATED: Applies the aspect to this Constructs node.
	// Deprecated: This API is going to be removed in the next major version of
	// the AWS CDK. Please use `Aspects.of(scope).add()` instead.
	ApplyAspect(aspect IAspect)
	// Return this construct and all of its children in the given order.
	// Experimental.
	FindAll(order ConstructOrder) *[]IConstruct
	// Return a direct child by id.
	//
	// Throws an error if the child is not found.
	//
	// Returns: Child with the given id.
	// Experimental.
	FindChild(id *string) IConstruct
	// This can be used to set contextual values.
	//
	// Context must be set before any children are added, since children may consult context info during construction.
	// If the key already exists, it will be overridden.
	// Experimental.
	SetContext(key *string, value interface{})
	// Return a direct child by id, or undefined.
	//
	// Returns: the child if found, or undefined.
	// Experimental.
	TryFindChild(id *string) IConstruct
	// Retrieves a value from tree context.
	//
	// Context is usually initialized at the root, but can be overridden at any point in the tree.
	//
	// Returns: The context value or `undefined` if there is no context value for the key.
	// Experimental.
	TryGetContext(key *string) interface{}
	// Remove the child with the given name, if present.
	//
	// Returns: Whether a child with the given name was deleted.
	// Experimental.
	TryRemoveChild(childName *string) *bool
}

// The jsii proxy struct for ConstructNode
type jsiiProxy_ConstructNode struct {
	_ byte // padding
}

func (j *jsiiProxy_ConstructNode) Addr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Children() *[]IConstruct {
	var returns *[]IConstruct
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) DefaultChild() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"defaultChild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Dependencies() *[]*Dependency {
	var returns *[]*Dependency
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Locked() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Metadata() *[]*cxapi.MetadataEntry {
	var returns *[]*cxapi.MetadataEntry
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) MetadataEntry() *[]*constructs.MetadataEntry {
	var returns *[]*constructs.MetadataEntry
	_jsii_.Get(
		j,
		"metadataEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Root() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Scope() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) Scopes() *[]IConstruct {
	var returns *[]IConstruct
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructNode) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}


// Experimental.
func NewConstructNode(host Construct, scope IConstruct, id *string) ConstructNode {
	_init_.Initialize()

	if err := validateNewConstructNodeParameters(host, scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConstructNode{}

	_jsii_.Create(
		"monocdk.ConstructNode",
		[]interface{}{host, scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewConstructNode_Override(c ConstructNode, host Construct, scope IConstruct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.ConstructNode",
		[]interface{}{host, scope, id},
		c,
	)
}

func (j *jsiiProxy_ConstructNode)SetDefaultChild(val IConstruct) {
	_jsii_.Set(
		j,
		"defaultChild",
		val,
	)
}

// Invokes "prepare" on all constructs (depth-first, post-order) in the tree under `node`.
// Deprecated: Use `app.synth()` instead
func ConstructNode_Prepare(node ConstructNode) {
	_init_.Initialize()

	if err := validateConstructNode_PrepareParameters(node); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.ConstructNode",
		"prepare",
		[]interface{}{node},
	)
}

// Synthesizes a CloudAssembly from a construct tree.
// Deprecated: Use `app.synth()` or `stage.synth()` instead
func ConstructNode_Synth(node ConstructNode, options *SynthesisOptions) cxapi.CloudAssembly {
	_init_.Initialize()

	if err := validateConstructNode_SynthParameters(node, options); err != nil {
		panic(err)
	}
	var returns cxapi.CloudAssembly

	_jsii_.StaticInvoke(
		"monocdk.ConstructNode",
		"synth",
		[]interface{}{node, options},
		&returns,
	)

	return returns
}

// Invokes "validate" on all constructs in the tree (depth-first, pre-order) and returns the list of all errors.
//
// An empty list indicates that there are no errors.
// Experimental.
func ConstructNode_Validate(node ConstructNode) *[]*ValidationError {
	_init_.Initialize()

	if err := validateConstructNode_ValidateParameters(node); err != nil {
		panic(err)
	}
	var returns *[]*ValidationError

	_jsii_.StaticInvoke(
		"monocdk.ConstructNode",
		"validate",
		[]interface{}{node},
		&returns,
	)

	return returns
}

func ConstructNode_PATH_SEP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.ConstructNode",
		"PATH_SEP",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConstructNode) AddDependency(dependencies ...IDependable) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

func (c *jsiiProxy_ConstructNode) AddError(message *string) {
	if err := c.validateAddErrorParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addError",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_ConstructNode) AddInfo(message *string) {
	if err := c.validateAddInfoParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addInfo",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_ConstructNode) AddMetadata(type_ *string, data interface{}, fromFunction interface{}) {
	if err := c.validateAddMetadataParameters(type_, data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{type_, data, fromFunction},
	)
}

func (c *jsiiProxy_ConstructNode) AddValidation(validation constructs.IValidation) {
	if err := c.validateAddValidationParameters(validation); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addValidation",
		[]interface{}{validation},
	)
}

func (c *jsiiProxy_ConstructNode) AddWarning(message *string) {
	if err := c.validateAddWarningParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addWarning",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_ConstructNode) ApplyAspect(aspect IAspect) {
	if err := c.validateApplyAspectParameters(aspect); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyAspect",
		[]interface{}{aspect},
	)
}

func (c *jsiiProxy_ConstructNode) FindAll(order ConstructOrder) *[]IConstruct {
	var returns *[]IConstruct

	_jsii_.Invoke(
		c,
		"findAll",
		[]interface{}{order},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructNode) FindChild(id *string) IConstruct {
	if err := c.validateFindChildParameters(id); err != nil {
		panic(err)
	}
	var returns IConstruct

	_jsii_.Invoke(
		c,
		"findChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructNode) SetContext(key *string, value interface{}) {
	if err := c.validateSetContextParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"setContext",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_ConstructNode) TryFindChild(id *string) IConstruct {
	if err := c.validateTryFindChildParameters(id); err != nil {
		panic(err)
	}
	var returns IConstruct

	_jsii_.Invoke(
		c,
		"tryFindChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructNode) TryGetContext(key *string) interface{} {
	if err := c.validateTryGetContextParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"tryGetContext",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructNode) TryRemoveChild(childName *string) *bool {
	if err := c.validateTryRemoveChildParameters(childName); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"tryRemoveChild",
		[]interface{}{childName},
		&returns,
	)

	return returns
}

