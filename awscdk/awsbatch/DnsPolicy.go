package awsbatch


// The DNS Policy for the pod used by the Job Definition.
// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy
//
type DnsPolicy string

const (
	// The Pod inherits the name resolution configuration from the node that the Pods run on.
	DnsPolicy_DEFAULT DnsPolicy = "DEFAULT"
	// Any DNS query that does not match the configured cluster domain suffix, such as `"www.kubernetes.io"`, is forwarded to an upstream nameserver by the DNS server. Cluster administrators may have extra stub-domain and upstream DNS servers configured.
	DnsPolicy_CLUSTER_FIRST DnsPolicy = "CLUSTER_FIRST"
	// For Pods running with `hostNetwork`, you should explicitly set its DNS policy to `CLUSTER_FIRST_WITH_HOST_NET`.
	//
	// Otherwise, Pods running with `hostNetwork` and `CLUSTER_FIRST` will fallback to the behavior of the `DEFAULT` policy.
	DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET DnsPolicy = "CLUSTER_FIRST_WITH_HOST_NET"
)

