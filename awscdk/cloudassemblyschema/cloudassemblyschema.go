package cloudassemblyschema

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Query to AMI context provider.
//
// TODO: EXAMPLE
//
type AmiContextQuery struct {
	// Account to query.
	Account *string `json:"account" yaml:"account"`
	// Filters to DescribeImages call.
	Filters *map[string]*[]*string `json:"filters" yaml:"filters"`
	// Region to query.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Owners to DescribeImages call.
	Owners *[]*string `json:"owners" yaml:"owners"`
}

// A manifest for a single artifact within the cloud assembly.
//
// TODO: EXAMPLE
//
type ArtifactManifest struct {
	// The type of artifact.
	Type ArtifactType `json:"type" yaml:"type"`
	// IDs of artifacts that must be deployed before this artifact.
	Dependencies *[]*string `json:"dependencies" yaml:"dependencies"`
	// A string that represents this artifact.
	//
	// Should only be used in user interfaces.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// The environment into which this artifact is deployed.
	Environment *string `json:"environment" yaml:"environment"`
	// Associated metadata.
	Metadata *map[string]*[]*MetadataEntry `json:"metadata" yaml:"metadata"`
	// The set of properties for this artifact (depends on type).
	Properties interface{} `json:"properties" yaml:"properties"`
}

// Type of artifact metadata entry.
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
type ArtifactType string

const (
	ArtifactType_NONE ArtifactType = "NONE"
	ArtifactType_AWS_CLOUDFORMATION_STACK ArtifactType = "AWS_CLOUDFORMATION_STACK"
	ArtifactType_CDK_TREE ArtifactType = "CDK_TREE"
	ArtifactType_ASSET_MANIFEST ArtifactType = "ASSET_MANIFEST"
	ArtifactType_NESTED_CLOUD_ASSEMBLY ArtifactType = "NESTED_CLOUD_ASSEMBLY"
)

// A manifest which describes the cloud assembly.
//
// TODO: EXAMPLE
//
type AssemblyManifest struct {
	// Protocol version.
	Version *string `json:"version" yaml:"version"`
	// The set of artifacts in this assembly.
	Artifacts *map[string]*ArtifactManifest `json:"artifacts" yaml:"artifacts"`
	// Missing context information.
	//
	// If this field has values, it means that the
	// cloud assembly is not complete and should not be deployed.
	Missing *[]*MissingContext `json:"missing" yaml:"missing"`
	// Runtime information.
	Runtime *RuntimeInfo `json:"runtime" yaml:"runtime"`
}

// Definitions for the asset manifest.
//
// TODO: EXAMPLE
//
type AssetManifest struct {
	// Version of the manifest.
	Version *string `json:"version" yaml:"version"`
	// The Docker image assets in this manifest.
	DockerImages *map[string]*DockerImageAsset `json:"dockerImages" yaml:"dockerImages"`
	// The file assets in this manifest.
	Files *map[string]*FileAsset `json:"files" yaml:"files"`
}

