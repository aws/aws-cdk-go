package cloudassemblyschema

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Query to AMI context provider.
//
// TODO: EXAMPLE
//
// Experimental.
type AmiContextQuery struct {
	// Account to query.
	// Experimental.
	Account *string `json:"account"`
	// Filters to DescribeImages call.
	// Experimental.
	Filters *map[string]*[]*string `json:"filters"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Owners to DescribeImages call.
	// Experimental.
	Owners *[]*string `json:"owners"`
	// Region to query.
	// Experimental.
	Region *string `json:"region"`
}

// A manifest for a single artifact within the cloud assembly.
//
// TODO: EXAMPLE
//
// Experimental.
type ArtifactManifest struct {
	// IDs of artifacts that must be deployed before this artifact.
	// Experimental.
	Dependencies *[]*string `json:"dependencies"`
	// A string that represents this artifact.
	//
	// Should only be used in user interfaces.
	// Experimental.
	DisplayName *string `json:"displayName"`
	// The environment into which this artifact is deployed.
	// Experimental.
	Environment *string `json:"environment"`
	// Associated metadata.
	// Experimental.
	Metadata *map[string]*[]*MetadataEntry `json:"metadata"`
	// The set of properties for this artifact (depends on type).
	// Experimental.
	Properties interface{} `json:"properties"`
	// The type of artifact.
	// Experimental.
	Type ArtifactType `json:"type"`
}

// Type of artifact metadata entry.
// Experimental.
type ArtifactMetadataEntryType string

const (
	ArtifactMetadataEntryType_ASSET ArtifactMetadataEntryType = "ASSET"
	ArtifactMetadataEntryType_INFO ArtifactMetadataEntryType = "INFO"
	ArtifactMetadataEntryType_WARN ArtifactMetadataEntryType = "WARN"
	ArtifactMetadataEntryType_ERROR ArtifactMetadataEntryType = "ERROR"
	ArtifactMetadataEntryType_LOGICAL_ID ArtifactMetadataEntryType = "LOGICAL_ID"
	ArtifactMetadataEntryType_STACK_TAGS ArtifactMetadataEntryType = "STACK_TAGS"
)

// Type of cloud artifact.
// Experimental.
type ArtifactType string

const (
	ArtifactType_ASSET_MANIFEST ArtifactType = "ASSET_MANIFEST"
	ArtifactType_AWS_CLOUDFORMATION_STACK ArtifactType = "AWS_CLOUDFORMATION_STACK"
	ArtifactType_CDK_TREE ArtifactType = "CDK_TREE"
	ArtifactType_NESTED_CLOUD_ASSEMBLY ArtifactType = "NESTED_CLOUD_ASSEMBLY"
	ArtifactType_NONE ArtifactType = "NONE"
)

// A manifest which describes the cloud assembly.
//
// TODO: EXAMPLE
//
// Experimental.
type AssemblyManifest struct {
	// The set of artifacts in this assembly.
	// Experimental.
	Artifacts *map[string]*ArtifactManifest `json:"artifacts"`
	// Missing context information.
	//
	// If this field has values, it means that the
	// cloud assembly is not complete and should not be deployed.
	// Experimental.
	Missing *[]*MissingContext `json:"missing"`
	// Runtime information.
	// Experimental.
	Runtime *RuntimeInfo `json:"runtime"`
	// Protocol version.
	// Experimental.
	Version *string `json:"version"`
}

// Definitions for the asset manifest.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetManifest struct {
	// The Docker image assets in this manifest.
	// Experimental.
	DockerImages *map[string]*DockerImageAsset `json:"dockerImages"`
	// The file assets in this manifest.
	// Experimental.
	Files *map[string]*FileAsset `json:"files"`
	// Version of the manifest.
	// Experimental.
	Version *string `json:"version"`
}

// Artifact properties for the Asset Manifest.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetManifestProperties struct {
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	// Experimental.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter"`
	// Filename of the asset manifest.
	// Experimental.
	File *string `json:"file"`
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion"`
}

// Query to availability zone context provider.
//
// TODO: EXAMPLE
//
// Experimental.
type AvailabilityZonesContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
}

