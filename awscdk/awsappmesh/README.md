# AWS App Mesh Construct Library

AWS App Mesh is a service mesh based on the [Envoy](https://www.envoyproxy.io/) proxy that makes it easy to monitor and control microservices. App Mesh standardizes how your microservices communicate, giving you end-to-end visibility and helping to ensure high-availability for your applications.

App Mesh gives you consistent visibility and network traffic controls for every microservice in an application.

App Mesh supports microservice applications that use service discovery naming for their components. To use App Mesh, you must have an existing application running on AWS Fargate, Amazon ECS, Amazon EKS, Kubernetes on AWS, or Amazon EC2.

For further information on **AWS App Mesh**, visit the [AWS App Mesh Documentation](https://docs.aws.amazon.com/app-mesh/index.html).

## Create the App and Stack

```go
// Example automatically generated from non-compiling source. May contain errors.
app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("stack"))
```

## Creating the Mesh

A service mesh is a logical boundary for network traffic between the services that reside within it.

After you create your service mesh, you can create virtual services, virtual nodes, virtual routers, and routes to distribute traffic between the applications in your mesh.

The following example creates the `AppMesh` service mesh with the default egress filter of `DROP_ALL`. See [the AWS CloudFormation `EgressFilter` resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-egressfilter.html) for more info on egress filters.

```go
// Example automatically generated from non-compiling source. May contain errors.
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &meshProps{
	meshName: jsii.String("myAwsMesh"),
})
```

The mesh can instead be created with the `ALLOW_ALL` egress filter by providing the `egressFilter` property.

```go
// Example automatically generated from non-compiling source. May contain errors.
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &meshProps{
	meshName: jsii.String("myAwsMesh"),
	egressFilter: appmesh.meshFilterType_ALLOW_ALL,
})
```

A mesh with an IP preference can be created by providing the property `serviceDiscovery` that specifes an `ipPreference`.

```go
// Example automatically generated from non-compiling source. May contain errors.
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &meshProps{
	meshName: jsii.String("myAwsMesh"),
	serviceDiscovery: &meshServiceDiscovery{
		ipPreference: appmesh.ipPreference_IPV4_ONLY,
	},
})
```

## Adding VirtualRouters

A *mesh* uses  *virtual routers* as logical units to route requests to *virtual nodes*.

Virtual routers handle traffic for one or more virtual services within your mesh.
After you create a virtual router, you can create and associate routes to your virtual router that direct incoming requests to different virtual nodes.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh

router := mesh.addVirtualRouter(jsii.String("router"), &virtualRouterBaseProps{
	listeners: []virtualRouterListener{
		appmesh.*virtualRouterListener.http(jsii.Number(8080)),
	},
})
```

Note that creating the router using the `addVirtualRouter()` method places it in the same stack as the mesh
(which might be different from the current stack).
The router can also be created using the `VirtualRouter` constructor (passing in the mesh) instead of calling the `addVirtualRouter()` method.
This is particularly useful when splitting your resources between many stacks: for example, defining the mesh itself as part of an infrastructure stack, but defining the other resources, such as routers, in the application stack:

```go
// Example automatically generated from non-compiling source. May contain errors.
var infraStack stack
var appStack stack


mesh := appmesh.NewMesh(infraStack, jsii.String("AppMesh"), &meshProps{
	meshName: jsii.String("myAwsMesh"),
	egressFilter: appmesh.meshFilterType_ALLOW_ALL,
})

// the VirtualRouter will belong to 'appStack',
// even though the Mesh belongs to 'infraStack'
router := appmesh.NewVirtualRouter(appStack, jsii.String("router"), &virtualRouterProps{
	mesh: mesh,
	 // notice that mesh is a required property when creating a router with the 'new' statement
	listeners: []virtualRouterListener{
		appmesh.*virtualRouterListener.http(jsii.Number(8081)),
	},
})
```

The same is true for other `add*()` methods in the App Mesh construct library.

The `VirtualRouterListener` class lets you define protocol-specific listeners.
The `http()`, `http2()`, `grpc()` and `tcp()` methods create listeners for the named protocols.
They accept a single parameter that defines the port to on which requests will be matched.
The port parameter defaults to 8080 if omitted.

## Adding a VirtualService

A *virtual service* is an abstraction of a real service that is provided by a virtual node directly, or indirectly by means of a virtual router. Dependent services call your virtual service by its `virtualServiceName`, and those requests are routed to the virtual node or virtual router specified as the provider for the virtual service.

We recommend that you use the service discovery name of the real service that you're targeting (such as `my-service.default.svc.cluster.local`).

When creating a virtual service:

* If you want the virtual service to spread traffic across multiple virtual nodes, specify a virtual router.
* If you want the virtual service to reach a virtual node directly, without a virtual router, specify a virtual node.

Adding a virtual router as the provider:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter


appmesh.NewVirtualService(this, jsii.String("virtual-service"), &virtualServiceProps{
	virtualServiceName: jsii.String("my-service.default.svc.cluster.local"),
	 // optional
	virtualServiceProvider: appmesh.virtualServiceProvider.virtualRouter(router),
})
```

Adding a virtual node as the provider:

```go
// Example automatically generated from non-compiling source. May contain errors.
var node virtualNode


appmesh.NewVirtualService(this, jsii.String("virtual-service"), &virtualServiceProps{
	virtualServiceName: jsii.String("my-service.default.svc.cluster.local"),
	 // optional
	virtualServiceProvider: appmesh.virtualServiceProvider.virtualNode(node),
})
```

## Adding a VirtualNode

A *virtual node* acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment.

When you create a virtual node, accept inbound traffic by specifying a *listener*. Outbound traffic that your virtual node expects to send should be specified as a *back end*.

The response metadata for your new virtual node contains the Amazon Resource Name (ARN) that is associated with the virtual node. Set this value (either the full ARN or the truncated resource name) as the `APPMESH_VIRTUAL_NODE_NAME` environment variable for your task group's Envoy proxy container in your task definition or pod spec. For example, the value could be `mesh/default/virtualNode/simpleapp`. This is then mapped to the `node.id` and `node.cluster` Envoy parameters.

> **Note**
> If you require your Envoy stats or tracing to use a different name, you can override the `node.cluster` value that is set by `APPMESH_VIRTUAL_NODE_NAME` with the `APPMESH_VIRTUAL_NODE_CLUSTER` environment variable.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh
vpc := ec2.NewVpc(this, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
	vpc: vpc,
	name: jsii.String("domain.local"),
})
service := namespace.createService(jsii.String("Svc"))
node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
			port: jsii.Number(8081),
			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
				healthyThreshold: jsii.Number(3),
				interval: cdk.duration.seconds(jsii.Number(5)),
				 // minimum
				path: jsii.String("/health-check-path"),
				timeout: cdk.*duration.seconds(jsii.Number(2)),
				 // minimum
				unhealthyThreshold: jsii.Number(2),
			}),
		}),
	},
	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
})
```

Create a `VirtualNode` with the constructor and add tags.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh
var service service


node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
			port: jsii.Number(8080),
			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
				healthyThreshold: jsii.Number(3),
				interval: cdk.duration.seconds(jsii.Number(5)),
				path: jsii.String("/ping"),
				timeout: cdk.*duration.seconds(jsii.Number(2)),
				unhealthyThreshold: jsii.Number(2),
			}),
			timeout: &httpTimeout{
				idle: cdk.*duration.seconds(jsii.Number(5)),
			},
		}),
	},
	backendDefaults: &backendDefaults{
		tlsClientPolicy: &tlsClientPolicy{
			validation: &tlsValidation{
				trust: appmesh.tlsValidationTrust.file(jsii.String("/keys/local_cert_chain.pem")),
			},
		},
	},
	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
})

cdk.tags.of(node).add(jsii.String("Environment"), jsii.String("Dev"))
```