// Artifact properties for the Asset Manifest.
//
// TODO: EXAMPLE
//
type AssetManifestProperties struct {
	// Filename of the asset manifest.
	File *string `json:"file" yaml:"file"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

// Query to availability zone context provider.
//
// TODO: EXAMPLE
//
type AvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Artifact properties for CloudFormation stacks.
//
// TODO: EXAMPLE
//
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	TemplateFile *string `json:"templateFile" yaml:"templateFile"`
	// The role that needs to be assumed to deploy the stack.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// Only used if `requiresBootstrapStackVersion` is set.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	CloudFormationExecutionRoleArn *string `json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	LookupRole *BootstrapRole `json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// The name to use for the CloudFormation stack.
	StackName *string `json:"stackName" yaml:"stackName"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	StackTemplateAssetObjectUrl *string `json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
	// Values for CloudFormation stack tags that should be passed when the stack is deployed.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `json:"terminationProtection" yaml:"terminationProtection"`
	// Whether this stack should be validated by the CLI after synthesis.
	ValidateOnSynth *bool `json:"validateOnSynth" yaml:"validateOnSynth"`
}

// Destination for assets that need to be uploaded to AWS.
//
// TODO: EXAMPLE
//
type AwsDestination struct {
	// The role that needs to be assumed while publishing this asset.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	Region *string `json:"region" yaml:"region"`
}

// Information needed to access an IAM role created as part of the bootstrap process.
//
// TODO: EXAMPLE
//
type BootstrapRole struct {
	// The ARN of the IAM role created as part of bootrapping e.g. lookupRoleArn.
	Arn *string `json:"arn" yaml:"arn"`
	// External ID to use when assuming the bootstrap role.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// Name of SSM parameter with bootstrap stack version.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to use this role.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

// Metadata Entry spec for container images.
//
// TODO: EXAMPLE
//
type ContainerImageAssetMetadataEntry struct {
	// Logical identifier for the asset.
	Id *string `json:"id" yaml:"id"`
	// Type of asset.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	Path *string `json:"path" yaml:"path"`
	// The hash of the asset source.
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
	// Build args to pass to the `docker build` command.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file" yaml:"file"`
	// The docker image tag to use for tagging pushed images.
	//
	// This field is
	// required if `imageParameterName` is ommited (otherwise, the app won't be
	// able to find the image).
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
	// Networking mode for the RUN commands during build.
	NetworkMode *string `json:"networkMode" yaml:"networkMode"`
	// ECR repository name, if omitted a default name based on the asset's ID is used instead.
	//
	// Specify this property if you need to statically address the
	// image, e.g. from a Kubernetes Pod. Note, this is only the repository name,
	// without the registry and the tag parts.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	Target *string `json:"target" yaml:"target"`
}

// Identifier for the context provider.
type ContextProvider string

const (
	ContextProvider_AMI_PROVIDER ContextProvider = "AMI_PROVIDER"
	ContextProvider_AVAILABILITY_ZONE_PROVIDER ContextProvider = "AVAILABILITY_ZONE_PROVIDER"
	ContextProvider_HOSTED_ZONE_PROVIDER ContextProvider = "HOSTED_ZONE_PROVIDER"
	ContextProvider_SSM_PARAMETER_PROVIDER ContextProvider = "SSM_PARAMETER_PROVIDER"
	ContextProvider_VPC_PROVIDER ContextProvider = "VPC_PROVIDER"
	ContextProvider_ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER ContextProvider = "ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER"
	ContextProvider_LOAD_BALANCER_PROVIDER ContextProvider = "LOAD_BALANCER_PROVIDER"
	ContextProvider_LOAD_BALANCER_LISTENER_PROVIDER ContextProvider = "LOAD_BALANCER_LISTENER_PROVIDER"
	ContextProvider_SECURITY_GROUP_PROVIDER ContextProvider = "SECURITY_GROUP_PROVIDER"
	ContextProvider_KEY_PROVIDER ContextProvider = "KEY_PROVIDER"
	ContextProvider_PLUGIN ContextProvider = "PLUGIN"
)