// Artifact properties for CloudFormation stacks.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsCloudFormationStackProperties struct {
	// The role that needs to be assumed to deploy the stack.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// Only used if `requiresBootstrapStackVersion` is set.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	// Experimental.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	// Experimental.
	CloudFormationExecutionRoleArn *string `json:"cloudFormationExecutionRoleArn"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	// Experimental.
	Parameters *map[string]*string `json:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion"`
	// The name to use for the CloudFormation stack.
	// Experimental.
	StackName *string `json:"stackName"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	// Experimental.
	StackTemplateAssetObjectUrl *string `json:"stackTemplateAssetObjectUrl"`
	// Values for CloudFormation stack tags that should be passed when the stack is deployed.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	// Experimental.
	TemplateFile *string `json:"templateFile"`
	// Whether to enable termination protection for this stack.
	// Experimental.
	TerminationProtection *bool `json:"terminationProtection"`
	// Whether this stack should be validated by the CLI after synthesis.
	// Experimental.
	ValidateOnSynth *bool `json:"validateOnSynth"`
}

// Destination for assets that need to be uploaded to AWS.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region"`
}

// Metadata Entry spec for container images.
//
// TODO: EXAMPLE
//
// Experimental.
type ContainerImageAssetMetadataEntry struct {
	// Build args to pass to the `docker build` command.
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `json:"file"`
	// Logical identifier for the asset.
	// Experimental.
	Id *string `json:"id"`
	// ECR Repository name and repo digest (separated by "@sha256:") where this image is stored.
	// Deprecated: specify `repositoryName` and `imageTag` instead, and then you
	// know where the image will go.
	ImageNameParameter *string `json:"imageNameParameter"`
	// The docker image tag to use for tagging pushed images.
	//
	// This field is
	// required if `imageParameterName` is ommited (otherwise, the app won't be
	// able to find the image).
	// Experimental.
	ImageTag *string `json:"imageTag"`
	// Type of asset.
	// Experimental.
	Packaging *string `json:"packaging"`
	// Path on disk to the asset.
	// Experimental.
	Path *string `json:"path"`
	// ECR repository name, if omitted a default name based on the asset's ID is used instead.
	//
	// Specify this property if you need to statically address the
	// image, e.g. from a Kubernetes Pod. Note, this is only the repository name,
	// without the registry and the tag parts.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
	// The hash of the asset source.
	// Experimental.
	SourceHash *string `json:"sourceHash"`
	// Docker target to build to.
	// Experimental.
	Target *string `json:"target"`
}

// Identifier for the context provider.
// Experimental.
type ContextProvider string

const (
	ContextProvider_AMI_PROVIDER ContextProvider = "AMI_PROVIDER"
	ContextProvider_AVAILABILITY_ZONE_PROVIDER ContextProvider = "AVAILABILITY_ZONE_PROVIDER"
	ContextProvider_ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER ContextProvider = "ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER"
	ContextProvider_HOSTED_ZONE_PROVIDER ContextProvider = "HOSTED_ZONE_PROVIDER"
	ContextProvider_KEY_PROVIDER ContextProvider = "KEY_PROVIDER"
	ContextProvider_LOAD_BALANCER_LISTENER_PROVIDER ContextProvider = "LOAD_BALANCER_LISTENER_PROVIDER"
	ContextProvider_LOAD_BALANCER_PROVIDER ContextProvider = "LOAD_BALANCER_PROVIDER"
	ContextProvider_SECURITY_GROUP_PROVIDER ContextProvider = "SECURITY_GROUP_PROVIDER"
	ContextProvider_SSM_PARAMETER_PROVIDER ContextProvider = "SSM_PARAMETER_PROVIDER"
	ContextProvider_VPC_PROVIDER ContextProvider = "VPC_PROVIDER"
)

// A file asset.
//
// TODO: EXAMPLE
//
// Experimental.
type DockerImageAsset struct {
	// Destinations for this file asset.
	// Experimental.
	Destinations *map[string]*DockerImageDestination `json:"destinations"`
	// Source description for file assets.
	// Experimental.
	Source *DockerImageSource `json:"source"`
}

// Where to publish docker images.
//
// TODO: EXAMPLE
//
// Experimental.
type DockerImageDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region"`
	// Tag of the image to publish.
	// Experimental.
	ImageTag *string `json:"imageTag"`
	// Name of the ECR repository to publish to.
	// Experimental.
	RepositoryName *string `json:"repositoryName"`
}

// Properties for how to produce a Docker image from a source.
//
// TODO: EXAMPLE
//
// Experimental.
type DockerImageSource struct {
	// The directory containing the Docker image build instructions.
	//
	// This path is relative to the asset manifest location.
	// Experimental.
	Directory *string `json:"directory"`
	// Additional build arguments.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildArgs *map[string]*string `json:"dockerBuildArgs"`
	// Target build stage in a Dockerfile with multiple build stages.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildTarget *string `json:"dockerBuildTarget"`
	// The name of the file with build instructions.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerFile *string `json:"dockerFile"`
	// A command-line executable that returns the name of a local Docker image on stdout after being run.
	// Experimental.
	Executable *[]*string `json:"executable"`
}