Create a `VirtualNode` with the constructor and add backend virtual service.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh
var router virtualRouter
var service service


node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
			port: jsii.Number(8080),
			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
				healthyThreshold: jsii.Number(3),
				interval: cdk.duration.seconds(jsii.Number(5)),
				path: jsii.String("/ping"),
				timeout: cdk.*duration.seconds(jsii.Number(2)),
				unhealthyThreshold: jsii.Number(2),
			}),
			timeout: &httpTimeout{
				idle: cdk.*duration.seconds(jsii.Number(5)),
			},
		}),
	},
	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
})

virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &virtualServiceProps{
	virtualServiceProvider: appmesh.virtualServiceProvider.virtualRouter(router),
	virtualServiceName: jsii.String("service1.domain.local"),
})

node.addBackend(appmesh.backend.virtualService(virtualService))
```

The `listeners` property can be left blank and added later with the `node.addListener()` method. The `serviceDiscovery` property must be specified when specifying a listener.

The `backends` property can be added with `node.addBackend()`. In the example, we define a virtual service and add it to the virtual node to allow egress traffic to other nodes.

The `backendDefaults` property is added to the node while creating the virtual node. These are the virtual node's default settings for all backends.

The `VirtualNode.addBackend()` method is especially useful if you want to create a circular traffic flow by having a Virtual Service as a backend whose provider is that same Virtual Node:

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh


node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
})

virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &virtualServiceProps{
	virtualServiceProvider: appmesh.virtualServiceProvider.virtualNode(node),
	virtualServiceName: jsii.String("service1.domain.local"),
})

node.addBackend(appmesh.backend.virtualService(virtualService))
```

