// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppRegistry construct to automatically create an application with the given name and description.
//
// The application name must be unique at the account level per region and it's immutable.
// This construct will automatically associate all stacks in the given scope, however
// in case of a `Pipeline` stack, stage underneath the pipeline will not automatically be associated and
// needs to be associated separately.
//
// If cross account stack is detected, then this construct will automatically share the application to consumer accounts.
// Cross account feature will only work for non environment agnostic stacks.
//
// Example:
//   app := awscdk.NewApp()
//   associatedApp := appreg.NewApplicationAssociator(app, jsii.String("AssociatedApplication"), &applicationAssociatorProps{
//   	applications: []targetApplication{
//   		appreg.*targetApplication.createApplicationStack(&createTargetApplicationOptions{
//   			applicationName: jsii.String("MyAssociatedApplication"),
//   			// 'Application containing stacks deployed via CDK.' is the default
//   			applicationDescription: jsii.String("Associated Application description"),
//   			stackName: jsii.String("MyAssociatedApplicationStack"),
//   			// AWS Account and Region that are implied by the current CLI configuration is the default
//   			env: &environment{
//   				account: jsii.String("123456789012"),
//   				region: jsii.String("us-east-1"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ApplicationAssociator interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Get the AppRegistry application.
	// Experimental.
	AppRegistryApplication() IApplication
	// Associate this application with the given stage.
	// Experimental.
	AssociateStage(stage awscdk.Stage) awscdk.Stage
	// Validates if a stage is already associated to the application.
	// Experimental.
	IsStageAssociated(stage awscdk.Stage) *bool
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationAssociator
type jsiiProxy_ApplicationAssociator struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationAssociator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplicationAssociator(scope awscdk.App, id *string, props *ApplicationAssociatorProps) ApplicationAssociator {
	_init_.Initialize()

	if err := validateNewApplicationAssociatorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationAssociator{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationAssociator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplicationAssociator_Override(a ApplicationAssociator, scope awscdk.App, id *string, props *ApplicationAssociatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationAssociator",
		[]interface{}{scope, id, props},
		a,
	)
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
func ApplicationAssociator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationAssociator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.ApplicationAssociator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAssociator) AppRegistryApplication() IApplication {
	var returns IApplication

	_jsii_.Invoke(
		a,
		"appRegistryApplication",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAssociator) AssociateStage(stage awscdk.Stage) awscdk.Stage {
	if err := a.validateAssociateStageParameters(stage); err != nil {
		panic(err)
	}
	var returns awscdk.Stage

	_jsii_.Invoke(
		a,
		"associateStage",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAssociator) IsStageAssociated(stage awscdk.Stage) *bool {
	if err := a.validateIsStageAssociatedParameters(stage); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"isStageAssociated",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAssociator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