// Query to endpoint service context provider.
//
// TODO: EXAMPLE
//
// Experimental.
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
	// Query service name.
	// Experimental.
	ServiceName *string `json:"serviceName"`
}

// A file asset.
//
// TODO: EXAMPLE
//
// Experimental.
type FileAsset struct {
	// Destinations for this file asset.
	// Experimental.
	Destinations *map[string]*FileDestination `json:"destinations"`
	// Source description for file assets.
	// Experimental.
	Source *FileSource `json:"source"`
}

// Metadata Entry spec for files.
//
// TODO: EXAMPLE
//
// Experimental.
type FileAssetMetadataEntry struct {
	// The name of the parameter where the hash of the bundled asset should be passed in.
	// Experimental.
	ArtifactHashParameter *string `json:"artifactHashParameter"`
	// Logical identifier for the asset.
	// Experimental.
	Id *string `json:"id"`
	// Requested packaging style.
	// Experimental.
	Packaging *string `json:"packaging"`
	// Path on disk to the asset.
	// Experimental.
	Path *string `json:"path"`
	// Name of parameter where S3 bucket should be passed in.
	// Experimental.
	S3BucketParameter *string `json:"s3BucketParameter"`
	// Name of parameter where S3 key should be passed in.
	// Experimental.
	S3KeyParameter *string `json:"s3KeyParameter"`
	// The hash of the asset source.
	// Experimental.
	SourceHash *string `json:"sourceHash"`
}

// Packaging strategy for file assets.
// Experimental.
type FileAssetPackaging string

const (
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

// Where in S3 a file asset needs to be published.
//
// TODO: EXAMPLE
//
// Experimental.
type FileDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region"`
	// The name of the bucket.
	// Experimental.
	BucketName *string `json:"bucketName"`
	// The destination object key.
	// Experimental.
	ObjectKey *string `json:"objectKey"`
}

// Describe the source of a file asset.
//
// TODO: EXAMPLE
//
// Experimental.
type FileSource struct {
	// External command which will produce the file asset to upload.
	// Experimental.
	Executable *[]*string `json:"executable"`
	// Packaging method.
	//
	// Only allowed when `path` is specified.
	// Experimental.
	Packaging FileAssetPackaging `json:"packaging"`
	// The filesystem object to upload.
	//
	// This path is relative to the asset manifest location.
	// Experimental.
	Path *string `json:"path"`
}

// Query to hosted zone context provider.
//
// TODO: EXAMPLE
//
// Experimental.
type HostedZoneContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The domain name e.g. example.com to lookup.
	// Experimental.
	DomainName *string `json:"domainName"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// True if the zone you want to find is a private hosted zone.
	// Experimental.
	PrivateZone *bool `json:"privateZone"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
	// The VPC ID to that the private zone must be associated with.
	//
	// If you provide VPC ID and privateZone is false, this will return no results
	// and raise an error.
	// Experimental.
	VpcId *string `json:"vpcId"`
}

// Query input for looking up a KMS Key.
//
// TODO: EXAMPLE
//
// Experimental.
type KeyContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// Alias name used to search the Key.
	// Experimental.
	AliasName *string `json:"aliasName"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
}

// Query input for looking up a load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancerContextQuery struct {
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags"`
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType"`
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
}

// Filters for selecting load balancers.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancerFilter struct {
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags"`
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType"`
}