### Adding TLS to a listener

The `tls` property specifies TLS configuration when creating a listener for a virtual node or a virtual gateway.
Provide the TLS certificate to the proxy in one of the following ways:

* A certificate from AWS Certificate Manager (ACM).
* A customer-provided certificate (specify a `certificateChain` path file and a `privateKey` file path).
* A certificate provided by a Secrets Discovery Service (SDS) endpoint over local Unix Domain Socket (specify its `secretName`).

```go
// Example automatically generated from non-compiling source. May contain errors.
// A Virtual Node with listener TLS from an ACM provided certificate
var cert certificate
var mesh mesh


node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
			port: jsii.Number(80),
			tls: &listenerTlsOptions{
				mode: appmesh.tlsMode_STRICT,
				certificate: appmesh.tlsCertificate.acm(cert),
			},
		}),
	},
})

// A Virtual Gateway with listener TLS from a customer provided file certificate
gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
	mesh: mesh,
	listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
			port: jsii.Number(8080),
			tls: &listenerTlsOptions{
				mode: appmesh.*tlsMode_STRICT,
				certificate: appmesh.*tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
			},
		}),
	},
	virtualGatewayName: jsii.String("gateway"),
})

// A Virtual Gateway with listener TLS from a SDS provided certificate
gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &virtualGatewayProps{
	mesh: mesh,
	listeners: []*virtualGatewayListener{
		appmesh.*virtualGatewayListener.http2(&http2GatewayListenerOptions{
			port: jsii.Number(8080),
			tls: &listenerTlsOptions{
				mode: appmesh.*tlsMode_STRICT,
				certificate: appmesh.*tlsCertificate.sds(jsii.String("secrete_certificate")),
			},
		}),
	},
	virtualGatewayName: jsii.String("gateway2"),
})
```

### Adding mutual TLS authentication

Mutual TLS authentication is an optional component of TLS that offers two-way peer authentication.
To enable mutual TLS authentication, add the `mutualTlsCertificate` property to TLS client policy and/or the `mutualTlsValidation` property to your TLS listener.

`tls.mutualTlsValidation` and `tlsClientPolicy.mutualTlsCertificate` can be sourced from either:

