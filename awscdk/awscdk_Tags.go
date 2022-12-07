// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages AWS tags for all resources within a construct scope.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8080),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				path: jsii.String("/ping"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   			timeout: &httpTimeout{
//   				idle: cdk.*duration.seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.file(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.tags.of(node).add(jsii.String("Environment"), jsii.String("Dev"))
//
type Tags interface {
	// add tags to the node of a construct and all its the taggable children.
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

