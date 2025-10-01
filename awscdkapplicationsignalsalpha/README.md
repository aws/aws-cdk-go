# AWS::ApplicationSignals Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

CloudWatch Application Signals is an auto-instrumentation solution built on OpenTelemetry that enables zero-code collection of monitoring data, such
as traces and metrics, from applications running across multiple platforms. It also supports topology auto-discovery based on collected monitoring data
and includes a new feature for managing service-level objectives (SLOs).

It supports Java, Python, .NET, and Node.js on platforms including EKS (and native Kubernetes), Lambda, ECS, and EC2. For more details, visit
[Application Signals](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Monitoring-Sections.html) on the AWS
public website.

## Application Signals Enablement L2 Constructs

A collection of L2 constructs which leverages native L1 CFN resources, simplifying the enablement steps and the creation of Application
Signals resources.

### ApplicationSignalsIntegration

`ApplicationSignalsIntegration` aims to address key challenges in the current CDK enablement process, which requires complex manual configurations for
ECS customers. Application Signals is designed to be flexible and is supported for other platforms as well. However, the initial focus is on supporting
ECS, with plans to potentially extend support to other platforms in the future.

#### Enable Application Signals on ECS with sidecar mode

1. Configure `instrumentation` to instrument the application with the ADOT SDK Agent.
2. Specify `cloudWatchAgentSidecar` to configure the CloudWatch Agent as a sidecar container.

```go
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

type myStack struct {
	stack
}

func newMyStack(scope construct, id *string, props stackProps) *myStack {
	if props == nil {
		props = &StackProps{
		}
	}
	this := &myStack{}
	cdk.NewStack_Override(this, )
	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
	})
	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
		Vpc: Vpc,
	})

	fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &FargateTaskDefinitionProps{
		Cpu: jsii.Number(2048),
		MemoryLimitMiB: jsii.Number(4096),
	})

	fargateTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
	})

	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
		TaskDefinition: fargateTaskDefinition,
		Instrumentation: &InstrumentationProps{
			SdkVersion: appsignals.JavaInstrumentationVersion_V2_10_0(),
		},
		ServiceName: jsii.String("sample-app"),
		CloudWatchAgentSidecar: &CloudWatchAgentOptions{
			ContainerName: jsii.String("cloudwatch-agent"),
			EnableLogging: jsii.Boolean(true),
			Cpu: jsii.Number(256),
			MemoryLimitMiB: jsii.Number(512),
		},
	})

	ecs.NewFargateService(this, jsii.String("MySampleApp"), &FargateServiceProps{
		Cluster: cluster,
		TaskDefinition: fargateTaskDefinition,
		DesiredCount: jsii.Number(1),
	})
	return this
}
```

#### Enable Application Signals on ECS with daemon mode

Note: Since the daemon deployment strategy is not supported on ECS Fargate, this mode is only supported on ECS on EC2.

1. Run CloudWatch Agent as a daemon service with HOST network mode.
2. Configure `instrumentation` to instrument the application with the ADOT Python Agent.
3. Override environment variables by configuring `overrideEnvironments` to use service connect endpoints to communicate to the CloudWatch agent server

```go
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

type myStack struct {
	stack
}

func newMyStack(scope construct, id *string, props stackProps) *myStack {
	if props == nil {
		props = &StackProps{
		}
	}
	this := &myStack{}
	cdk.NewStack_Override(this, scope, id, props)

	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
	})
	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
		Vpc: Vpc,
	})

	// Define Task Definition for CloudWatch agent (Daemon)
	cwAgentTaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("CloudWatchAgentTaskDefinition"), &Ec2TaskDefinitionProps{
		NetworkMode: ecs.NetworkMode_HOST,
	})

	appsignals.NewCloudWatchAgentIntegration(this, jsii.String("CloudWatchAgentIntegration"), &CloudWatchAgentIntegrationProps{
		TaskDefinition: cwAgentTaskDefinition,
		ContainerName: jsii.String("ecs-cwagent"),
		EnableLogging: jsii.Boolean(false),
		Cpu: jsii.Number(128),
		MemoryLimitMiB: jsii.Number(64),
		PortMappings: []portMapping{
			&portMapping{
				ContainerPort: jsii.Number(4316),
				HostPort: jsii.Number(4316),
			},
			&portMapping{
				ContainerPort: jsii.Number(2000),
				HostPort: jsii.Number(2000),
			},
		},
	})

	// Create the CloudWatch Agent daemon service
	// Create the CloudWatch Agent daemon service
	ecs.NewEc2Service(this, jsii.String("CloudWatchAgentDaemon"), &Ec2ServiceProps{
		Cluster: Cluster,
		TaskDefinition: cwAgentTaskDefinition,
		Daemon: jsii.Boolean(true),
	})

	// Define Task Definition for user application
	sampleAppTaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &Ec2TaskDefinitionProps{
		NetworkMode: ecs.NetworkMode_HOST,
	})

	sampleAppTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
		Cpu: jsii.Number(0),
		MemoryLimitMiB: jsii.Number(512),
	})

	// No CloudWatch Agent side car is needed as application container communicates to CloudWatch Agent daemon through host network
	// No CloudWatch Agent side car is needed as application container communicates to CloudWatch Agent daemon through host network
	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
		TaskDefinition: sampleAppTaskDefinition,
		Instrumentation: &InstrumentationProps{
			SdkVersion: appsignals.PythonInstrumentationVersion_V0_8_0(),
		},
		ServiceName: jsii.String("sample-app"),
	})

	ecs.NewEc2Service(this, jsii.String("MySampleApp"), &Ec2ServiceProps{
		Cluster: Cluster,
		TaskDefinition: sampleAppTaskDefinition,
		DesiredCount: jsii.Number(1),
	})
	return this
}
```

