# AWS App Mesh Construct Library

AWS App Mesh is a service mesh based on the [Envoy](https://www.envoyproxy.io/) proxy that makes it easy to monitor and control microservices. App Mesh standardizes how your microservices communicate, giving you end-to-end visibility and helping to ensure high-availability for your applications.

App Mesh gives you consistent visibility and network traffic controls for every microservice in an application.

App Mesh supports microservice applications that use service discovery naming for their components. To use App Mesh, you must have an existing application running on AWS Fargate, Amazon ECS, Amazon EKS, Kubernetes on AWS, or Amazon EC2.

For further information on **AWS App Mesh**, visit the [AWS App Mesh Documentation](https://docs.aws.amazon.com/app-mesh/index.html).

## Create the App and Stack

```go
app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("stack"))
```

## Creating the Mesh

A service mesh is a logical boundary for network traffic between the services that reside within it.

After you create your service mesh, you can create virtual services, virtual nodes, virtual routers, and routes to distribute traffic between the applications in your mesh.

The following example creates the `AppMesh` service mesh with the default egress filter of `DROP_ALL`. See [the AWS CloudFormation `EgressFilter` resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-egressfilter.html) for more info on egress filters.

```go
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
	MeshName: jsii.String("myAwsMesh"),
})
```

The mesh can instead be created with the `ALLOW_ALL` egress filter by providing the `egressFilter` property.

```go
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
	MeshName: jsii.String("myAwsMesh"),
	EgressFilter: appmesh.MeshFilterType_ALLOW_ALL,
})
```

A mesh with an IP preference can be created by providing the property `serviceDiscovery` that specifes an `ipPreference`.

```go
mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
	MeshName: jsii.String("myAwsMesh"),
	ServiceDiscovery: &MeshServiceDiscovery{
		IpPreference: appmesh.IpPreference_IPV4_ONLY,
	},
})
```

## Adding VirtualRouters

A *mesh* uses  *virtual routers* as logical units to route requests to *virtual nodes*.

Virtual routers handle traffic for one or more virtual services within your mesh.
After you create a virtual router, you can create and associate routes to your virtual router that direct incoming requests to different virtual nodes.

```go
var mesh mesh

router := mesh.addVirtualRouter(jsii.String("router"), &VirtualRouterBaseProps{
	Listeners: []virtualRouterListener{
		appmesh.*virtualRouterListener_Http(jsii.Number(8080)),
	},
})
```

Note that creating the router using the `addVirtualRouter()` method places it in the same stack as the mesh
(which might be different from the current stack).
The router can also be created using the `VirtualRouter` constructor (passing in the mesh) instead of calling the `addVirtualRouter()` method.
This is particularly useful when splitting your resources between many stacks: for example, defining the mesh itself as part of an infrastructure stack, but defining the other resources, such as routers, in the application stack:

```go
var infraStack stack
var appStack stack


mesh := appmesh.NewMesh(infraStack, jsii.String("AppMesh"), &MeshProps{
	MeshName: jsii.String("myAwsMesh"),
	EgressFilter: appmesh.MeshFilterType_ALLOW_ALL,
})

// the VirtualRouter will belong to 'appStack',
// even though the Mesh belongs to 'infraStack'
router := appmesh.NewVirtualRouter(appStack, jsii.String("router"), &VirtualRouterProps{
	Mesh: Mesh,
	 // notice that mesh is a required property when creating a router with the 'new' statement
	Listeners: []virtualRouterListener{
		appmesh.*virtualRouterListener_Http(jsii.Number(8081)),
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
var router virtualRouter


appmesh.NewVirtualService(this, jsii.String("virtual-service"), &VirtualServiceProps{
	VirtualServiceName: jsii.String("my-service.default.svc.cluster.local"),
	 // optional
	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualRouter(router),
})
```

Adding a virtual node as the provider:

```go
var node virtualNode


appmesh.NewVirtualService(this, jsii.String("virtual-service"), &VirtualServiceProps{
	VirtualServiceName: jsii.String("my-service.default.svc.cluster.local"),
	 // optional
	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualNode(node),
})
```

## Adding a VirtualNode

A *virtual node* acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment.

When you create a virtual node, accept inbound traffic by specifying a *listener*. Outbound traffic that your virtual node expects to send should be specified as a *back end*.

The response metadata for your new virtual node contains the Amazon Resource Name (ARN) that is associated with the virtual node. Set this value (either the full ARN or the truncated resource name) as the `APPMESH_VIRTUAL_NODE_NAME` environment variable for your task group's Envoy proxy container in your task definition or pod spec. For example, the value could be `mesh/default/virtualNode/simpleapp`. This is then mapped to the `node.id` and `node.cluster` Envoy parameters.

> **Note**
> If you require your Envoy stats or tracing to use a different name, you can override the `node.cluster` value that is set by `APPMESH_VIRTUAL_NODE_NAME` with the `APPMESH_VIRTUAL_NODE_CLUSTER` environment variable.

```go
var mesh mesh
vpc := ec2.NewVpc(this, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
	Vpc: Vpc,
	Name: jsii.String("domain.local"),
})
service := namespace.CreateService(jsii.String("Svc"))
node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			Port: jsii.Number(8081),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				HealthyThreshold: jsii.Number(3),
				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
				 // minimum
				Path: jsii.String("/health-check-path"),
				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
				 // minimum
				UnhealthyThreshold: jsii.Number(2),
			}),
		}),
	},
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
})
```

Create a `VirtualNode` with the constructor and add tags.

```go
var mesh mesh
var service service


node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			Port: jsii.Number(8080),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				HealthyThreshold: jsii.Number(3),
				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
				Path: jsii.String("/ping"),
				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
				UnhealthyThreshold: jsii.Number(2),
			}),
			Timeout: &HttpTimeout{
				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
			},
		}),
	},
	BackendDefaults: &BackendDefaults{
		TlsClientPolicy: &TlsClientPolicy{
			Validation: &TlsValidation{
				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
			},
		},
	},
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
})

cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
```

Create a `VirtualNode` with the customized access logging format.

```go
var mesh mesh
var service service

node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			Port: jsii.Number(8080),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				HealthyThreshold: jsii.Number(3),
				Interval: cdk.Duration_Seconds(jsii.Number(5)),
				Path: jsii.String("/ping"),
				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
				UnhealthyThreshold: jsii.Number(2),
			}),
			Timeout: &HttpTimeout{
				Idle: cdk.Duration_*Seconds(jsii.Number(5)),
			},
		}),
	},
	BackendDefaults: &BackendDefaults{
		TlsClientPolicy: &TlsClientPolicy{
			Validation: &TlsValidation{
				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
			},
		},
	},
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout"), appmesh.LoggingFormat_FromJson(map[string]*string{
		"testKey1": jsii.String("testValue1"),
		"testKey2": jsii.String("testValue2"),
	})),
})
```

By using a key-value pair indexed signature, you can specify json key pairs to customize the log entry pattern. You can also use text format as below. You can only specify one of these 2 formats.

```go
// Example automatically generated from non-compiling source. May contain errors.
accessLog: appmesh.AccessLog.fromFilePath('/dev/stdout', appmesh.LoggingFormat.fromText('test_pattern')),
```

For what values and operators you can use for these two formats, please visit the latest envoy documentation. (https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage)
Create a `VirtualNode` with the constructor and add backend virtual service.

```go
var mesh mesh
var router virtualRouter
var service service


node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			Port: jsii.Number(8080),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				HealthyThreshold: jsii.Number(3),
				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
				Path: jsii.String("/ping"),
				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
				UnhealthyThreshold: jsii.Number(2),
			}),
			Timeout: &HttpTimeout{
				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
			},
		}),
	},
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
})

virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &VirtualServiceProps{
	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualRouter(router),
	VirtualServiceName: jsii.String("service1.domain.local"),
})

node.AddBackend(appmesh.Backend_VirtualService(virtualService))
```

The `listeners` property can be left blank and added later with the `node.addListener()` method. The `serviceDiscovery` property must be specified when specifying a listener.

The `backends` property can be added with `node.addBackend()`. In the example, we define a virtual service and add it to the virtual node to allow egress traffic to other nodes.

The `backendDefaults` property is added to the node while creating the virtual node. These are the virtual node's default settings for all backends.

The `VirtualNode.addBackend()` method is especially useful if you want to create a circular traffic flow by having a Virtual Service as a backend whose provider is that same Virtual Node:

```go
var mesh mesh


node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
})

virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &VirtualServiceProps{
	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualNode(node),
	VirtualServiceName: jsii.String("service1.domain.local"),
})

node.AddBackend(appmesh.Backend_VirtualService(virtualService))
```

### Adding TLS to a listener

The `tls` property specifies TLS configuration when creating a listener for a virtual node or a virtual gateway.
Provide the TLS certificate to the proxy in one of the following ways:

* A certificate from AWS Certificate Manager (ACM).
* A customer-provided certificate (specify a `certificateChain` path file and a `privateKey` file path).
* A certificate provided by a Secrets Discovery Service (SDS) endpoint over local Unix Domain Socket (specify its `secretName`).

```go
// A Virtual Node with listener TLS from an ACM provided certificate
var cert certificate
var mesh mesh


node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
			Port: jsii.Number(80),
			Tls: &ListenerTlsOptions{
				Mode: appmesh.TlsMode_STRICT,
				Certificate: appmesh.TlsCertificate_Acm(cert),
			},
		}),
	},
})

// A Virtual Gateway with listener TLS from a customer provided file certificate
gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
	Mesh: Mesh,
	Listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
			Port: jsii.Number(8080),
			Tls: &ListenerTlsOptions{
				Mode: appmesh.TlsMode_STRICT,
				Certificate: appmesh.TlsCertificate_File(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
			},
		}),
	},
	VirtualGatewayName: jsii.String("gateway"),
})

// A Virtual Gateway with listener TLS from a SDS provided certificate
gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &VirtualGatewayProps{
	Mesh: Mesh,
	Listeners: []*virtualGatewayListener{
		appmesh.*virtualGatewayListener_Http2(&Http2GatewayListenerOptions{
			Port: jsii.Number(8080),
			Tls: &ListenerTlsOptions{
				Mode: appmesh.TlsMode_STRICT,
				Certificate: appmesh.TlsCertificate_Sds(jsii.String("secrete_certificate")),
			},
		}),
	},
	VirtualGatewayName: jsii.String("gateway2"),
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
var mesh mesh


node1 := appmesh.NewVirtualNode(this, jsii.String("node1"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
			Port: jsii.Number(80),
			Tls: &ListenerTlsOptions{
				Mode: appmesh.TlsMode_STRICT,
				Certificate: appmesh.TlsCertificate_File(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
				// Validate a file client certificates to enable mutual TLS authentication when a client provides a certificate.
				MutualTlsValidation: &MutualTlsValidation{
					Trust: appmesh.TlsValidationTrust_File(jsii.String("path-to-certificate")),
				},
			},
		}),
	},
})

certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
node2 := appmesh.NewVirtualNode(this, jsii.String("node2"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_*Dns(jsii.String("node2")),
	BackendDefaults: &BackendDefaults{
		TlsClientPolicy: &TlsClientPolicy{
			Ports: []*f64{
				jsii.Number(8080),
				jsii.Number(8081),
			},
			Validation: &TlsValidation{
				SubjectAlternativeNames: appmesh.SubjectAlternativeNames_MatchingExactly(jsii.String("mesh-endpoint.apps.local")),
				Trust: appmesh.TlsValidationTrust_Acm([]iCertificateAuthority{
					acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
				}),
			},
			// Provide a SDS client certificate when a server requests it and enable mutual TLS authentication.
			MutualTlsCertificate: appmesh.TlsCertificate_Sds(jsii.String("secret_certificate")),
		},
	},
})
```

### Adding outlier detection to a Virtual Node listener

The `outlierDetection` property adds outlier detection to a Virtual Node listener. The properties
`baseEjectionDuration`, `interval`, `maxEjectionPercent`, and `maxServerErrors` are required.

```go
var mesh mesh
// Cloud Map service discovery is currently required for host ejection by outlier detection
vpc := ec2.NewVpc(this, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
	Vpc: Vpc,
	Name: jsii.String("domain.local"),
})
service := namespace.CreateService(jsii.String("Svc"))
node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			OutlierDetection: &OutlierDetection{
				BaseEjectionDuration: awscdk.Duration_Seconds(jsii.Number(10)),
				Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
				MaxEjectionPercent: jsii.Number(50),
				MaxServerErrors: jsii.Number(5),
			},
		}),
	},
})
```

### Adding a connection pool to a listener

The `connectionPool` property can be added to a Virtual Node listener or Virtual Gateway listener to add a request connection pool. Each listener protocol type has its own connection pool properties.

```go
// A Virtual Node with a gRPC listener with a connection pool set
var mesh mesh

node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
	Mesh: Mesh,
	// DNS service discovery can optionally specify the DNS response type as either LOAD_BALANCER or ENDPOINTS.
	// LOAD_BALANCER means that the DNS resolver returns a loadbalanced set of endpoints,
	// whereas ENDPOINTS means that the DNS resolver is returning all the endpoints.
	// By default, the response type is assumed to be LOAD_BALANCER
	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node"), appmesh.DnsResponseType_ENDPOINTS),
	Listeners: []virtualNodeListener{
		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
			Port: jsii.Number(80),
			ConnectionPool: &HttpConnectionPool{
				MaxConnections: jsii.Number(100),
				MaxPendingRequests: jsii.Number(10),
			},
		}),
	},
})

// A Virtual Gateway with a gRPC listener with a connection pool set
gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
	Mesh: Mesh,
	Listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
			Port: jsii.Number(8080),
			ConnectionPool: &GrpcConnectionPool{
				MaxRequests: jsii.Number(10),
			},
		}),
	},
	VirtualGatewayName: jsii.String("gateway"),
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
mesh := appmesh.NewMesh(this, jsii.String("mesh"), &MeshProps{
	MeshName: jsii.String("mesh-with-preference"),
})

// Virtual Node with DNS service discovery and an IP preference
dnsNode := appmesh.NewVirtualNode(this, jsii.String("dns-node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("test"), appmesh.DnsResponseType_LOAD_BALANCER, appmesh.IpPreference_IPV4_ONLY),
})

// Virtual Node with CloudMap service discovery and an IP preference
vpc := ec2.NewVpc(this, jsii.String("vpc"))
namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
	Vpc: Vpc,
	Name: jsii.String("domain.local"),
})
service := namespace.CreateService(jsii.String("Svc"))

instanceAttribute := map[string]interface{}{
}
instanceAttribute.testKey = "testValue"

cloudmapNode := appmesh.NewVirtualNode(this, jsii.String("cloudmap-node"), &VirtualNodeProps{
	Mesh: Mesh,
	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service, instanceAttribute, appmesh.IpPreference_IPV4_ONLY),
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
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Http(&HttpRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &HttpRouteMatch{
			// Path that is passed to this method must start with '/'.
			Path: appmesh.HttpRoutePathMatch_StartsWith(jsii.String("/path-to-app")),
		},
	}),
})
```

Add an HTTP2 route that matches based on exact path, method, scheme, headers, and query parameters:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http2"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &HttpRouteMatch{
			Path: appmesh.HttpRoutePathMatch_Exactly(jsii.String("/exact")),
			Method: appmesh.HttpRouteMethod_POST,
			Protocol: appmesh.HttpRouteProtocol_HTTPS,
			Headers: []headerMatch{
				appmesh.*headerMatch_ValueIs(jsii.String("Content-Type"), jsii.String("application/json")),
				appmesh.*headerMatch_ValueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
			},
			QueryParameters: []queryParameterMatch{
				appmesh.*queryParameterMatch_ValueIs(jsii.String("query-field"), jsii.String("value")),
			},
		},
	}),
})
```

Add a single route with two targets and split traffic 50/50:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Http(&HttpRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
				Weight: jsii.Number(50),
			},
			&weightedTarget{
				VirtualNode: node,
				Weight: jsii.Number(50),
			},
		},
		Match: &HttpRouteMatch{
			Path: appmesh.HttpRoutePathMatch_StartsWith(jsii.String("/path-to-app")),
		},
	}),
})
```

