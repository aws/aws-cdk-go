package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// OpenTelemetry trace exporter and sampling configurations.
//
// Contains constants for trace endpoints, sampling strategies, and propagation formats.
//
// Example:
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string, props stackProps) *myStack {
//   	if props == nil {
//   		props = &StackProps{
//   		}
//   	}
//   	this := &myStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	vpc := ec2.NewVpc(this, jsii.String("TestVpc"), &VpcProps{
//   	})
//   	cluster := ecs.NewCluster(this, jsii.String("TestCluster"), &ClusterProps{
//   		Vpc: Vpc,
//   	})
//   	dnsNamespace := awscdk.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
//   		Vpc: Vpc,
//   		Name: jsii.String("local"),
//   	})
//   	securityGroup := ec2.NewSecurityGroup(this, jsii.String("ECSSG"), &SecurityGroupProps{
//   		Vpc: Vpc,
//   	})
//   	securityGroup.AddIngressRule(securityGroup, ec2.Port_TcpRange(jsii.Number(0), jsii.Number(65535)))
//
//   	// Define Task Definition for CloudWatch agent (Replica)
//   	cwAgentTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("CloudWatchAgentTaskDefinition"), &FargateTaskDefinitionProps{
//   	})
//
//   	appsignals.NewCloudWatchAgentIntegration(this, jsii.String("CloudWatchAgentIntegration"), &CloudWatchAgentIntegrationProps{
//   		TaskDefinition: cwAgentTaskDefinition,
//   		ContainerName: jsii.String("ecs-cwagent"),
//   		EnableLogging: jsii.Boolean(false),
//   		Cpu: jsii.Number(128),
//   		MemoryLimitMiB: jsii.Number(64),
//   		PortMappings: []portMapping{
//   			&portMapping{
//   				Name: jsii.String("cwagent-4316"),
//   				ContainerPort: jsii.Number(4316),
//   				HostPort: jsii.Number(4316),
//   			},
//   			&portMapping{
//   				Name: jsii.String("cwagent-2000"),
//   				ContainerPort: jsii.Number(2000),
//   				HostPort: jsii.Number(2000),
//   			},
//   		},
//   	})
//
//   	// Create the CloudWatch Agent replica service with service connect
//   	// Create the CloudWatch Agent replica service with service connect
//   	ecs.NewFargateService(this, jsii.String("CloudWatchAgentService"), &FargateServiceProps{
//   		Cluster: cluster,
//   		TaskDefinition: cwAgentTaskDefinition,
//   		SecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   		ServiceConnectConfiguration: &ServiceConnectProps{
//   			Namespace: dnsNamespace.NamespaceArn,
//   			Services: []serviceConnectService{
//   				&serviceConnectService{
//   					PortMappingName: jsii.String("cwagent-4316"),
//   					DnsName: jsii.String("cwagent-4316-http"),
//   					Port: jsii.Number(4316),
//   				},
//   				&serviceConnectService{
//   					PortMappingName: jsii.String("cwagent-2000"),
//   					DnsName: jsii.String("cwagent-2000-http"),
//   					Port: jsii.Number(2000),
//   				},
//   			},
//   		},
//   		DesiredCount: jsii.Number(1),
//   	})
//
//   	// Define Task Definition for user application
//   	sampleAppTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("SampleAppTaskDefinition"), &FargateTaskDefinitionProps{
//   	})
//
//   	sampleAppTaskDefinition.AddContainer(jsii.String("app"), &ContainerDefinitionOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test/sample-app")),
//   		Cpu: jsii.Number(0),
//   		MemoryLimitMiB: jsii.Number(512),
//   	})
//
//   	// Overwrite environment variables to connect to the CloudWatch Agent service just created
//   	// Overwrite environment variables to connect to the CloudWatch Agent service just created
//   	appsignals.NewApplicationSignalsIntegration(this, jsii.String("ApplicationSignalsIntegration"), &ApplicationSignalsIntegrationProps{
//   		TaskDefinition: sampleAppTaskDefinition,
//   		Instrumentation: &InstrumentationProps{
//   			SdkVersion: appsignals.PythonInstrumentationVersion_V0_8_0(),
//   		},
//   		ServiceName: jsii.String("sample-app"),
//   		OverrideEnvironments: []environmentExtension{
//   			&environmentExtension{
//   				Name: appsignals.CommonExporting_OTEL_AWS_APPLICATION_SIGNALS_EXPORTER_ENDPOINT(),
//   				Value: jsii.String("http://cwagent-4316-http:4316/v1/metrics"),
//   			},
//   			&environmentExtension{
//   				Name: appsignals.TraceExporting_OTEL_EXPORTER_OTLP_TRACES_ENDPOINT(),
//   				Value: jsii.String("http://cwagent-4316-http:4316/v1/traces"),
//   			},
//   			&environmentExtension{
//   				Name: appsignals.TraceExporting_OTEL_TRACES_SAMPLER_ARG(),
//   				Value: jsii.String("endpoint=http://cwagent-2000-http:2000"),
//   			},
//   		},
//   	})
//
//   	// Create ECS Service with service connect configuration
//   	// Create ECS Service with service connect configuration
//   	ecs.NewFargateService(this, jsii.String("MySampleApp"), &FargateServiceProps{
//   		Cluster: cluster,
//   		TaskDefinition: sampleAppTaskDefinition,
//   		ServiceConnectConfiguration: &ServiceConnectProps{
//   			Namespace: dnsNamespace.*NamespaceArn,
//   		},
//   		DesiredCount: jsii.Number(1),
//   	})
//   	return this
//   }
//
// Experimental.
type TraceExporting interface {
}

// The jsii proxy struct for TraceExporting
type jsiiProxy_TraceExporting struct {
	_ byte // padding
}

// Experimental.
func NewTraceExporting() TraceExporting {
	_init_.Initialize()

	j := jsiiProxy_TraceExporting{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTraceExporting_Override(t TraceExporting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		nil, // no parameters
		t,
	)
}

func TraceExporting_OTEL_EXPORTER_OTLP_TRACES_ENDPOINT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_EXPORTER_OTLP_TRACES_ENDPOINT_LOCAL_CWA() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT_LOCAL_CWA",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_PROPAGATORS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_PROPAGATORS",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_PROPAGATORS_APPLICATION_SIGNALS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_PROPAGATORS_APPLICATION_SIGNALS",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_ALWAYS_OFF() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_ALWAYS_OFF",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_ALWAYS_ON() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_ALWAYS_ON",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_ARG() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_ARG",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_ARG_LOCAL_CWA() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_ARG_LOCAL_CWA",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_PARENT_BASED_ALWAYS_OFF() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_PARENT_BASED_ALWAYS_OFF",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_PARENT_BASED_ALWAYS_ON() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_PARENT_BASED_ALWAYS_ON",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_PARENT_BASED_TRACEID_RATIO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_PARENT_BASED_TRACEID_RATIO",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_TRACEID_RATIO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_TRACEID_RATIO",
		&returns,
	)
	return returns
}

func TraceExporting_OTEL_TRACES_SAMPLER_XRAY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.TraceExporting",
		"OTEL_TRACES_SAMPLER_XRAY",
		&returns,
	)
	return returns
}