#### Enable Application Signals on ECS with replica mode

**Note**
*Running CloudWatch Agent service using replica mode requires specific security group configurations to enable communication with other services.
For Application Signals functionality, configure the security group with the following minimum inbound rules: Port 2000 (HTTP) and Port 4316 (HTTP).
This configuration ensures proper connectivity between the CloudWatch Agent and dependent services.*

1. Run CloudWatch Agent as a replica service with service connect.
2. Configure `instrumentation` to instrument the application with the ADOT Python Agent.
3. Override environment variables by configuring `overrideEnvironments` to use service connect endpoints to communicate to the CloudWatch agent server

```go
import "github.com/aws/constructs-go/constructs"
import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
import cdk "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

type myStack struct {
	stack
}

func newMyStack(scope construct, id *string, props stackProps) *myStack {
	if props == nil {
		props = &StackProps{
		}
	}
	this := &myStack{}
	cdk.NewStack_Override(this, scope, id, props)

	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
	})
	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
		Vpc: Vpc,
	})
	dnsNamespace := awscdk.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
		Vpc: Vpc,
		Name: jsii.String("local"),
	})
	securityGroup := ec2.NewSecurityGroup(this, jsii.String("ECSSG"), &SecurityGroupProps{
		Vpc: Vpc,
	})
	securityGroup.AddIngressRule(securityGroup, ec2.Port_TcpRange(jsii.Number(0), jsii.Number(65535)))

	// Define Task Definition for CloudWatch agent (Replica)
	cwAgentTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("CloudWatchAgentTaskDefinition"), &FargateTaskDefinitionProps{
	})

	appsignals.NewCloudWatchAgentIntegration(this, jsii.String("CloudWatchAgentIntegration"), &CloudWatchAgentIntegrationProps{
		TaskDefinition: cwAgentTaskDefinition,
		ContainerName: jsii.String("ecs-cwagent"),
		EnableLogging: jsii.Boolean(false),
		Cpu: jsii.Number(128),
		MemoryLimitMiB: jsii.Number(64),
		PortMappings: []portMapping{
			&portMapping{
				Name: jsii.String("cwagent-4316"),
				ContainerPort: jsii.Number(4316),
				HostPort: jsii.Number(4316),
			},
			&portMapping{
				Name: jsii.String("cwagent-2000"),
				ContainerPort: jsii.Number(2000),
				HostPort: jsii.Number(2000),
			},
		},
	})

	// Create the CloudWatch Agent replica service with service connect
	// Create the CloudWatch Agent replica service with service connect
	ecs.NewFargateService(this, jsii.String("CloudWatchAgentService"), &FargateServiceProps{
		Cluster: cluster,
		TaskDefinition: cwAgentTaskDefinition,
		SecurityGroups: []iSecurityGroup{
			securityGroup,
		},
		ServiceConnectConfiguration: &ServiceConnectProps{
			Namespace: dnsNamespace.NamespaceArn,
			Services: []serviceConnectService{
				&serviceConnectService{
					PortMappingName: jsii.String("cwagent-4316"),
					DnsName: jsii.String("cwagent-4316-http"),
					Port: jsii.Number(4316),
				},
				&serviceConnectService{
					PortMappingName: jsii.String("cwagent-2000"),
					DnsName: jsii.String("cwagent-2000-http"),
					Port: jsii.Number(2000),
				},
			},
		},
		DesiredCount: jsii.Number(1),
	})

	// Define Task Definition for user application
	sampleAppTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &FargateTaskDefinitionProps{
	})

	sampleAppTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
		Cpu: jsii.Number(0),
		MemoryLimitMiB: jsii.Number(512),
	})

	// Overwrite environment variables to connect to the CloudWatch Agent service just created
	// Overwrite environment variables to connect to the CloudWatch Agent service just created
	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
		TaskDefinition: sampleAppTaskDefinition,
		Instrumentation: &InstrumentationProps{
			SdkVersion: appsignals.PythonInstrumentationVersion_V0_8_0(),
		},
		ServiceName: jsii.String("sample-app"),
		OverrideEnvironments: []environmentExtension{
			&environmentExtension{
				Name: appsignals.CommonExporting_OTEL_AWS_APPLICATION_SIGNALS_EXPORTER_ENDPOINT(),
				Value: jsii.String("http://cwagent-4316-http:4316/v1/metrics"),
			},
			&environmentExtension{
				Name: appsignals.TraceExporting_OTEL_EXPORTER_OTLP_TRACES_ENDPOINT(),
				Value: jsii.String("http://cwagent-4316-http:4316/v1/traces"),
			},
			&environmentExtension{
				Name: appsignals.TraceExporting_OTEL_TRACES_SAMPLER_ARG(),
				Value: jsii.String("endpoint=http://cwagent-2000-http:2000"),
			},
		},
	})

	// Create ECS Service with service connect configuration
	// Create ECS Service with service connect configuration
	ecs.NewFargateService(this, jsii.String("MySampleApp"), &FargateServiceProps{
		Cluster: cluster,
		TaskDefinition: sampleAppTaskDefinition,
		ServiceConnectConfiguration: &ServiceConnectProps{
			Namespace: dnsNamespace.*NamespaceArn,
		},
		DesiredCount: jsii.Number(1),
	})
	return this
}
```