Add an http2 route with retries:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http2-retry"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		RetryPolicy: &HttpRetryPolicy{
			// Retry if the connection failed
			TcpRetryEvents: []cONNECTION_ERROR{
				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
			},
			// Retry if HTTP responds with a gateway error (502, 503, 504)
			HttpRetryEvents: []httpRetryEvent{
				appmesh.*httpRetryEvent_GATEWAY_ERROR,
			},
			// Retry five times
			RetryAttempts: jsii.Number(5),
			// Use a 1 second timeout per retry
			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
		},
	}),
})
```

Add a gRPC route with retries:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-grpc-retry"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &GrpcRouteMatch{
			ServiceName: jsii.String("servicename"),
		},
		RetryPolicy: &GrpcRetryPolicy{
			TcpRetryEvents: []cONNECTION_ERROR{
				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
			},
			HttpRetryEvents: []httpRetryEvent{
				appmesh.*httpRetryEvent_GATEWAY_ERROR,
			},
			// Retry if gRPC responds that the request was cancelled, a resource
			// was exhausted, or if the service is unavailable
			GrpcRetryEvents: []grpcRetryEvent{
				appmesh.*grpcRetryEvent_CANCELLED,
				appmesh.*grpcRetryEvent_RESOURCE_EXHAUSTED,
				appmesh.*grpcRetryEvent_UNAVAILABLE,
			},
			RetryAttempts: jsii.Number(5),
			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
		},
	}),
})
```