* A customer-provided certificate (specify a `certificateChain` path file and a `privateKey` file path).
* A certificate provided by a Secrets Discovery Service (SDS) endpoint over local Unix Domain Socket (specify its `secretName`).

> **Note**
> Currently, a certificate from AWS Certificate Manager (ACM) cannot be used for mutual TLS authentication.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh


node1 := appmesh.NewVirtualNode(this, jsii.String("node1"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
			port: jsii.Number(80),
			tls: &listenerTlsOptions{
				mode: appmesh.tlsMode_STRICT,
				certificate: appmesh.tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
				// Validate a file client certificates to enable mutual TLS authentication when a client provides a certificate.
				mutualTlsValidation: &mutualTlsValidation{
					trust: appmesh.tlsValidationTrust.file(jsii.String("path-to-certificate")),
				},
			},
		}),
	},
})

certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
node2 := appmesh.NewVirtualNode(this, jsii.String("node2"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.*serviceDiscovery.dns(jsii.String("node2")),
	backendDefaults: &backendDefaults{
		tlsClientPolicy: &tlsClientPolicy{
			ports: []*f64{
				jsii.Number(8080),
				jsii.Number(8081),
			},
			validation: &tlsValidation{
				subjectAlternativeNames: appmesh.subjectAlternativeNames.matchingExactly(jsii.String("mesh-endpoint.apps.local")),
				trust: appmesh.*tlsValidationTrust.acm([]iCertificateAuthority{
					acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
				}),
			},
			// Provide a SDS client certificate when a server requests it and enable mutual TLS authentication.
			mutualTlsCertificate: appmesh.*tlsCertificate.sds(jsii.String("secret_certificate")),
		},
	},
})
```

### Adding outlier detection to a Virtual Node listener

The `outlierDetection` property adds outlier detection to a Virtual Node listener. The properties
`baseEjectionDuration`, `interval`, `maxEjectionPercent`, and `maxServerErrors` are required.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh
// Cloud Map service discovery is currently required for host ejection by outlier detection
vpc := ec2.NewVpc(this, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
	vpc: vpc,
	name: jsii.String("domain.local"),
})
service := namespace.createService(jsii.String("Svc"))
node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
			outlierDetection: &outlierDetection{
				baseEjectionDuration: cdk.duration.seconds(jsii.Number(10)),
				interval: cdk.*duration.seconds(jsii.Number(30)),
				maxEjectionPercent: jsii.Number(50),
				maxServerErrors: jsii.Number(5),
			},
		}),
	},
})
```

### Adding a connection pool to a listener

The `connectionPool` property can be added to a Virtual Node listener or Virtual Gateway listener to add a request connection pool. Each listener protocol type has its own connection pool properties.

```go
// Example automatically generated from non-compiling source. May contain errors.
// A Virtual Node with a gRPC listener with a connection pool set
var mesh mesh

node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
	mesh: mesh,
	// DNS service discovery can optionally specify the DNS response type as either LOAD_BALANCER or ENDPOINTS.
	// LOAD_BALANCER means that the DNS resolver returns a loadbalanced set of endpoints,
	// whereas ENDPOINTS means that the DNS resolver is returning all the endpoints.
	// By default, the response type is assumed to be LOAD_BALANCER
	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node"), appmesh.dnsResponseType_ENDPOINTS),
	listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
			port: jsii.Number(80),
			connectionPool: &httpConnectionPool{
				maxConnections: jsii.Number(100),
				maxPendingRequests: jsii.Number(10),
			},
		}),
	},
})

// A Virtual Gateway with a gRPC listener with a connection pool set
gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
	mesh: mesh,
	listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
			port: jsii.Number(8080),
			connectionPool: &grpcConnectionPool{
				maxRequests: jsii.Number(10),
			},
		}),
	},
	virtualGatewayName: jsii.String("gateway"),
})
```

### Adding an IP Preference to a Virtual Node

An `ipPreference` can be specified as part of a Virtual Node's service discovery. An IP preference defines how clients for this Virtual Node will interact with it.

