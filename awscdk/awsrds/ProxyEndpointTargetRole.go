package awsrds


// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
//
// Example:
//   var vpc Vpc
//   var secrets []Secret
//   var dbInstance DatabaseInstance
//
//
//   proxy := dbInstance.AddProxy(jsii.String("Proxy"), &DatabaseProxyOptions{
//   	Secrets: Secrets,
//   	Vpc: Vpc,
//   })
//
//   // Add a reader endpoint
//   proxy.AddEndpoint(jsii.String("ProxyEndpoint"), &DatabaseProxyEndpointOptions{
//   	Vpc: Vpc,
//   	TargetRole: rds.ProxyEndpointTargetRole_READ_ONLY,
//   })
//
type ProxyEndpointTargetRole string

const (
	// The proxy endpoint can be used for both read and write operations.
	ProxyEndpointTargetRole_READ_WRITE ProxyEndpointTargetRole = "READ_WRITE"
	// The proxy endpoint can be used only for read operations.
	ProxyEndpointTargetRole_READ_ONLY ProxyEndpointTargetRole = "READ_ONLY"
)