Add an gRPC route that matches based on method name and metadata:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-grpc-retry"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &GrpcRouteMatch{
			// When method name is specified, service name must be also specified.
			MethodName: jsii.String("methodname"),
			ServiceName: jsii.String("servicename"),
			Metadata: []headerMatch{
				appmesh.*headerMatch_ValueStartsWith(jsii.String("Content-Type"), jsii.String("application/")),
				appmesh.*headerMatch_ValueDoesNotStartWith(jsii.String("Content-Type"), jsii.String("text/")),
			},
		},
	}),
})
```

Add a gRPC route that matches based on port:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-grpc-port"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &GrpcRouteMatch{
			Port: jsii.Number(1234),
		},
	}),
})
```

Add a gRPC route with timeout:

```go
var router virtualRouter
var node virtualNode


router.addRoute(jsii.String("route-http"), &RouteBaseProps{
	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
		WeightedTargets: []weightedTarget{
			&weightedTarget{
				VirtualNode: node,
			},
		},
		Match: &GrpcRouteMatch{
			ServiceName: jsii.String("my-service.default.svc.cluster.local"),
		},
		Timeout: &GrpcTimeout{
			Idle: awscdk.Duration_Seconds(jsii.Number(2)),
			PerRequest: awscdk.Duration_*Seconds(jsii.Number(1)),
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
var mesh mesh

certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"

gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
	Mesh: mesh,
	Listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener_Http(&HttpGatewayListenerOptions{
			Port: jsii.Number(443),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				Interval: awscdk.Duration_Seconds(jsii.Number(10)),
			}),
		}),
	},
	BackendDefaults: &BackendDefaults{
		TlsClientPolicy: &TlsClientPolicy{
			Ports: []*f64{
				jsii.Number(8080),
				jsii.Number(8081),
			},
			Validation: &TlsValidation{
				Trust: appmesh.TlsValidationTrust_Acm([]iCertificateAuthority{
					acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
				}),
			},
		},
	},
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
	VirtualGatewayName: jsii.String("virtualGateway"),
})
```

