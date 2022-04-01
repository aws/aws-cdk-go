package cloudassemblyschema

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Query to AMI context provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   amiContextQuery := &amiContextQuery{
//   	account: jsii.String("account"),
//   	filters: map[string][]*string{
//   		"filtersKey": []*string{
//   			jsii.String("filters"),
//   		},
//   	},
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	owners: []*string{
//   		jsii.String("owners"),
//   	},
//   }
//
// Experimental.
type AmiContextQuery struct {
	// Account to query.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Filters to DescribeImages call.
	// Experimental.
	Filters *map[string]*[]*string `json:"filters" yaml:"filters"`
	// Region to query.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Owners to DescribeImages call.
	// Experimental.
	Owners *[]*string `json:"owners" yaml:"owners"`
}

// A manifest for a single artifact within the cloud assembly.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   artifactManifest := &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
//
//   	// the properties below are optional
//   	dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	displayName: jsii.String("displayName"),
//   	environment: jsii.String("environment"),
//   	metadata: map[string][]metadataEntry{
//   		"metadataKey": []*metadataEntry{
//   			&metadataEntry{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"data": jsii.String("data"),
//   				"trace": []*string{
//   					jsii.String("trace"),
//   				},
//   			},
//   		},
//   	},
//   	properties: &awsCloudFormationStackProperties{
//   		templateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		assumeRoleArn: jsii.String("assumeRoleArn"),
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		lookupRole: &bootstrapRole{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			requiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		requiresBootstrapStackVersion: jsii.Number(123),
//   		stackName: jsii.String("stackName"),
//   		stackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		terminationProtection: jsii.Boolean(false),
//   		validateOnSynth: jsii.Boolean(false),
//   	},
//   }
//
// Experimental.
type ArtifactManifest struct {
	// The type of artifact.
	// Experimental.
	Type ArtifactType `json:"type" yaml:"type"`
	// IDs of artifacts that must be deployed before this artifact.
	// Experimental.
	Dependencies *[]*string `json:"dependencies" yaml:"dependencies"`
	// A string that represents this artifact.
	//
	// Should only be used in user interfaces.
	// Experimental.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// The environment into which this artifact is deployed.
	// Experimental.
	Environment *string `json:"environment" yaml:"environment"`
	// Associated metadata.
	// Experimental.
	Metadata *map[string]*[]*MetadataEntry `json:"metadata" yaml:"metadata"`
	// The set of properties for this artifact (depends on type).
	// Experimental.
	Properties interface{} `json:"properties" yaml:"properties"`
}

// Type of artifact metadata entry.
// Experimental.
type ArtifactMetadataEntryType string

const (
	// Asset in metadata.
	// Experimental.
	ArtifactMetadataEntryType_ASSET ArtifactMetadataEntryType = "ASSET"
	// Metadata key used to print INFO-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_INFO ArtifactMetadataEntryType = "INFO"
	// Metadata key used to print WARNING-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_WARN ArtifactMetadataEntryType = "WARN"
	// Metadata key used to print ERROR-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_ERROR ArtifactMetadataEntryType = "ERROR"
	// Represents the CloudFormation logical ID of a resource at a certain path.
	// Experimental.
	ArtifactMetadataEntryType_LOGICAL_ID ArtifactMetadataEntryType = "LOGICAL_ID"
	// Represents tags of a stack.
	// Experimental.
	ArtifactMetadataEntryType_STACK_TAGS ArtifactMetadataEntryType = "STACK_TAGS"
)

// Type of cloud artifact.
// Experimental.
type ArtifactType string

const (
	// Stub required because of JSII.
	// Experimental.
	ArtifactType_NONE ArtifactType = "NONE"
	// The artifact is an AWS CloudFormation stack.
	// Experimental.
	ArtifactType_AWS_CLOUDFORMATION_STACK ArtifactType = "AWS_CLOUDFORMATION_STACK"
	// The artifact contains the CDK application's construct tree.
	// Experimental.
	ArtifactType_CDK_TREE ArtifactType = "CDK_TREE"
	// Manifest for all assets in the Cloud Assembly.
	// Experimental.
	ArtifactType_ASSET_MANIFEST ArtifactType = "ASSET_MANIFEST"
	// Nested Cloud Assembly.
	// Experimental.
	ArtifactType_NESTED_CLOUD_ASSEMBLY ArtifactType = "NESTED_CLOUD_ASSEMBLY"
)

// A manifest which describes the cloud assembly.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   assemblyManifest := &assemblyManifest{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	artifacts: map[string]artifactManifest{
//   		"artifactsKey": &artifactManifest{
//   			"type": cloud_assembly_schema.ArtifactType_NONE,
//
//   			// the properties below are optional
//   			"dependencies": []*string{
//   				jsii.String("dependencies"),
//   			},
//   			"displayName": jsii.String("displayName"),
//   			"environment": jsii.String("environment"),
//   			"metadata": map[string][]MetadataEntry{
//   				"metadataKey": []MetadataEntry{
//   					&MetadataEntry{
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"data": jsii.String("data"),
//   						"trace": []*string{
//   							jsii.String("trace"),
//   						},
//   					},
//   				},
//   			},
//   			"properties": &AwsCloudFormationStackProperties{
//   				"templateFile": jsii.String("templateFile"),
//
//   				// the properties below are optional
//   				"assumeRoleArn": jsii.String("assumeRoleArn"),
//   				"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   				"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   				"cloudFormationExecutionRoleArn": jsii.String("cloudFormationExecutionRoleArn"),
//   				"lookupRole": &BootstrapRole{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   					"requiresBootstrapStackVersion": jsii.Number(123),
//   				},
//   				"parameters": map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				"requiresBootstrapStackVersion": jsii.Number(123),
//   				"stackName": jsii.String("stackName"),
//   				"stackTemplateAssetObjectUrl": jsii.String("stackTemplateAssetObjectUrl"),
//   				"tags": map[string]*string{
//   					"tagsKey": jsii.String("tags"),
//   				},
//   				"terminationProtection": jsii.Boolean(false),
//   				"validateOnSynth": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	missing: []missingContext{
//   		&missingContext{
//   			key: jsii.String("key"),
//   			props: &amiContextQuery{
//   				account: jsii.String("account"),
//   				filters: map[string][]*string{
//   					"filtersKey": []*string{
//   						jsii.String("filters"),
//   					},
//   				},
//   				region: jsii.String("region"),
//
//   				// the properties below are optional
//   				lookupRoleArn: jsii.String("lookupRoleArn"),
//   				owners: []interface{}{
//   					jsii.String("owners"),
//   				},
//   			},
//   			provider: cloud_assembly_schema.contextProvider_AMI_PROVIDER,
//   		},
//   	},
//   	runtime: &runtimeInfo{
//   		libraries: map[string]*string{
//   			"librariesKey": jsii.String("libraries"),
//   		},
//   	},
//   }
//
// Experimental.
type AssemblyManifest struct {
	// Protocol version.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// The set of artifacts in this assembly.
	// Experimental.
	Artifacts *map[string]*ArtifactManifest `json:"artifacts" yaml:"artifacts"`
	// Missing context information.
	//
	// If this field has values, it means that the
	// cloud assembly is not complete and should not be deployed.
	// Experimental.
	Missing *[]*MissingContext `json:"missing" yaml:"missing"`
	// Runtime information.
	// Experimental.
	Runtime *RuntimeInfo `json:"runtime" yaml:"runtime"`
}

// Definitions for the asset manifest.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   assetManifest := &assetManifest{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	dockerImages: map[string]dockerImageAsset{
//   		"dockerImagesKey": &dockerImageAsset{
//   			"destinations": map[string]DockerImageDestination{
//   				"destinationsKey": &DockerImageDestination{
//   					"imageTag": jsii.String("imageTag"),
//   					"repositoryName": jsii.String("repositoryName"),
//
//   					// the properties below are optional
//   					"assumeRoleArn": jsii.String("assumeRoleArn"),
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"region": jsii.String("region"),
//   				},
//   			},
//   			"source": &DockerImageSource{
//   				"directory": jsii.String("directory"),
//   				"dockerBuildArgs": map[string]*string{
//   					"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   				},
//   				"dockerBuildTarget": jsii.String("dockerBuildTarget"),
//   				"dockerFile": jsii.String("dockerFile"),
//   				"executable": []*string{
//   					jsii.String("executable"),
//   				},
//   				"networkMode": jsii.String("networkMode"),
//   			},
//   		},
//   	},
//   	files: map[string]fileAsset{
//   		"filesKey": &fileAsset{
//   			"destinations": map[string]FileDestination{
//   				"destinationsKey": &FileDestination{
//   					"bucketName": jsii.String("bucketName"),
//   					"objectKey": jsii.String("objectKey"),
//
//   					// the properties below are optional
//   					"assumeRoleArn": jsii.String("assumeRoleArn"),
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"region": jsii.String("region"),
//   				},
//   			},
//   			"source": &FileSource{
//   				"executable": []*string{
//   					jsii.String("executable"),
//   				},
//   				"packaging": cloud_assembly_schema.FileAssetPackaging_FILE,
//   				"path": jsii.String("path"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AssetManifest struct {
	// Version of the manifest.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// The Docker image assets in this manifest.
	// Experimental.
	DockerImages *map[string]*DockerImageAsset `json:"dockerImages" yaml:"dockerImages"`
	// The file assets in this manifest.
	// Experimental.
	Files *map[string]*FileAsset `json:"files" yaml:"files"`
}

// Artifact properties for the Asset Manifest.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   assetManifestProperties := &assetManifestProperties{
//   	file: jsii.String("file"),
//
//   	// the properties below are optional
//   	bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	requiresBootstrapStackVersion: jsii.Number(123),
//   }
//
// Experimental.
type AssetManifestProperties struct {
	// Filename of the asset manifest.
	// Experimental.
	File *string `json:"file" yaml:"file"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	// Experimental.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

// Query to availability zone context provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   availabilityZonesContextQuery := &availabilityZonesContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type AvailabilityZonesContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Artifact properties for CloudFormation stacks.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   awsCloudFormationStackProperties := &awsCloudFormationStackProperties{
//   	templateFile: jsii.String("templateFile"),
//
//   	// the properties below are optional
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	lookupRole: &bootstrapRole{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		requiresBootstrapStackVersion: jsii.Number(123),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	requiresBootstrapStackVersion: jsii.Number(123),
//   	stackName: jsii.String("stackName"),
//   	stackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	terminationProtection: jsii.Boolean(false),
//   	validateOnSynth: jsii.Boolean(false),
//   }
//
// Experimental.
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	// Experimental.
	TemplateFile *string `json:"templateFile" yaml:"templateFile"`
	// The role that needs to be assumed to deploy the stack.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	// Experimental.
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
	// Experimental.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	// Experimental.
	CloudFormationExecutionRoleArn *string `json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	// Experimental.
	LookupRole *BootstrapRole `json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	// Experimental.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// The name to use for the CloudFormation stack.
	// Experimental.
	StackName *string `json:"stackName" yaml:"stackName"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	// Experimental.
	StackTemplateAssetObjectUrl *string `json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
	// Values for CloudFormation stack tags that should be passed when the stack is deployed.
	// Experimental.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Experimental.
	TerminationProtection *bool `json:"terminationProtection" yaml:"terminationProtection"`
	// Whether this stack should be validated by the CLI after synthesis.
	// Experimental.
	ValidateOnSynth *bool `json:"validateOnSynth" yaml:"validateOnSynth"`
}

// Destination for assets that need to be uploaded to AWS.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   awsDestination := &awsDestination{
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	region: jsii.String("region"),
//   }
//
// Experimental.
type AwsDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
}

// Information needed to access an IAM role created as part of the bootstrap process.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   bootstrapRole := &bootstrapRole{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	requiresBootstrapStackVersion: jsii.Number(123),
//   }
//
// Experimental.
type BootstrapRole struct {
	// The ARN of the IAM role created as part of bootrapping e.g. lookupRoleArn.
	// Experimental.
	Arn *string `json:"arn" yaml:"arn"`
	// External ID to use when assuming the bootstrap role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// Name of SSM parameter with bootstrap stack version.
	// Experimental.
	BootstrapStackVersionSsmParameter *string `json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to use this role.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

// Metadata Entry spec for container images.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   containerImageAssetMetadataEntry := &containerImageAssetMetadataEntry{
//   	id: jsii.String("id"),
//   	packaging: jsii.String("packaging"),
//   	path: jsii.String("path"),
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	file: jsii.String("file"),
//   	imageNameParameter: jsii.String("imageNameParameter"),
//   	imageTag: jsii.String("imageTag"),
//   	networkMode: jsii.String("networkMode"),
//   	repositoryName: jsii.String("repositoryName"),
//   	target: jsii.String("target"),
//   }
//
// Experimental.
type ContainerImageAssetMetadataEntry struct {
	// Logical identifier for the asset.
	// Experimental.
	Id *string `json:"id" yaml:"id"`
	// Type of asset.
	// Experimental.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
	// The hash of the asset source.
	// Experimental.
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
	// Build args to pass to the `docker build` command.
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `json:"file" yaml:"file"`
	// ECR Repository name and repo digest (separated by "@sha256:") where this image is stored.
	// Deprecated: specify `repositoryName` and `imageTag` instead, and then you
	// know where the image will go.
	ImageNameParameter *string `json:"imageNameParameter" yaml:"imageNameParameter"`
	// The docker image tag to use for tagging pushed images.
	//
	// This field is
	// required if `imageParameterName` is ommited (otherwise, the app won't be
	// able to find the image).
	// Experimental.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
	// Networking mode for the RUN commands during build.
	// Experimental.
	NetworkMode *string `json:"networkMode" yaml:"networkMode"`
	// ECR repository name, if omitted a default name based on the asset's ID is used instead.
	//
	// Specify this property if you need to statically address the
	// image, e.g. from a Kubernetes Pod. Note, this is only the repository name,
	// without the registry and the tag parts.
	// Experimental.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
}

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

// A file asset.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   dockerImageAsset := &dockerImageAsset{
//   	destinations: map[string]dockerImageDestination{
//   		"destinationsKey": &dockerImageDestination{
//   			"imageTag": jsii.String("imageTag"),
//   			"repositoryName": jsii.String("repositoryName"),
//
//   			// the properties below are optional
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	source: &dockerImageSource{
//   		directory: jsii.String("directory"),
//   		dockerBuildArgs: map[string]*string{
//   			"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   		},
//   		dockerBuildTarget: jsii.String("dockerBuildTarget"),
//   		dockerFile: jsii.String("dockerFile"),
//   		executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		networkMode: jsii.String("networkMode"),
//   	},
//   }
//
// Experimental.
type DockerImageAsset struct {
	// Destinations for this file asset.
	// Experimental.
	Destinations *map[string]*DockerImageDestination `json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	// Experimental.
	Source *DockerImageSource `json:"source" yaml:"source"`
}

// Where to publish docker images.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   dockerImageDestination := &dockerImageDestination{
//   	imageTag: jsii.String("imageTag"),
//   	repositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	region: jsii.String("region"),
//   }
//
// Experimental.
type DockerImageDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Tag of the image to publish.
	// Experimental.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
	// Name of the ECR repository to publish to.
	// Experimental.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// Properties for how to produce a Docker image from a source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   dockerImageSource := &dockerImageSource{
//   	directory: jsii.String("directory"),
//   	dockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	dockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	dockerFile: jsii.String("dockerFile"),
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	networkMode: jsii.String("networkMode"),
//   }
//
// Experimental.
type DockerImageSource struct {
	// The directory containing the Docker image build instructions.
	//
	// This path is relative to the asset manifest location.
	// Experimental.
	Directory *string `json:"directory" yaml:"directory"`
	// Additional build arguments.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildArgs *map[string]*string `json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Target build stage in a Dockerfile with multiple build stages.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildTarget *string `json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// The name of the file with build instructions.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerFile *string `json:"dockerFile" yaml:"dockerFile"`
	// A command-line executable that returns the name of a local Docker image on stdout after being run.
	// Experimental.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	// Experimental.
	NetworkMode *string `json:"networkMode" yaml:"networkMode"`
}

// Query to endpoint service context provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   endpointServiceAvailabilityZonesContextQuery := &endpointServiceAvailabilityZonesContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Query service name.
	// Experimental.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// A file asset.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   fileAsset := &fileAsset{
//   	destinations: map[string]fileDestination{
//   		"destinationsKey": &fileDestination{
//   			"bucketName": jsii.String("bucketName"),
//   			"objectKey": jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	source: &fileSource{
//   		executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		packaging: cloud_assembly_schema.fileAssetPackaging_FILE,
//   		path: jsii.String("path"),
//   	},
//   }
//
// Experimental.
type FileAsset struct {
	// Destinations for this file asset.
	// Experimental.
	Destinations *map[string]*FileDestination `json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	// Experimental.
	Source *FileSource `json:"source" yaml:"source"`
}

// Metadata Entry spec for files.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   fileAssetMetadataEntry := &fileAssetMetadataEntry{
//   	artifactHashParameter: jsii.String("artifactHashParameter"),
//   	id: jsii.String("id"),
//   	packaging: jsii.String("packaging"),
//   	path: jsii.String("path"),
//   	s3BucketParameter: jsii.String("s3BucketParameter"),
//   	s3KeyParameter: jsii.String("s3KeyParameter"),
//   	sourceHash: jsii.String("sourceHash"),
//   }
//
// Experimental.
type FileAssetMetadataEntry struct {
	// The name of the parameter where the hash of the bundled asset should be passed in.
	// Experimental.
	ArtifactHashParameter *string `json:"artifactHashParameter" yaml:"artifactHashParameter"`
	// Logical identifier for the asset.
	// Experimental.
	Id *string `json:"id" yaml:"id"`
	// Requested packaging style.
	// Experimental.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
	// Name of parameter where S3 bucket should be passed in.
	// Experimental.
	S3BucketParameter *string `json:"s3BucketParameter" yaml:"s3BucketParameter"`
	// Name of parameter where S3 key should be passed in.
	// Experimental.
	S3KeyParameter *string `json:"s3KeyParameter" yaml:"s3KeyParameter"`
	// The hash of the asset source.
	// Experimental.
	SourceHash *string `json:"sourceHash" yaml:"sourceHash"`
}

// Packaging strategy for file assets.
// Experimental.
type FileAssetPackaging string

const (
	// Upload the given path as a file.
	// Experimental.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	// The given path is a directory, zip it and upload.
	// Experimental.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

// Where in S3 a file asset needs to be published.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   fileDestination := &fileDestination{
//   	bucketName: jsii.String("bucketName"),
//   	objectKey: jsii.String("objectKey"),
//
//   	// the properties below are optional
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	region: jsii.String("region"),
//   }
//
// Experimental.
type FileDestination struct {
	// The role that needs to be assumed while publishing this asset.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Experimental.
	AssumeRoleExternalId *string `json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The name of the bucket.
	// Experimental.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The destination object key.
	// Experimental.
	ObjectKey *string `json:"objectKey" yaml:"objectKey"`
}

// Describe the source of a file asset.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   fileSource := &fileSource{
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	packaging: cloud_assembly_schema.fileAssetPackaging_FILE,
//   	path: jsii.String("path"),
//   }
//
// Experimental.
type FileSource struct {
	// External command which will produce the file asset to upload.
	// Experimental.
	Executable *[]*string `json:"executable" yaml:"executable"`
	// Packaging method.
	//
	// Only allowed when `path` is specified.
	// Experimental.
	Packaging FileAssetPackaging `json:"packaging" yaml:"packaging"`
	// The filesystem object to upload.
	//
	// This path is relative to the asset manifest location.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
}

// Query to hosted zone context provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   hostedZoneContextQuery := &hostedZoneContextQuery{
//   	account: jsii.String("account"),
//   	domainName: jsii.String("domainName"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	privateZone: jsii.Boolean(false),
//   	vpcId: jsii.String("vpcId"),
//   }
//
// Experimental.
type HostedZoneContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// The domain name e.g. example.com to lookup.
	// Experimental.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// True if the zone you want to find is a private hosted zone.
	// Experimental.
	PrivateZone *bool `json:"privateZone" yaml:"privateZone"`
	// The VPC ID to that the private zone must be associated with.
	//
	// If you provide VPC ID and privateZone is false, this will return no results
	// and raise an error.
	// Experimental.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Query input for looking up a KMS Key.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   keyContextQuery := &keyContextQuery{
//   	account: jsii.String("account"),
//   	aliasName: jsii.String("aliasName"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type KeyContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Alias name used to search the Key.
	// Experimental.
	AliasName *string `json:"aliasName" yaml:"aliasName"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Query input for looking up a load balancer.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   loadBalancerContextQuery := &loadBalancerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type LoadBalancerContextQuery struct {
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Filters for selecting load balancers.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   loadBalancerFilter := &loadBalancerFilter{
//   	loadBalancerType: cloud_assembly_schema.loadBalancerType_NETWORK,
//
//   	// the properties below are optional
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type LoadBalancerFilter struct {
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Query input for looking up a load balancer listener.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   loadBalancerListenerContextQuery := &loadBalancerListenerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	listenerArn: jsii.String("listenerArn"),
//   	listenerPort: jsii.Number(123),
//   	listenerProtocol: cloud_assembly_schema.loadBalancerListenerProtocol_HTTP,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type LoadBalancerListenerContextQuery struct {
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Find by listener's arn.
	// Experimental.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener port.
	// Experimental.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// Filter by listener protocol.
	// Experimental.
	ListenerProtocol LoadBalancerListenerProtocol `json:"listenerProtocol" yaml:"listenerProtocol"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// The protocol for connections from clients to the load balancer.
// Experimental.
type LoadBalancerListenerProtocol string

const (
	// HTTP protocol.
	// Experimental.
	LoadBalancerListenerProtocol_HTTP LoadBalancerListenerProtocol = "HTTP"
	// HTTPS protocol.
	// Experimental.
	LoadBalancerListenerProtocol_HTTPS LoadBalancerListenerProtocol = "HTTPS"
	// TCP protocol.
	// Experimental.
	LoadBalancerListenerProtocol_TCP LoadBalancerListenerProtocol = "TCP"
	// TLS protocol.
	// Experimental.
	LoadBalancerListenerProtocol_TLS LoadBalancerListenerProtocol = "TLS"
	// UDP protocol.
	// Experimental.
	LoadBalancerListenerProtocol_UDP LoadBalancerListenerProtocol = "UDP"
	// TCP and UDP protocol.
	// Experimental.
	LoadBalancerListenerProtocol_TCP_UDP LoadBalancerListenerProtocol = "TCP_UDP"
)

// Type of load balancer.
// Experimental.
type LoadBalancerType string

const (
	// Network load balancer.
	// Experimental.
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
	// Application load balancer.
	// Experimental.
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
)

// Options for the loadManifest operation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   loadManifestOptions := &loadManifestOptions{
//   	skipEnumCheck: jsii.Boolean(false),
//   	skipVersionCheck: jsii.Boolean(false),
//   }
//
// Experimental.
type LoadManifestOptions struct {
	// Skip enum checks.
	//
	// This means you may read enum values you don't know about yet. Make sure to always
	// check the values of enums you encounter in the manifest.
	// Experimental.
	SkipEnumCheck *bool `json:"skipEnumCheck" yaml:"skipEnumCheck"`
	// Skip the version check.
	//
	// This means you may read a newer cloud assembly than the CX API is designed
	// to support, and your application may not be aware of all features that in use
	// in the Cloud Assembly.
	// Experimental.
	SkipVersionCheck *bool `json:"skipVersionCheck" yaml:"skipVersionCheck"`
}

// Protocol utility class.
// Experimental.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Deprecated.
// Deprecated: use `loadAssemblyManifest()`.
func Manifest_Load(filePath *string) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"load",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the cloud assembly manifest from file.
// Experimental.
func Manifest_LoadAssemblyManifest(filePath *string, options *LoadManifestOptions) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"loadAssemblyManifest",
		[]interface{}{filePath, options},
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
		"monocdk.cloud_assembly_schema.Manifest",
		"loadAssetManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Deprecated.
// Deprecated: use `saveAssemblyManifest()`.
func Manifest_Save(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"save",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the cloud assembly manifest to file.
// Experimental.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
// Experimental.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
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
		"monocdk.cloud_assembly_schema.Manifest",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A metadata entry in a cloud assembly artifact.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   metadataEntry := &metadataEntry{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	data: jsii.String("data"),
//   	trace: []*string{
//   		jsii.String("trace"),
//   	},
//   }
//
// Experimental.
type MetadataEntry struct {
	// The type of the metadata entry.
	// Experimental.
	Type *string `json:"type" yaml:"type"`
	// The data.
	// Experimental.
	Data interface{} `json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	// Experimental.
	Trace *[]*string `json:"trace" yaml:"trace"`
}

// Represents a missing piece of context.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   missingContext := &missingContext{
//   	key: jsii.String("key"),
//   	props: &amiContextQuery{
//   		account: jsii.String("account"),
//   		filters: map[string][]*string{
//   			"filtersKey": []*string{
//   				jsii.String("filters"),
//   			},
//   		},
//   		region: jsii.String("region"),
//
//   		// the properties below are optional
//   		lookupRoleArn: jsii.String("lookupRoleArn"),
//   		owners: []interface{}{
//   			jsii.String("owners"),
//   		},
//   	},
//   	provider: cloud_assembly_schema.contextProvider_AMI_PROVIDER,
//   }
//
// Experimental.
type MissingContext struct {
	// The missing context key.
	// Experimental.
	Key *string `json:"key" yaml:"key"`
	// A set of provider-specific options.
	// Experimental.
	Props interface{} `json:"props" yaml:"props"`
	// The provider from which we expect this context key to be obtained.
	// Experimental.
	Provider ContextProvider `json:"provider" yaml:"provider"`
}

// Artifact properties for nested cloud assemblies.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   nestedCloudAssemblyProperties := &nestedCloudAssemblyProperties{
//   	directoryName: jsii.String("directoryName"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   }
//
// Experimental.
type NestedCloudAssemblyProperties struct {
	// Relative path to the nested cloud assembly.
	// Experimental.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Display name for the cloud assembly.
	// Experimental.
	DisplayName *string `json:"displayName" yaml:"displayName"`
}

// Query input for plugins.
//
// This alternate branch is necessary because it needs to be able to escape all type checking
// we do on on the cloud assembly -- we cannot know the properties that will be used a priori.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   pluginContextQuery := &pluginContextQuery{
//   	pluginName: jsii.String("pluginName"),
//   }
//
// Experimental.
type PluginContextQuery struct {
	// The name of the plugin.
	// Experimental.
	PluginName *string `json:"pluginName" yaml:"pluginName"`
}

// Information about the application's runtime components.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   runtimeInfo := &runtimeInfo{
//   	libraries: map[string]*string{
//   		"librariesKey": jsii.String("libraries"),
//   	},
//   }
//
// Experimental.
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	// Experimental.
	Libraries *map[string]*string `json:"libraries" yaml:"libraries"`
}

// Query to SSM Parameter Context Provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   sSMParameterContextQuery := &sSMParameterContextQuery{
//   	account: jsii.String("account"),
//   	parameterName: jsii.String("parameterName"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type SSMParameterContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Parameter name to query.
	// Experimental.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

// Query input for looking up a security group.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   securityGroupContextQuery := &securityGroupContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	securityGroupId: jsii.String("securityGroupId"),
//   	securityGroupName: jsii.String("securityGroupName"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
// Experimental.
type SecurityGroupContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Security group id.
	// Experimental.
	SecurityGroupId *string `json:"securityGroupId" yaml:"securityGroupId"`
	// Security group name.
	// Experimental.
	SecurityGroupName *string `json:"securityGroupName" yaml:"securityGroupName"`
	// VPC ID.
	// Experimental.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Metadata Entry spec for stack tag.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   tag := &tag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// Experimental.
type Tag struct {
	// Tag key.
	//
	// (In the actual file on disk this will be cased as "Key", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	// Experimental.
	Key *string `json:"key" yaml:"key"`
	// Tag value.
	//
	// (In the actual file on disk this will be cased as "Value", and the structure is
	// patched to match this structure upon loading:
	// https://github.com/aws/aws-cdk/blob/4aadaa779b48f35838cccd4e25107b2338f05547/packages/%40aws-cdk/cloud-assembly-schema/lib/manifest.ts#L137)
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}

// Artifact properties for the Construct Tree Artifact.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   treeArtifactProperties := &treeArtifactProperties{
//   	file: jsii.String("file"),
//   }
//
// Experimental.
type TreeArtifactProperties struct {
	// Filename of the tree artifact.
	// Experimental.
	File *string `json:"file" yaml:"file"`
}

// Query input for looking up a VPC.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"
//   vpcContextQuery := &vpcContextQuery{
//   	account: jsii.String("account"),
//   	filter: map[string]*string{
//   		"filterKey": jsii.String("filter"),
//   	},
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   	returnAsymmetricSubnets: jsii.Boolean(false),
//   	subnetGroupNameTag: jsii.String("subnetGroupNameTag"),
//   }
//
// Experimental.
type VpcContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Filters to apply to the VPC.
	//
	// Filter parameters are the same as passed to DescribeVpcs.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html
	//
	// Experimental.
	Filter *map[string]*string `json:"filter" yaml:"filter"`
	// Query region.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Whether to populate the subnetGroups field of the {@link VpcContextResponse}, which contains potentially asymmetric subnet groups.
	// Experimental.
	ReturnAsymmetricSubnets *bool `json:"returnAsymmetricSubnets" yaml:"returnAsymmetricSubnets"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	// Experimental.
	SubnetGroupNameTag *string `json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
}