// A file asset.
//
// TODO: EXAMPLE
//
type DockerImageAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*DockerImageDestination `json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *DockerImageSource `json:"source" yaml:"source"`
}

// Where to publish docker images.
//
// TODO: EXAMPLE
//
type DockerImageDestination struct {
	// The role that needs to be assumed while publishing this asset.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	Region *string `json:"region" yaml:"region"`
	// Tag of the image to publish.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
	// Name of the ECR repository to publish to.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// Properties for how to produce a Docker image from a source.
//
// TODO: EXAMPLE
//
type DockerImageSource struct {
	// The directory containing the Docker image build instructions.
	//
	// This path is relative to the asset manifest location.
	Directory *string `json:"directory" yaml:"directory"`
	// Additional build arguments.
	//
	// Only allowed when `directory` is set.
	DockerBuildArgs *map[string]*string `json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Target build stage in a Dockerfile with multiple build stages.
	//
	// Only allowed when `directory` is set.
	DockerBuildTarget *string `json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// The name of the file with build instructions.
	//
	// Only allowed when `directory` is set.
	DockerFile *string `json:"dockerFile" yaml:"dockerFile"`
	// A command-line executable that returns the name of a local Docker image on stdout after being run.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	NetworkMode *string `json:"networkMode" yaml:"networkMode"`
}

// Query to endpoint service context provider.
//
// TODO: EXAMPLE
//
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// Query service name.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// A file asset.
//
// TODO: EXAMPLE
//
type FileAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*FileDestination `json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *FileSource `json:"source" yaml:"source"`
}

// Metadata Entry spec for files.
//
// TODO: EXAMPLE
//
type FileAssetMetadataEntry struct {
	// The name of the parameter where the hash of the bundled asset should be passed in.
	ArtifactHashParameter *string `json:"artifactHashParameter" yaml:"artifactHashParameter"`
	// Logical identifier for the asset.
	Id *string `json:"id" yaml:"id"`
	// Requested packaging style.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	Path *string `json:"path" yaml:"path"`
	// Name of parameter where S3 bucket should be passed in.
	S3BucketParameter *string `json:"s3BucketParameter" yaml:"s3BucketParameter"`
	// Name of parameter where S3 key should be passed in.
	S3KeyParameter *string `json:"s3KeyParameter" yaml:"s3KeyParameter"`
	// The hash of the asset source.
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
}

// Packaging strategy for file assets.
type FileAssetPackaging string

const (
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

// Where in S3 a file asset needs to be published.
//
// TODO: EXAMPLE
//
type FileDestination struct {
	// The role that needs to be assumed while publishing this asset.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	Region *string `json:"region" yaml:"region"`
	// The name of the bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The destination object key.
	ObjectKey *string `json:"objectKey" yaml:"objectKey"`
}

// Describe the source of a file asset.
//
// TODO: EXAMPLE
//
type FileSource struct {
	// External command which will produce the file asset to upload.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// Packaging method.
	//
	// Only allowed when `path` is specified.
	Packaging FileAssetPackaging `json:"packaging" yaml:"packaging"`
	// The filesystem object to upload.
	//
	// This path is relative to the asset manifest location.
	Path *string `json:"path" yaml:"path"`
}

// Query to hosted zone context provider.
//
// TODO: EXAMPLE
//
type HostedZoneContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// The domain name e.g. example.com to lookup.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// True if the zone you want to find is a private hosted zone.
	PrivateZone *bool `json:"privateZone" yaml:"privateZone"`
	// The VPC ID to that the private zone must be associated with.
	//
	// If you provide VPC ID and privateZone is false, this will return no results
	// and raise an error.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Query input for looking up a KMS Key.
//
// TODO: EXAMPLE
//
type KeyContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Alias name used to search the Key.
	AliasName *string `json:"aliasName" yaml:"aliasName"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Query input for looking up a load balancer.
//
// TODO: EXAMPLE
//
type LoadBalancerContextQuery struct {
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Filters for selecting load balancers.
//
// TODO: EXAMPLE
//
type LoadBalancerFilter struct {
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Query input for looking up a load balancer listener.
//
// TODO: EXAMPLE
//
type LoadBalancerListenerContextQuery struct {
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// Find by listener's arn.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener port.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// Filter by listener protocol.
	ListenerProtocol LoadBalancerListenerProtocol `json:"listenerProtocol" yaml:"listenerProtocol"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// The protocol for connections from clients to the load balancer.
type LoadBalancerListenerProtocol string

const (
	LoadBalancerListenerProtocol_HTTP LoadBalancerListenerProtocol = "HTTP"
	LoadBalancerListenerProtocol_HTTPS LoadBalancerListenerProtocol = "HTTPS"
	LoadBalancerListenerProtocol_TCP LoadBalancerListenerProtocol = "TCP"
	LoadBalancerListenerProtocol_TLS LoadBalancerListenerProtocol = "TLS"
	LoadBalancerListenerProtocol_UDP LoadBalancerListenerProtocol = "UDP"
	LoadBalancerListenerProtocol_TCP_UDP LoadBalancerListenerProtocol = "TCP_UDP"
)

// Type of load balancer.
type LoadBalancerType string

const (
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
)

// Protocol utility class.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Load and validates the cloud assembly manifest from file.
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

// Validates and saves the cloud assembly manifest to file.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssetManifest",
		[]interface{}{manifest, filePath},
	)
}