Add a virtual gateway directly to the mesh:

```go
var mesh mesh


gateway := mesh.addVirtualGateway(jsii.String("gateway"), &VirtualGatewayBaseProps{
	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
	VirtualGatewayName: jsii.String("virtualGateway"),
	Listeners: []virtualGatewayListener{
		appmesh.*virtualGatewayListener_Http(&HttpGatewayListenerOptions{
			Port: jsii.Number(443),
			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
				Interval: awscdk.Duration_Seconds(jsii.Number(10)),
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
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_Http(&HttpGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &HttpGatewayRouteMatch{
			Path: appmesh.HttpGatewayRoutePathMatch_Regex(jsii.String("regex")),
		},
	}),
})
```

For gRPC-based gateway routes, the `match` field can be used to match on service name, host name, port and metadata.

```go
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_Grpc(&GrpcGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &GrpcGatewayRouteMatch{
			Hostname: appmesh.GatewayRouteHostnameMatch_EndsWith(jsii.String(".example.com")),
		},
	}),
})
```

For HTTP based gateway routes, App Mesh automatically rewrites the matched prefix path in Gateway Route to “/”.
This automatic rewrite configuration can be overwritten in following ways:

```go
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_Http(&HttpGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &HttpGatewayRouteMatch{
			// This disables the default rewrite to '/', and retains original path.
			Path: appmesh.HttpGatewayRoutePathMatch_StartsWith(jsii.String("/path-to-app/"), jsii.String("")),
		},
	}),
})

gateway.addGatewayRoute(jsii.String("gateway-route-http-1"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_*Http(&HttpGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &HttpGatewayRouteMatch{
			// If the request full path is '/path-to-app/xxxxx', this rewrites the path to '/rewrittenUri/xxxxx'.
			// Please note both `prefixPathMatch` and `rewriteTo` must start and end with the `/` character.
			Path: appmesh.HttpGatewayRoutePathMatch_*StartsWith(jsii.String("/path-to-app/"), jsii.String("/rewrittenUri/")),
		},
	}),
})
```