There a four different IP preferences available to use which each specify what IP versions this Virtual Node will use and prefer.

* `IPv4_ONLY` - Only use IPv4. For CloudMap service discovery, only IPv4 addresses returned from CloudMap will be used. For DNS service discovery, Envoy's DNS resolver will only resolve DNS queries for IPv4.
* `IPv4_PREFERRED` - Prefer IPv4 and fall back to IPv6. For CloudMap service discovery, an IPv4 address will be used if returned from CloudMap. Otherwise, an IPv6 address will be used if available. For DNS service discovery, Envoy's DNS resolver will first attempt to resolve DNS queries using IPv4 and fall back to IPv6.
* `IPv6_ONLY` - Only use IPv6. For CloudMap service discovery, only IPv6 addresses returned from CloudMap will be used. For DNS service discovery, Envoy's DNS resolver will only resolve DNS queries for IPv6.
* `IPv6_PREFERRED` - Prefer IPv6 and fall back to IPv4. For CloudMap service discovery, an IPv6 address will be used if returned from CloudMap. Otherwise, an IPv4 address will be used if available. For DNS service discovery, Envoy's DNS resolver will first attempt to resolve DNS queries using IPv6 and fall back to IPv4.

```go
// Example automatically generated from non-compiling source. May contain errors.
mesh := appmesh.NewMesh(stack, jsii.String("mesh"), &meshProps{
	meshName: jsii.String("mesh-with-preference"),
})

// Virtual Node with DNS service discovery and an IP preference
dnsNode := appmesh.NewVirtualNode(stack, jsii.String("dns-node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("test"), appmesh.dnsResponseType_LOAD_BALANCER, appmesh.ipPreference_IPV4_ONLY),
})

// Virtual Node with CloudMap service discovery and an IP preference
vpc := ec2.NewVpc(stack, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(stack, jsii.String("test-namespace"), &privateDnsNamespaceProps{
	vpc: vpc,
	name: jsii.String("domain.local"),
})
service := namespace.createService(jsii.String("Svc"))

instanceAttribute := map[string]interface{}{
}
instanceAttribute.testKey = "testValue"

cloudmapNode := appmesh.NewVirtualNode(stack, jsii.String("cloudmap-node"), &virtualNodeProps{
	mesh: mesh,
	serviceDiscovery: appmesh.*serviceDiscovery.cloudMap(service, instanceAttribute, appmesh.*ipPreference_IPV4_ONLY),
})
```

## Adding a Route

A *route* matches requests with an associated virtual router and distributes traffic to its associated virtual nodes.
The route distributes matching requests to one or more target virtual nodes with relative weighting.

The `RouteSpec` class lets you define protocol-specific route specifications.
The `tcp()`, `http()`, `http2()`, and `grpc()` methods create a specification for the named protocols.

For HTTP-based routes, the match field can match on path (prefix, exact, or regex), HTTP method, scheme,
HTTP headers, and query parameters. By default, HTTP-based routes match all requests.

For gRPC-based routes, the match field can  match on service name, method name, and metadata.
When specifying the method name, the service name must also be specified.

For example, here's how to add an HTTP route that matches based on a prefix of the URL path:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.http(&httpRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		match: &httpRouteMatch{
			// Path that is passed to this method must start with '/'.
			path: appmesh.httpRoutePathMatch.startsWith(jsii.String("/path-to-app")),
		},
	}),
})
```

Add an HTTP2 route that matches based on exact path, method, scheme, headers, and query parameters:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http2"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		match: &httpRouteMatch{
			path: appmesh.httpRoutePathMatch.exactly(jsii.String("/exact")),
			method: appmesh.httpRouteMethod_POST,
			protocol: appmesh.httpRouteProtocol_HTTPS,
			headers: []headerMatch{
				appmesh.*headerMatch.valueIs(jsii.String("Content-Type"), jsii.String("application/json")),
				appmesh.*headerMatch.valueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
			},
			queryParameters: []queryParameterMatch{
				appmesh.*queryParameterMatch.valueIs(jsii.String("query-field"), jsii.String("value")),
			},
		},
	}),
})
```

