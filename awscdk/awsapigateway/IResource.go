package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

type IResource interface {
	awscdk.IResource
	// Adds an OPTIONS method to this resource which responds to Cross-Origin Resource Sharing (CORS) preflight requests.
	//
	// Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional
	// HTTP headers to tell browsers to give a web application running at one
	// origin, access to selected resources from a different origin. A web
	// application executes a cross-origin HTTP request when it requests a
	// resource that has a different origin (domain, protocol, or port) from its
	// own.
	//
	// Returns: a `Method` object.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
	//
	AddCorsPreflight(options *CorsOptions) Method
	// Defines a new method for this resource.
	//
	// Returns: The newly created `Method` object.
	AddMethod(httpMethod *string, target Integration, options *MethodOptions) Method
	// Adds a greedy proxy resource ("{proxy+}") and an ANY method to this route.
	AddProxy(options *ProxyResourceOptions) ProxyResource
	// Defines a new child resource where this resource is the parent.
	//
	// Returns: A Resource object.
	AddResource(pathPart *string, options *ResourceOptions) Resource
	// Retrieves a child resource by path part.
	//
	// Returns: the child resource or undefined if not found.
	GetResource(pathPart *string) IResource
	// Gets or create all resources leading up to the specified path.
	//
	// - Path may only start with "/" if this method is called on the root resource.
	// - All resources are created using default options.
	//
	// Returns: a new or existing resource.
	ResourceForPath(path *string) Resource
	// The rest API that this resource is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	Api() IRestApi
	// Default options for CORS preflight OPTIONS method.
	DefaultCorsPreflightOptions() *CorsOptions
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration() Integration
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions() *MethodOptions
	// The parent of this resource or undefined for the root resource.
	ParentResource() IResource
	// The full path of this resource.
	Path() *string
	// The ID of the resource.
	ResourceId() *string
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IResource) AddCorsPreflight(options *CorsOptions) Method {
	if err := i.validateAddCorsPreflightParameters(options); err != nil {
		panic(err)
	}
	var returns Method

	_jsii_.Invoke(
		i,
		"addCorsPreflight",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddMethod(httpMethod *string, target Integration, options *MethodOptions) Method {
	if err := i.validateAddMethodParameters(httpMethod, options); err != nil {
		panic(err)
	}
	var returns Method

	_jsii_.Invoke(
		i,
		"addMethod",
		[]interface{}{httpMethod, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddProxy(options *ProxyResourceOptions) ProxyResource {
	if err := i.validateAddProxyParameters(options); err != nil {
		panic(err)
	}
	var returns ProxyResource

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) AddResource(pathPart *string, options *ResourceOptions) Resource {
	if err := i.validateAddResourceParameters(pathPart, options); err != nil {
		panic(err)
	}
	var returns Resource

	_jsii_.Invoke(
		i,
		"addResource",
		[]interface{}{pathPart, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) GetResource(pathPart *string) IResource {
	if err := i.validateGetResourceParameters(pathPart); err != nil {
		panic(err)
	}
	var returns IResource

	_jsii_.Invoke(
		i,
		"getResource",
		[]interface{}{pathPart},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResource) ResourceForPath(path *string) Resource {
	if err := i.validateResourceForPathParameters(path); err != nil {
		panic(err)
	}
	var returns Resource

	_jsii_.Invoke(
		i,
		"resourceForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResource) Api() IRestApi {
	var returns IRestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultCorsPreflightOptions() *CorsOptions {
	var returns *CorsOptions
	_jsii_.Get(
		j,
		"defaultCorsPreflightOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultIntegration() Integration {
	var returns Integration
	_jsii_.Get(
		j,
		"defaultIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) DefaultMethodOptions() *MethodOptions {
	var returns *MethodOptions
	_jsii_.Get(
		j,
		"defaultMethodOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ParentResource() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"parentResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

