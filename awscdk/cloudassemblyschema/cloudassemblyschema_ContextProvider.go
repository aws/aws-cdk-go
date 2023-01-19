package cloudassemblyschema


// Identifier for the context provider.
// Experimental.
type ContextProvider string

const (
	// AMI provider.
	// Experimental.
	ContextProvider_AMI_PROVIDER ContextProvider = "AMI_PROVIDER"
	// AZ provider.
	// Experimental.
	ContextProvider_AVAILABILITY_ZONE_PROVIDER ContextProvider = "AVAILABILITY_ZONE_PROVIDER"
	// Route53 Hosted Zone provider.
	// Experimental.
	ContextProvider_HOSTED_ZONE_PROVIDER ContextProvider = "HOSTED_ZONE_PROVIDER"
	// SSM Parameter Provider.
	// Experimental.
	ContextProvider_SSM_PARAMETER_PROVIDER ContextProvider = "SSM_PARAMETER_PROVIDER"
	// VPC Provider.
	// Experimental.
	ContextProvider_VPC_PROVIDER ContextProvider = "VPC_PROVIDER"
	// VPC Endpoint Service AZ Provider.
	// Experimental.
	ContextProvider_ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER ContextProvider = "ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER"
	// Load balancer provider.
	// Experimental.
	ContextProvider_LOAD_BALANCER_PROVIDER ContextProvider = "LOAD_BALANCER_PROVIDER"
	// Load balancer listener provider.
	// Experimental.
	ContextProvider_LOAD_BALANCER_LISTENER_PROVIDER ContextProvider = "LOAD_BALANCER_LISTENER_PROVIDER"
	// Security group provider.
	// Experimental.
	ContextProvider_SECURITY_GROUP_PROVIDER ContextProvider = "SECURITY_GROUP_PROVIDER"
	// KMS Key Provider.
	// Experimental.
	ContextProvider_KEY_PROVIDER ContextProvider = "KEY_PROVIDER"
	// A plugin provider (the actual plugin name will be in the properties).
	// Experimental.
	ContextProvider_PLUGIN ContextProvider = "PLUGIN"
)

