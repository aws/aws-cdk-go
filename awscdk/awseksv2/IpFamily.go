package awseksv2


// EKS cluster IP family.
type IpFamily string

const (
	// Use IPv4 for pods and services in your cluster.
	IpFamily_IP_V4 IpFamily = "IP_V4"
	// Use IPv6 for pods and services in your cluster.
	IpFamily_IP_V6 IpFamily = "IP_V6"
)