Add a single route with two targets and split traffic 50/50:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.http(&httpRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
				weight: jsii.Number(50),
			},
			&weightedTarget{
				virtualNode: node,
				weight: jsii.Number(50),
			},
		},
		match: &httpRouteMatch{
			path: appmesh.httpRoutePathMatch.startsWith(jsii.String("/path-to-app")),
		},
	}),
})
```

Add an http2 route with retries:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http2-retry"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		retryPolicy: &httpRetryPolicy{
			// Retry if the connection failed
			tcpRetryEvents: []cONNECTION_ERROR{
				appmesh.tcpRetryEvent_*cONNECTION_ERROR,
			},
			// Retry if HTTP responds with a gateway error (502, 503, 504)
			httpRetryEvents: []httpRetryEvent{
				appmesh.*httpRetryEvent_GATEWAY_ERROR,
			},
			// Retry five times
			retryAttempts: jsii.Number(5),
			// Use a 1 second timeout per retry
			retryTimeout: cdk.duration.seconds(jsii.Number(1)),
		},
	}),
})
```

Add a gRPC route with retries:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-grpc-retry"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		match: &grpcRouteMatch{
			serviceName: jsii.String("servicename"),
		},
		retryPolicy: &grpcRetryPolicy{
			tcpRetryEvents: []cONNECTION_ERROR{
				appmesh.tcpRetryEvent_*cONNECTION_ERROR,
			},
			httpRetryEvents: []httpRetryEvent{
				appmesh.*httpRetryEvent_GATEWAY_ERROR,
			},
			// Retry if gRPC responds that the request was cancelled, a resource
			// was exhausted, or if the service is unavailable
			grpcRetryEvents: []grpcRetryEvent{
				appmesh.*grpcRetryEvent_CANCELLED,
				appmesh.*grpcRetryEvent_RESOURCE_EXHAUSTED,
				appmesh.*grpcRetryEvent_UNAVAILABLE,
			},
			retryAttempts: jsii.Number(5),
			retryTimeout: cdk.duration.seconds(jsii.Number(1)),
		},
	}),
})
```

Add an gRPC route that matches based on method name and metadata:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-grpc-retry"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		match: &grpcRouteMatch{
			// When method name is specified, service name must be also specified.
			methodName: jsii.String("methodname"),
			serviceName: jsii.String("servicename"),
			metadata: []headerMatch{
				appmesh.*headerMatch.valueStartsWith(jsii.String("Content-Type"), jsii.String("application/")),
				appmesh.*headerMatch.valueDoesNotStartWith(jsii.String("Content-Type"), jsii.String("text/")),
			},
		},
	}),
})
```

Add a gRPC route with timeout:

```go
// Example automatically generated from non-compiling source. May contain errors.
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &routeBaseProps{
	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
		weightedTargets: []weightedTarget{
			&weightedTarget{
				virtualNode: node,
			},
		},
		match: &grpcRouteMatch{
			serviceName: jsii.String("my-service.default.svc.cluster.local"),
		},
		timeout: &grpcTimeout{
			idle: cdk.duration.seconds(jsii.Number(2)),
			perRequest: cdk.*duration.seconds(jsii.Number(1)),
		},
	}),
})
```

## Adding a Virtual Gateway

A *virtual gateway* allows resources outside your mesh to communicate with resources inside your mesh.
The virtual gateway represents an Envoy proxy running in an Amazon ECS task, in a Kubernetes service, or on an Amazon EC2 instance.
Unlike a virtual node, which represents Envoy running with an application, a virtual gateway represents Envoy deployed by itself.

A virtual gateway is similar to a virtual node in that it has a listener that accepts traffic for a particular port and protocol (HTTP, HTTP2, gRPC).
Traffic received by the virtual gateway is directed to other services in your mesh
using rules defined in gateway routes which can be added to your virtual gateway.

