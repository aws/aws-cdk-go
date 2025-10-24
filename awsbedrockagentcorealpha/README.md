# Amazon Bedrock AgentCore Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

| **Language**                                                                                   | **Package**                             |
| :--------------------------------------------------------------------------------------------- | --------------------------------------- |
| ![Typescript Logo](https://docs.aws.amazon.com/cdk/api/latest/img/typescript32.png) TypeScript | `@aws-cdk/aws-bedrock-agentcore-alpha` |

[Amazon Bedrock AgentCore](https://aws.amazon.com/bedrock/agentcore/) enables you to deploy and operate highly capable AI agents securely, at scale. It offers infrastructure purpose-built for dynamic agent workloads, powerful tools to enhance agents, and essential controls for real-world deployment. AgentCore services can be used together or independently and work with any framework including CrewAI, LangGraph, LlamaIndex, and Strands Agents, as well as any foundation model in or outside of Amazon Bedrock, giving you ultimate flexibility. AgentCore eliminates the undifferentiated heavy lifting of building specialized agent infrastructure, so you can accelerate agents to production.

This construct library facilitates the deployment of Bedrock AgentCore primitives, enabling you to create sophisticated AI applications that can interact with your systems and data sources.

## Table of contents

* [AgentCore Runtime](#agentcore-runtime)

  * [Runtime Versioning](#runtime-versioning)
  * [Runtime Endpoints](#runtime-endpoints)
  * [AgentCore Runtime Properties](#agentcore-runtime-properties)
  * [Runtime Endpoint Properties](#runtime-endpoint-properties)
  * [Creating a Runtime](#creating-a-runtime)

    * [Option 1: Use an existing image in ECR](#option-1-use-an-existing-image-in-ecr)
    * [Managing Endpoints and Versions](#managing-endpoints-and-versions)
    * [Option 2: Use a local asset](#option-2-use-a-local-asset)
* [Browser Custom tool](#browser)

  * [Browser properties](#browser-properties)
  * [Browser Network modes](#browser-network-modes)
  * [Basic Browser Creation](#basic-browser-creation)
  * [Browser IAM permissions](#browser-iam-permissions)
* [Code Interpreter Custom tool](#code-interpreter)

  * [Code Interpreter properties](#code-interpreter-properties)
  * [Code Interpreter Network Modes](#code-interpreter-network-modes)
  * [Basic Code Interpreter Creation](#basic-code-interpreter-creation)
  * [Code Interpreter IAM permissions](#code-interpreter-iam-permissions)

## AgentCore Runtime

The AgentCore Runtime construct enables you to deploy containerized agents on Amazon Bedrock AgentCore.
This L2 construct simplifies runtime creation just pass your ECR repository name
and the construct handles all the configuration with sensible defaults.

### Runtime Endpoints

Endpoints provide a stable way to invoke specific versions of your agent runtime, enabling controlled deployments across different environments.
When you create an agent runtime, Amazon Bedrock AgentCore automatically creates a "DEFAULT" endpoint which always points to the latest version
of runtime.

You can create additional endpoints in two ways:

1. **Using Runtime.addEndpoint()** - Convenient method when creating endpoints alongside the runtime.
2. **Using RuntimeEndpoint** - Flexible approach for existing runtimes.

For example, you might keep a "production" endpoint on a stable version while testing newer versions
through a "staging" endpoint. This separation allows you to test changes thoroughly before promoting them
to production by simply updating the endpoint to point to the newer version.

### AgentCore Runtime Properties

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `runtimeName` | `string` | Yes | The name of the agent runtime. Valid characters are a-z, A-Z, 0-9, _ (underscore). Must start with a letter and can be up to 48 characters long |
| `agentRuntimeArtifact` | `AgentRuntimeArtifact` | Yes | The artifact configuration for the agent runtime containing the container configuration with ECR URI |
| `executionRole` | `iam.IRole` | No | The IAM role that provides permissions for the agent runtime. If not provided, a role will be created automatically |
| `networkConfiguration` | `NetworkConfiguration` | No | Network configuration for the agent runtime. Defaults to `RuntimeNetworkConfiguration.usingPublicNetwork()` |
| `description` | `string` | No | Optional description for the agent runtime |
| `protocolConfiguration` | `ProtocolType` | No | Protocol configuration for the agent runtime. Defaults to `ProtocolType.HTTP` |
| `authorizerConfiguration` | `RuntimeAuthorizerConfiguration` | No | Authorizer configuration for the agent runtime. Use `RuntimeAuthorizerConfiguration` static methods to create configurations for IAM, Cognito, JWT, or OAuth authentication |
| `environmentVariables` | `{ [key: string]: string }` | No | Environment variables for the agent runtime. Maximum 50 environment variables |
| `tags` | `{ [key: string]: string }` | No | Tags for the agent runtime. A list of key:value pairs of tags to apply to this Runtime resource |

### Runtime Endpoint Properties

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `endpointName` | `string` | Yes | The name of the runtime endpoint. Valid characters are a-z, A-Z, 0-9, _ (underscore). Must start with a letter and can be up to 48 characters long |
| `agentRuntimeId` | `string` | Yes | The Agent Runtime ID for this endpoint |
| `agentRuntimeVersion` | `string` | Yes | The Agent Runtime version for this endpoint. Must be between 1 and 5 characters long.|
| `description` | `string` | No | Optional description for the runtime endpoint |
| `tags` | `{ [key: string]: string }` | No | Tags for the runtime endpoint |

### Creating a Runtime

#### Option 1: Use an existing image in ECR

Reference an image available within ECR.

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

// The runtime by default create ECR permission only for the repository available in the account the stack is being deployed
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

// Create runtime using the built image
runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
})
```

To grant the runtime permission to invoke a Bedrock model or inference profile:

```text
// Note: This example uses @aws-cdk/aws-bedrock-alpha which must be installed separately
import * as bedrock from '@aws-cdk/aws-bedrock-alpha';

// Create a cross-region inference profile for Claude 3.7 Sonnet
const inferenceProfile = bedrock.CrossRegionInferenceProfile.fromConfig({
  geoRegion: bedrock.CrossRegionInferenceProfileRegion.US,
  model: bedrock.BedrockFoundationModel.ANTHROPIC_CLAUDE_3_7_SONNET_V1_0
});

// Grant the runtime permission to invoke the inference profile
inferenceProfile.grantInvoke(runtime);
```

#### Option 2: Use a local asset

Reference a local directory containing a Dockerfile.
Images are built from a local Docker context directory (with a Dockerfile), uploaded to Amazon Elastic Container Registry (ECR)
by the CDK toolkit,and can be naturally referenced in your CDK app .

```go
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromAsset(path.join(__dirname, jsii.String("path to agent dockerfile directory")))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
})
```

### Runtime Versioning

Amazon Bedrock AgentCore automatically manages runtime versioning to ensure safe deployments and rollback capabilities.
When you create an agent runtime, AgentCore automatically creates version 1 (V1). Each subsequent update to the
runtime configuration (such as updating the container image, modifying network settings, or changing protocol configurations)
creates a new immutable version. These versions contain complete, self-contained configurations that can be referenced by endpoints,
allowing you to maintain different versions for different environments or gradually roll out updates.

#### Managing Endpoints and Versions

Amazon Bedrock AgentCore automatically manages runtime versioning to provide safe deployments and rollback capabilities. You can follow
the steps below to understand how to use versioning with runtime for controlled deployments across different environments.

##### Step 1: Initial Deployment

When you first create an agent runtime, AgentCore automatically creates Version 1 of your runtime. At this point, a DEFAULT endpoint is
automatically created that points to Version 1. This DEFAULT endpoint serves as the main access point for your runtime.

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0")),
})
```

##### Step 2: Creating Custom Endpoints

After the initial deployment, you can create additional endpoints for different environments. For example, you might create a "production"
endpoint that explicitly points to Version 1. This allows you to maintain stable access points for specific environments while keeping the
flexibility to test newer versions elsewhere.

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0")),
})

prodEndpoint := runtime.AddEndpoint(jsii.String("production"), &AddEndpointOptions{
	Version: jsii.String("1"),
	Description: jsii.String("Stable production endpoint - pinned to v1"),
})
```

##### Step 3: Runtime Update Deployment

When you update the runtime configuration (such as updating the container image, modifying network settings, or changing protocol
configurations), AgentCore automatically creates a new version (Version 2). Upon this update:

* Version 2 is created automatically with the new configuration
* The DEFAULT endpoint automatically updates to point to Version 2
* Any explicitly pinned endpoints (like the production endpoint) remain on their specified versions

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

agentRuntimeArtifactNew := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v2.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifactNew,
})
```

##### Step 4: Testing with Staging Endpoints

Once Version 2 exists, you can create a staging endpoint that points to the new version. This staging endpoint allows you to test the
new version in a controlled environment before promoting it to production. This separation ensures that production traffic continues
to use the stable version while you validate the new version.

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

agentRuntimeArtifactNew := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v2.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifactNew,
})

stagingEndpoint := runtime.AddEndpoint(jsii.String("staging"), &AddEndpointOptions{
	Version: jsii.String("2"),
	Description: jsii.String("Staging environment for testing new version"),
})
```

##### Step 5: Promoting to Production

After thoroughly testing the new version through the staging endpoint, you can update the production endpoint to point to Version 2.
This controlled promotion process ensures that you can validate changes before they affect production traffic.

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})

agentRuntimeArtifactNew := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v2.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifactNew,
})

prodEndpoint := runtime.AddEndpoint(jsii.String("production"), &AddEndpointOptions{
	Version: jsii.String("2"),
	 // New version added here
	Description: jsii.String("Stable production endpoint"),
})
```

### Creating Standalone Runtime Endpoints

RuntimeEndpoint can also be created as a standalone resource.

#### Example: Creating an endpoint for an existing runtime

```go
// Reference an existing runtime by its ID
existingRuntimeId := "abc123-runtime-id" // The ID of an existing runtime

// Create a standalone endpoint
endpoint := agentcore.NewRuntimeEndpoint(this, jsii.String("MyEndpoint"), &RuntimeEndpointProps{
	EndpointName: jsii.String("production"),
	AgentRuntimeId: existingRuntimeId,
	AgentRuntimeVersion: jsii.String("1"),
	 // Specify which version to use
	Description: jsii.String("Production endpoint for existing runtime"),
})
```

### Runtime Authentication Configuration

The AgentCore Runtime supports multiple authentication modes to secure access to your agent endpoints. Authentication is configured during runtime creation using the `RuntimeAuthorizerConfiguration` class's static factory methods.

#### IAM Authentication (Default)

IAM authentication is the default mode, when no authorizerConfiguration is set then the underlying service use IAM.

#### Cognito Authentication

To configure AWS Cognito User Pool authentication:

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingCognito(jsii.String("us-west-2_ABC123"), jsii.String("client123"), jsii.String("us-west-2")),
})
```

#### JWT Authentication

To configure custom JWT authentication with your own OpenID Connect (OIDC) provider:

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingJWT(jsii.String("https://example.com/.well-known/openid-configuration"), []*string{
		jsii.String("client1"),
		jsii.String("client2"),
	}, []*string{
		jsii.String("audience1"),
	}),
})
```

**Note**: The discovery URL must end with `/.well-known/openid-configuration`.

#### OAuth Authentication

To configure OAuth 2.0 authentication:

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingOAuth(jsii.String("https://github.com/.well-known/openid-configuration"), jsii.String("oauth_client_123")),
})
```

#### Using a Custom IAM Role

Instead of using the auto-created execution role, you can provide your own IAM role with specific permissions:
The auto-created role includes all necessary baseline permissions for ECR access, CloudWatch logging, and X-Ray tracing. When providing a custom role, ensure these permissions are included.

### Runtime Network Configuration

The AgentCore Runtime supports two network modes for deployment:

#### Public Network Mode (Default)

By default, runtimes are deployed in PUBLIC network mode, which provides internet access suitable for less sensitive or open-use scenarios:

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

// Explicitly using public network (this is the default)
runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	NetworkConfiguration: agentcore.RuntimeNetworkConfiguration_UsingPublicNetwork(),
})
```

#### VPC Network Mode

For enhanced security and network isolation, you can deploy your runtime within a VPC:

```go
repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

// Create or use an existing VPC
vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &VpcProps{
	MaxAzs: jsii.Number(2),
})

// Configure runtime with VPC
runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	NetworkConfiguration: agentcore.RuntimeNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: vpc,
		VpcSubnets: &SubnetSelection{
			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
	}),
})
```

#### Managing Security Groups with VPC Configuration

When using VPC mode, the Runtime implements `ec2.IConnectable`, allowing you to manage network access using the `connections` property:

```go
vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &VpcProps{
	MaxAzs: jsii.Number(2),
})

repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
	RepositoryName: jsii.String("test-agent-runtime"),
})
agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))

// Create runtime with VPC configuration
runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
	RuntimeName: jsii.String("myAgent"),
	AgentRuntimeArtifact: agentRuntimeArtifact,
	NetworkConfiguration: agentcore.RuntimeNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: vpc,
		VpcSubnets: &SubnetSelection{
			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
	}),
})

// Now you can manage network access using the connections property
// Allow inbound HTTPS traffic from a specific security group
webServerSecurityGroup := ec2.NewSecurityGroup(this, jsii.String("WebServerSG"), &SecurityGroupProps{
	Vpc: Vpc,
})
runtime.connections.AllowFrom(webServerSecurityGroup, ec2.Port_Tcp(jsii.Number(443)), jsii.String("Allow HTTPS from web servers"))

// Allow outbound connections to a database
databaseSecurityGroup := ec2.NewSecurityGroup(this, jsii.String("DatabaseSG"), &SecurityGroupProps{
	Vpc: Vpc,
})
runtime.connections.AllowTo(databaseSecurityGroup, ec2.Port_Tcp(jsii.Number(5432)), jsii.String("Allow PostgreSQL connection"))

// Allow outbound HTTPS to anywhere (for external API calls)
runtime.connections.AllowToAnyIpv4(ec2.Port_Tcp(jsii.Number(443)), jsii.String("Allow HTTPS outbound"))
```

## Browser

The Amazon Bedrock AgentCore Browser provides a secure, cloud-based browser that enables AI agents to interact with websites. It includes security features such as session isolation, built-in observability through live viewing, CloudTrail logging, and session replay capabilities.

Additional information about the browser tool can be found in the [official documentation](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/browser-tool.html)

### Browser Network modes

The Browser construct supports the following network modes:

1. **Public Network Mode** (`BrowserNetworkMode.usingPublicNetwork()`) - Default

   * Allows internet access for web browsing and external API calls
   * Suitable for scenarios where agents need to interact with publicly available websites
   * Enables full web browsing capabilities
   * VPC mode is not supported with this option
2. **VPC (Virtual Private Cloud)** (`BrowserNetworkMode.usingVpc()`)

   * Select whether to run the browser in a virtual private cloud (VPC).
   * By configuring VPC connectivity, you enable secure access to private resources such as databases, internal APIs, and services within your VPC.

   While the VPC itself is mandatory, these are optional:

   * Subnets - if not provided, CDK will select appropriate subnets from the VPC
   * Security Groups - if not provided, CDK will create a default security group
   * Specific subnet selection criteria - you can let CDK choose automatically

For more information on VPC connectivity for Amazon Bedrock AgentCore Browser, please refer to the [official documentation](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agentcore-vpc.html).

### Browser Properties

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `browserCustomName` | `string` | Yes | The name of the browser. Must start with a letter and can be up to 48 characters long. Pattern: `[a-zA-Z][a-zA-Z0-9_]{0,47}` |
| `description` | `string` | No | Optional description for the browser. Can have up to 200 characters |
| `networkConfiguration` | `BrowserNetworkConfiguration` | No | Network configuration for browser. Defaults to PUBLIC network mode |
| `recordingConfig` | `RecordingConfig` | No | Recording configuration for browser. Defaults to no recording |
| `executionRole` | `iam.IRole` | No | The IAM role that provides permissions for the browser to access AWS services. A new role will be created if not provided |
| `tags` | `{ [key: string]: string }` | No | Tags to apply to the browser resource |

### Basic Browser Creation

```go
// Create a basic browser with public network access
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("A browser for web automation"),
})
```

### Browser with Tags

```go
// Create a browser with custom tags
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("A browser for web automation with tags"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
	Tags: map[string]*string{
		"Environment": jsii.String("Production"),
		"Team": jsii.String("AI/ML"),
		"Project": jsii.String("AgentCore"),
	},
})
```

### Browser with VPC

```go
browser := agentcore.NewBrowserCustom(this, jsii.String("BrowserVpcWithRecording"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("browser_recording"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
			RestrictDefaultSecurityGroup: jsii.Boolean(false),
		}),
	}),
})
```

Browser exposes a [connections](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Connections.html) property. This property returns a connections object, which simplifies the process of defining and managing ingress and egress rules for security groups in your AWS CDK applications. Instead of directly manipulating security group rules, you interact with the Connections object of a construct, which then translates your connectivity requirements into the appropriate security group rules. For instance:

```go
vpc := ec2.NewVpc(this, jsii.String("testVPC"))

browser := agentcore.NewBrowserCustom(this, jsii.String("test-browser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("test_browser"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: vpc,
	}),
})

browser.connections.AddSecurityGroup(ec2.NewSecurityGroup(this, jsii.String("AdditionalGroup"), &SecurityGroupProps{
	Vpc: Vpc,
}))
```

So security groups can be added after the browser construct creation. You can use methods like allowFrom() and allowTo() to grant ingress access to/egress access from a specified peer over a given portRange. The Connections object automatically adds the necessary ingress or egress rules to the security group(s) associated with the calling construct.

### Browser with Recording Configuration

```go
// Create an S3 bucket for recordings
recordingBucket := s3.NewBucket(this, jsii.String("RecordingBucket"), &BucketProps{
	BucketName: jsii.String("my-browser-recordings"),
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Create browser with recording enabled
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("Browser with recording enabled"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
	RecordingConfig: &RecordingConfig{
		Enabled: jsii.Boolean(true),
		S3Location: &Location{
			BucketName: recordingBucket.BucketName,
			ObjectKey: jsii.String("browser-recordings/"),
		},
	},
})
```

### Browser with Custom Execution Role

```go
// Create a custom execution role
executionRole := iam.NewRole(this, jsii.String("BrowserExecutionRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
	ManagedPolicies: []IManagedPolicy{
		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonBedrockAgentCoreBrowserExecutionRolePolicy")),
	},
})

// Create browser with custom execution role
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("Browser with custom execution role"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
	ExecutionRole: executionRole,
})
```

### Browser with S3 Recording and Permissions

```go
// Create an S3 bucket for recordings
recordingBucket := s3.NewBucket(this, jsii.String("RecordingBucket"), &BucketProps{
	BucketName: jsii.String("my-browser-recordings"),
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Create browser with recording enabled
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("Browser with recording enabled"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
	RecordingConfig: &RecordingConfig{
		Enabled: jsii.Boolean(true),
		S3Location: &Location{
			BucketName: recordingBucket.BucketName,
			ObjectKey: jsii.String("browser-recordings/"),
		},
	},
})
```

### Browser IAM Permissions

The Browser construct provides convenient methods for granting IAM permissions:

```go
// Create a browser
browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
	BrowserCustomName: jsii.String("my_browser"),
	Description: jsii.String("Browser for web automation"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
})

// Create a role that needs access to the browser
userRole := iam.NewRole(this, jsii.String("UserRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

// Grant read permissions (Get and List actions)
browser.GrantRead(userRole)

// Grant use permissions (Start, Update, Stop actions)
browser.GrantUse(userRole)

// Grant specific custom permissions
browser.Grant(userRole, jsii.String("bedrock-agentcore:GetBrowserSession"))
```

## Code Interpreter

The Amazon Bedrock AgentCore Code Interpreter enables AI agents to write and execute code securely in sandbox environments, enhancing their accuracy and expanding their ability to solve complex end-to-end tasks. This is critical in Agentic AI applications where the agents may execute arbitrary code that can lead to data compromise or security risks. The AgentCore Code Interpreter tool provides secure code execution, which helps you avoid running into these issues.

For more information about code interpreter, please refer to the [official documentation](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/code-interpreter-tool.html)

### Code Interpreter Network Modes

The Code Interpreter construct supports the following network modes:

1. **Public Network Mode** (`CodeInterpreterNetworkMode.usingPublicNetwork()`) - Default

   * Allows internet access for package installation and external API calls
   * Suitable for development and testing environments
   * Enables downloading Python packages from PyPI
2. **Sandbox Network Mode** (`CodeInterpreterNetworkMode.usingSandboxNetwork()`)

   * Isolated network environment with no internet access
   * Suitable for production environments with strict security requirements
   * Only allows access to pre-installed packages and local resources
3. **VPC (Virtual Private Cloud)** (`CodeInterpreterNetworkMode.usingVpc()`)

   * Select whether to run the browser in a virtual private cloud (VPC).
   * By configuring VPC connectivity, you enable secure access to private resources such as databases, internal APIs, and services within your VPC.

   While the VPC itself is mandatory, these are optional:

   * Subnets - if not provided, CDK will select appropriate subnets from the VPC
   * Security Groups - if not provided, CDK will create a default security group
   * Specific subnet selection criteria - you can let CDK choose automatically

For more information on VPC connectivity for Amazon Bedrock AgentCore Browser, please refer to the [official documentation](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agentcore-vpc.html).

### Code Interpreter Properties

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `codeInterpreterCustomName` | `string` | Yes | The name of the code interpreter. Must start with a letter and can be up to 48 characters long. Pattern: `[a-zA-Z][a-zA-Z0-9_]{0,47}` |
| `description` | `string` | No | Optional description for the code interpreter. Can have up to 200 characters |
| `executionRole` | `iam.IRole` | No | The IAM role that provides permissions for the code interpreter to access AWS services. A new role will be created if not provided |
| `networkConfiguration` | `CodeInterpreterNetworkConfiguration` | No | Network configuration for code interpreter. Defaults to PUBLIC network mode |
| `tags` | `{ [key: string]: string }` | No | Tags to apply to the code interpreter resource |

### Basic Code Interpreter Creation

```go
// Create a basic code interpreter with public network access
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_code_interpreter"),
	Description: jsii.String("A code interpreter for Python execution"),
})
```

### Code Interpreter with VPC

```go
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
	Description: jsii.String("Code interpreter with isolated network access"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
			RestrictDefaultSecurityGroup: jsii.Boolean(false),
		}),
	}),
})
```

Code Interpreter exposes a [connections](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Connections.html) property. This property returns a connections object, which simplifies the process of defining and managing ingress and egress rules for security groups in your AWS CDK applications. Instead of directly manipulating security group rules, you interact with the Connections object of a construct, which then translates your connectivity requirements into the appropriate security group rules. For instance:

```go
vpc := ec2.NewVpc(this, jsii.String("testVPC"))

codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
	Description: jsii.String("Code interpreter with isolated network access"),
	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
		Vpc: vpc,
	}),
})

codeInterpreter.connections.AddSecurityGroup(ec2.NewSecurityGroup(this, jsii.String("AdditionalGroup"), &SecurityGroupProps{
	Vpc: Vpc,
}))
```

So security groups can be added after the browser construct creation. You can use methods like allowFrom() and allowTo() to grant ingress access to/egress access from a specified peer over a given portRange. The Connections object automatically adds the necessary ingress or egress rules to the security group(s) associated with the calling construct.

### Code Interpreter with Sandbox Network Mode

```go
// Create code interpreter with sandbox network mode (isolated)
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
	Description: jsii.String("Code interpreter with isolated network access"),
	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingSandboxNetwork(),
})
```

### Code Interpreter with Custom Execution Role

```go
// Create a custom execution role
executionRole := iam.NewRole(this, jsii.String("CodeInterpreterExecutionRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
})

// Create code interpreter with custom execution role
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_code_interpreter"),
	Description: jsii.String("Code interpreter with custom execution role"),
	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingPublicNetwork(),
	ExecutionRole: executionRole,
})
```

### Code Interpreter IAM Permissions

The Code Interpreter construct provides convenient methods for granting IAM permissions:

```go
// Create a code interpreter
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_code_interpreter"),
	Description: jsii.String("Code interpreter for Python execution"),
	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingPublicNetwork(),
})

// Create a role that needs access to the code interpreter
userRole := iam.NewRole(this, jsii.String("UserRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

// Grant read permissions (Get and List actions)
codeInterpreter.GrantRead(userRole)

// Grant use permissions (Start, Invoke, Stop actions)
codeInterpreter.GrantUse(userRole)

// Grant specific custom permissions
codeInterpreter.Grant(userRole, jsii.String("bedrock-agentcore:GetCodeInterpreterSession"))
```

### Code interpreter with tags

```go
// Create code interpreter with sandbox network mode (isolated)
codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
	Description: jsii.String("Code interpreter with isolated network access"),
	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingPublicNetwork(),
	Tags: map[string]*string{
		"Environment": jsii.String("Production"),
		"Team": jsii.String("AI/ML"),
		"Project": jsii.String("AgentCore"),
	},
})
```
