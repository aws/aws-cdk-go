package cloudassemblyschema


// Identifier for the context provider.
type ContextProvider string

const (
	// AMI provider.
	ContextProvider_AMI_PROVIDER ContextProvider = "AMI_PROVIDER"
	// AZ provider.
	ContextProvider_AVAILABILITY_ZONE_PROVIDER ContextProvider = "AVAILABILITY_ZONE_PROVIDER"
	// Route53 Hosted Zone provider.
	ContextProvider_HOSTED_ZONE_PROVIDER ContextProvider = "HOSTED_ZONE_PROVIDER"
	// SSM Parameter Provider.
	ContextProvider_SSM_PARAMETER_PROVIDER ContextProvider = "SSM_PARAMETER_PROVIDER"
	// VPC Provider.
	ContextProvider_VPC_PROVIDER ContextProvider = "VPC_PROVIDER"
	// VPC Endpoint Service AZ Provider.
	ContextProvider_ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER ContextProvider = "ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER"
	// Load balancer provider.
	ContextProvider_LOAD_BALANCER_PROVIDER ContextProvider = "LOAD_BALANCER_PROVIDER"
	// Load balancer listener provider.
	ContextProvider_LOAD_BALANCER_LISTENER_PROVIDER ContextProvider = "LOAD_BALANCER_LISTENER_PROVIDER"
	// Security group provider.
	ContextProvider_SECURITY_GROUP_PROVIDER ContextProvider = "SECURITY_GROUP_PROVIDER"
	// KMS Key Provider.
	ContextProvider_KEY_PROVIDER ContextProvider = "KEY_PROVIDER"
	// A plugin provider (the actual plugin name will be in the properties).
	ContextProvider_PLUGIN ContextProvider = "PLUGIN"
)