Create a virtual gateway with the constructor:

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh

certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"

gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
	mesh: mesh,
	listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener.http(&httpGatewayListenerOptions{
			port: jsii.Number(443),
			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
				interval: cdk.duration.seconds(jsii.Number(10)),
			}),
		}),
	},
	backendDefaults: &backendDefaults{
		tlsClientPolicy: &tlsClientPolicy{
			ports: []*f64{
				jsii.Number(8080),
				jsii.Number(8081),
			},
			validation: &tlsValidation{
				trust: appmesh.tlsValidationTrust.acm([]iCertificateAuthority{
					acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
				}),
			},
		},
	},
	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
	virtualGatewayName: jsii.String("virtualGateway"),
})
```

Add a virtual gateway directly to the mesh:

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh


gateway := mesh.addVirtualGateway(jsii.String("gateway"), &virtualGatewayBaseProps{
	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
	virtualGatewayName: jsii.String("virtualGateway"),
	listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener.http(&httpGatewayListenerOptions{
			port: jsii.Number(443),
			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
				interval: cdk.duration.seconds(jsii.Number(10)),
			}),
		}),
	},
})
```

The `listeners` field defaults to an HTTP Listener on port 8080 if omitted.
A gateway route can be added using the `gateway.addGatewayRoute()` method.

The `backendDefaults` property, provided when creating the virtual gateway, specifies the virtual gateway's default settings for all backends.

## Adding a Gateway Route

A *gateway route* is attached to a virtual gateway and routes matching traffic to an existing virtual service.

For HTTP-based gateway routes, the `match` field can be used to match on
path (prefix, exact, or regex), HTTP method, host name, HTTP headers, and query parameters.
By default, HTTP-based gateway routes match all requests.

```go
// Example automatically generated from non-compiling source. May contain errors.
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http"), &gatewayRouteBaseProps{
	routeSpec: appmesh.gatewayRouteSpec.http(&httpGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &httpGatewayRouteMatch{
			path: appmesh.httpGatewayRoutePathMatch.regex(jsii.String("regex")),
		},
	}),
})
```

For gRPC-based gateway routes, the `match` field can be used to match on service name, host name, and metadata.

```go
// Example automatically generated from non-compiling source. May contain errors.
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &gatewayRouteBaseProps{
	routeSpec: appmesh.gatewayRouteSpec.grpc(&grpcGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &grpcGatewayRouteMatch{
			hostname: appmesh.gatewayRouteHostnameMatch.endsWith(jsii.String(".example.com")),
		},
	}),
})
```

For HTTP based gateway routes, App Mesh automatically rewrites the matched prefix path in Gateway Route to “/”.
This automatic rewrite configuration can be overwritten in following ways:

```go
// Example automatically generated from non-compiling source. May contain errors.
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http"), &gatewayRouteBaseProps{
	routeSpec: appmesh.gatewayRouteSpec.http(&httpGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &httpGatewayRouteMatch{
			// This disables the default rewrite to '/', and retains original path.
			path: appmesh.httpGatewayRoutePathMatch.startsWith(jsii.String("/path-to-app/"), jsii.String("")),
		},
	}),
})

gateway.addGatewayRoute(jsii.String("gateway-route-http-1"), &gatewayRouteBaseProps{
	routeSpec: appmesh.*gatewayRouteSpec.http(&httpGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &httpGatewayRouteMatch{
			// If the request full path is '/path-to-app/xxxxx', this rewrites the path to '/rewrittenUri/xxxxx'.
			// Please note both `prefixPathMatch` and `rewriteTo` must start and end with the `/` character.
			path: appmesh.*httpGatewayRoutePathMatch.startsWith(jsii.String("/path-to-app/"), jsii.String("/rewrittenUri/")),
		},
	}),
})
```

If matching other path (exact or regex), only specific rewrite path can be specified.
Unlike `startsWith()` method above, no default rewrite is performed.