// Query input for looking up a load balancer listener.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancerListenerContextQuery struct {
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags"`
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType"`
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// Find by listener's arn.
	// Experimental.
	ListenerArn *string `json:"listenerArn"`
	// Filter listeners by listener port.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort"`
	// Filter by listener protocol.
	// Experimental.
	ListenerProtocol LoadBalancerListenerProtocol `json:"listenerProtocol"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
}

// The protocol for connections from clients to the load balancer.
// Experimental.
type LoadBalancerListenerProtocol string

const (
	LoadBalancerListenerProtocol_HTTP LoadBalancerListenerProtocol = "HTTP"
	LoadBalancerListenerProtocol_HTTPS LoadBalancerListenerProtocol = "HTTPS"
	LoadBalancerListenerProtocol_TCP LoadBalancerListenerProtocol = "TCP"
	LoadBalancerListenerProtocol_TCP_UDP LoadBalancerListenerProtocol = "TCP_UDP"
	LoadBalancerListenerProtocol_TLS LoadBalancerListenerProtocol = "TLS"
	LoadBalancerListenerProtocol_UDP LoadBalancerListenerProtocol = "UDP"
)

// Type of load balancer.
// Experimental.
type LoadBalancerType string

const (
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
)

// Protocol utility class.
// Experimental.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Deprecated.
// Deprecated: use `loadAssemblyManifest()`
func Manifest_Load(filePath *string) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"load",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the cloud assembly manifest from file.
// Experimental.
func Manifest_LoadAssemblyManifest(filePath *string) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"loadAssemblyManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the asset manifest from file.
// Experimental.
func Manifest_LoadAssetManifest(filePath *string) *AssetManifest {
	_init_.Initialize()

	var returns *AssetManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"loadAssetManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Deprecated.
// Deprecated: use `saveAssemblyManifest()`
func Manifest_Save(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"save",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the cloud assembly manifest to file.
// Experimental.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
// Experimental.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssetManifest",
		[]interface{}{manifest, filePath},
	)
}

// Fetch the current schema version number.
// Experimental.
func Manifest_Version() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A metadata entry in a cloud assembly artifact.
//
// TODO: EXAMPLE
//
// Experimental.
type MetadataEntry struct {
	// The data.
	// Experimental.
	Data interface{} `json:"data"`
	// A stack trace for when the entry was created.
	// Experimental.
	Trace *[]*string `json:"trace"`
	// The type of the metadata entry.
	// Experimental.
	Type *string `json:"type"`
}

// Represents a missing piece of context.
//
// TODO: EXAMPLE
//
// Experimental.
type MissingContext struct {
	// The missing context key.
	// Experimental.
	Key *string `json:"key"`
	// A set of provider-specific options.
	// Experimental.
	Props interface{} `json:"props"`
	// The provider from which we expect this context key to be obtained.
	// Experimental.
	Provider ContextProvider `json:"provider"`
}

// Artifact properties for nested cloud assemblies.
//
// TODO: EXAMPLE
//
// Experimental.
type NestedCloudAssemblyProperties struct {
	// Relative path to the nested cloud assembly.
	// Experimental.
	DirectoryName *string `json:"directoryName"`
	// Display name for the cloud assembly.
	// Experimental.
	DisplayName *string `json:"displayName"`
}

// Information about the application's runtime components.
//
// TODO: EXAMPLE
//
// Experimental.
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	// Experimental.
	Libraries *map[string]*string `json:"libraries"`
}

// Query to SSM Parameter Context Provider.
//
// TODO: EXAMPLE
//
// Experimental.
type SSMParameterContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Parameter name to query.
	// Experimental.
	ParameterName *string `json:"parameterName"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
}

// Query input for looking up a security group.
//
// TODO: EXAMPLE
//
// Experimental.
type SecurityGroupContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
	// Security group id.
	// Experimental.
	SecurityGroupId *string `json:"securityGroupId"`
	// Security group name.
	// Experimental.
	SecurityGroupName *string `json:"securityGroupName"`
	// VPC ID.
	// Experimental.
	VpcId *string `json:"vpcId"`
}

// Metadata Entry spec for stack tag.
//
// TODO: EXAMPLE
//
// Experimental.
type Tag struct {
	// Tag key.
	//
	// (In the actual file on disk this will be cased as "Key", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	// Experimental.
	Key *string `json:"key"`
	// Tag value.
	//
	// (In the actual file on disk this will be cased as "Value", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	// Experimental.
	Value *string `json:"value"`
}

// Artifact properties for the Construct Tree Artifact.
//
// TODO: EXAMPLE
//
// Experimental.
type TreeArtifactProperties struct {
	// Filename of the tree artifact.
	// Experimental.
	File *string `json:"file"`
}

// Query input for looking up a VPC.
//
// TODO: EXAMPLE
//
// Experimental.
type VpcContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account"`
	// Filters to apply to the VPC.
	//
	// Filter parameters are the same as passed to DescribeVpcs.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html
	//
	// Experimental.
	Filter *map[string]*string `json:"filter"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn"`
	// Query region.
	// Experimental.
	Region *string `json:"region"`
	// Whether to populate the subnetGroups field of the {@link VpcContextResponse}, which contains potentially asymmetric subnet groups.
	// Experimental.
	ReturnAsymmetricSubnets *bool `json:"returnAsymmetricSubnets"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	// Experimental.
	SubnetGroupNameTag *string `json:"subnetGroupNameTag"`
}