// Fetch the current schema version number.
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
type MetadataEntry struct {
	// The type of the metadata entry.
	Type *string `json:"type" yaml:"type"`
	// The data.
	Data interface{} `json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	Trace *[]*string `json:"trace" yaml:"trace"`
}

// Represents a missing piece of context.
//
// TODO: EXAMPLE
//
type MissingContext struct {
	// The missing context key.
	Key *string `json:"key" yaml:"key"`
	// A set of provider-specific options.
	Props interface{} `json:"props" yaml:"props"`
	// The provider from which we expect this context key to be obtained.
	Provider ContextProvider `json:"provider" yaml:"provider"`
}

// Artifact properties for nested cloud assemblies.
//
// TODO: EXAMPLE
//
type NestedCloudAssemblyProperties struct {
	// Relative path to the nested cloud assembly.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Display name for the cloud assembly.
	DisplayName *string `json:"displayName" yaml:"displayName"`
}

// Query input for plugins.
//
// This alternate branch is necessary because it needs to be able to escape all type checking
// we do on on the cloud assembly -- we cannot know the properties that will be used a priori.
//
// TODO: EXAMPLE
//
type PluginContextQuery struct {
	// The name of the plugin.
	PluginName *string `json:"pluginName" yaml:"pluginName"`
}

// Information about the application's runtime components.
//
// TODO: EXAMPLE
//
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	Libraries *map[string]*string `json:"libraries" yaml:"libraries"`
}

// Query to SSM Parameter Context Provider.
//
// TODO: EXAMPLE
//
type SSMParameterContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Parameter name to query.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Query input for looking up a security group.
//
// TODO: EXAMPLE
//
type SecurityGroupContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Security group id.
	SecurityGroupId *string `json:"securityGroupId" yaml:"securityGroupId"`
	// Security group name.
	SecurityGroupName *string `json:"securityGroupName" yaml:"securityGroupName"`
	// VPC ID.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Metadata Entry spec for stack tag.
//
// TODO: EXAMPLE
//
type Tag struct {
	// Tag key.
	//
	// (In the actual file on disk this will be cased as "Key", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	Key *string `json:"key" yaml:"key"`
	// Tag value.
	//
	// (In the actual file on disk this will be cased as "Value", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	Value *string `json:"value" yaml:"value"`
}

// Artifact properties for the Construct Tree Artifact.
//
// TODO: EXAMPLE
//
type TreeArtifactProperties struct {
	// Filename of the tree artifact.
	File *string `json:"file" yaml:"file"`
}

// Query input for looking up a VPC.
//
// TODO: EXAMPLE
//
type VpcContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Filters to apply to the VPC.
	//
	// Filter parameters are the same as passed to DescribeVpcs.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html
	//
	Filter *map[string]*string `json:"filter" yaml:"filter"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Whether to populate the subnetGroups field of the {@link VpcContextResponse}, which contains potentially asymmetric subnet groups.
	ReturnAsymmetricSubnets *bool `json:"returnAsymmetricSubnets" yaml:"returnAsymmetricSubnets"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	SubnetGroupNameTag *string `json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
}