If matching other path (exact or regex), only specific rewrite path can be specified.
Unlike `startsWith()` method above, no default rewrite is performed.

```go
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-http-2"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_Http(&HttpGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &HttpGatewayRouteMatch{
			// This rewrites the path from '/test' to '/rewrittenPath'.
			Path: appmesh.HttpGatewayRoutePathMatch_Exactly(jsii.String("/test"), jsii.String("/rewrittenPath")),
		},
	}),
})
```

For HTTP/gRPC based routes, App Mesh automatically rewrites
the original request received at the Virtual Gateway to the destination Virtual Service name.
This default host name rewrite can be configured by specifying the rewrite rule as one of the `match` property:

```go
var gateway virtualGateway
var virtualService virtualService


gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &GatewayRouteBaseProps{
	RouteSpec: appmesh.GatewayRouteSpec_Grpc(&GrpcGatewayRouteSpecOptions{
		RouteTarget: virtualService,
		Match: &GrpcGatewayRouteMatch{
			Hostname: appmesh.GatewayRouteHostnameMatch_Exactly(jsii.String("example.com")),
			// This disables the default rewrite to virtual service name and retain original request.
			RewriteRequestHostname: jsii.Boolean(false),
		},
	}),
})
```

## Importing Resources

Each App Mesh resource class comes with two static methods, `from<Resource>Arn` and `from<Resource>Attributes` (where `<Resource>` is replaced with the resource name, such as `VirtualNode`) for importing a reference to an existing App Mesh resource.
These imported resources can be used with other resources in your mesh as if they were defined directly in your CDK application.