```go
// Example automatically generated from non-compiling source. May contain errors.
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http-2"), &gatewayRouteBaseProps{
	routeSpec: appmesh.gatewayRouteSpec.http(&httpGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &httpGatewayRouteMatch{
			// This rewrites the path from '/test' to '/rewrittenPath'.
			path: appmesh.httpGatewayRoutePathMatch.exactly(jsii.String("/test"), jsii.String("/rewrittenPath")),
		},
	}),
})
```

For HTTP/gRPC based routes, App Mesh automatically rewrites
the original request received at the Virtual Gateway to the destination Virtual Service name.
This default host name rewrite can be configured by specifying the rewrite rule as one of the `match` property:

```go
// Example automatically generated from non-compiling source. May contain errors.
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &gatewayRouteBaseProps{
	routeSpec: appmesh.gatewayRouteSpec.grpc(&grpcGatewayRouteSpecOptions{
		routeTarget: virtualService,
		match: &grpcGatewayRouteMatch{
			hostname: appmesh.gatewayRouteHostnameMatch.exactly(jsii.String("example.com")),
			// This disables the default rewrite to virtual service name and retain original request.
			rewriteRequestHostname: jsii.Boolean(false),
		},
	}),
})
```

## Importing Resources

Each App Mesh resource class comes with two static methods, `from<Resource>Arn` and `from<Resource>Attributes` (where `<Resource>` is replaced with the resource name, such as `VirtualNode`) for importing a reference to an existing App Mesh resource.
These imported resources can be used with other resources in your mesh as if they were defined directly in your CDK application.

```go
// Example automatically generated from non-compiling source. May contain errors.
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh/virtualNode/testNode"
appmesh.virtualNode.fromVirtualNodeArn(this, jsii.String("importedVirtualNode"), arn)
```

```go
// Example automatically generated from non-compiling source. May contain errors.
virtualNodeName := "my-virtual-node"
appmesh.virtualNode.fromVirtualNodeAttributes(this, jsii.String("imported-virtual-node"), &virtualNodeAttributes{
	mesh: appmesh.mesh.fromMeshName(this, jsii.String("Mesh"), jsii.String("testMesh")),
	virtualNodeName: virtualNodeName,
})
```

To import a mesh, again there are two static methods, `fromMeshArn` and `fromMeshName`.

```go
// Example automatically generated from non-compiling source. May contain errors.
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh"
appmesh.mesh.fromMeshArn(this, jsii.String("imported-mesh"), arn)
```

```go
// Example automatically generated from non-compiling source. May contain errors.
appmesh.mesh.fromMeshName(this, jsii.String("imported-mesh"), jsii.String("abc"))
```

## IAM Grants

`VirtualNode` and `VirtualGateway` provide `grantStreamAggregatedResources` methods that grant identities that are running
Envoy access to stream generated config from App Mesh.

```go
// Example automatically generated from non-compiling source. May contain errors.
var mesh mesh

gateway := appmesh.NewVirtualGateway(this, jsii.String("testGateway"), &virtualGatewayProps{
	mesh: mesh,
})
envoyUser := iam.NewUser(this, jsii.String("envoyUser"))

/**
 * This will grant `grantStreamAggregatedResources` ONLY for this gateway.
 */
gateway.grantStreamAggregatedResources(envoyUser)
```

## Adding Resources to shared meshes

A shared mesh allows resources created by different accounts to communicate with each other in the same mesh:

```go
// Example automatically generated from non-compiling source. May contain errors.
// This is the ARN for the mesh from different AWS IAM account ID.
// Ensure mesh is properly shared with your account. For more details, see: https://github.com/aws/aws-cdk/issues/15404
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh"
sharedMesh := appmesh.mesh.fromMeshArn(this, jsii.String("imported-mesh"), arn)

// This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
// This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
appmesh.NewVirtualNode(this, jsii.String("test-node"), &virtualNodeProps{
	mesh: sharedMesh,
})
```
