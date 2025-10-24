package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages AWS tags for all resources within a construct scope.
//
// Example:
//   var mesh Mesh
//   var service Service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type Tags interface {
	// Add tags to the node of a construct and all its the taggable children.
	//
	// ## Tagging and CloudFormation Stacks
	//
	// If the feature flag `@aws-cdk/core:explicitStackTags` is set to `true`
	// (recommended modern behavior), Stacks will not automatically be tagged.
	// Stack tags should be configured on Stacks directly (preferred), or
	// you must explicitly include the resource type `aws:cdk:stack` in the
	// `includeResourceTypes` array.
	//
	// If the feature flag is set to `false` (legacy behavior) then both Stacks
	// and resources in the indicated scope will both be tagged by default, which
	// leads to tags being applied twice (once in the template, and then once
	// again automatically by CloudFormation as part of the stack deployment).
	// That behavior leads to loss of control as `excludeResourceTypes` will
	// prevent tags from appearing in the template, but they will still be
	// applied to the Stack and hence CloudFormation will still apply them
	// to the resource.
	Add(key *string, value *string, props *TagProps)
	// remove tags to the node of a construct and all its the taggable children.
	Remove(key *string, props *TagProps)
}

// The jsii proxy struct for Tags
type jsiiProxy_Tags struct {
	_ byte // padding
}

// Returns the tags API for this scope.
func Tags_Of(scope constructs.IConstruct) Tags {
	_init_.Initialize()

	if err := validateTags_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns Tags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Tags",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tags) Add(key *string, value *string, props *TagProps) {
	if err := t.validateAddParameters(key, value, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"add",
		[]interface{}{key, value, props},
	)
}

func (t *jsiiProxy_Tags) Remove(key *string, props *TagProps) {
	if err := t.validateRemoveParameters(key, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"remove",
		[]interface{}{key, props},
	)
}