```go
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh/virtualNode/testNode"
appmesh.VirtualNode_FromVirtualNodeArn(this, jsii.String("importedVirtualNode"), arn)
```

```go
virtualNodeName := "my-virtual-node"
appmesh.VirtualNode_FromVirtualNodeAttributes(this, jsii.String("imported-virtual-node"), &VirtualNodeAttributes{
	Mesh: appmesh.Mesh_FromMeshName(this, jsii.String("Mesh"), jsii.String("testMesh")),
	VirtualNodeName: virtualNodeName,
})
```

To import a mesh, again there are two static methods, `fromMeshArn` and `fromMeshName`.

```go
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh"
appmesh.Mesh_FromMeshArn(this, jsii.String("imported-mesh"), arn)
```

```go
appmesh.Mesh_FromMeshName(this, jsii.String("imported-mesh"), jsii.String("abc"))
```

## IAM Grants

`VirtualNode` and `VirtualGateway` provide `grantStreamAggregatedResources` methods that grant identities that are running
Envoy access to stream generated config from App Mesh.

```go
var mesh mesh

gateway := appmesh.NewVirtualGateway(this, jsii.String("testGateway"), &VirtualGatewayProps{
	Mesh: Mesh,
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
// This is the ARN for the mesh from different AWS IAM account ID.
// Ensure mesh is properly shared with your account. For more details, see: https://github.com/aws/aws-cdk/issues/15404
arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh"
sharedMesh := appmesh.Mesh_FromMeshArn(this, jsii.String("imported-mesh"), arn)

// This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
// This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
appmesh.NewVirtualNode(this, jsii.String("test-node"), &VirtualNodeProps{
	Mesh: sharedMesh,
})
```
