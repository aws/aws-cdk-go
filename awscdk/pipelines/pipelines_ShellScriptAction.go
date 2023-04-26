package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
)

// Validate a revision using shell commands.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var stackOutput stackOutput
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
//   shellScriptAction := awscdk.Pipelines.NewShellScriptAction(&ShellScriptActionProps{
//   	ActionName: jsii.String("actionName"),
//   	Commands: []*string{
//   		jsii.String("commands"),
//   	},
//
//   	// the properties below are optional
//   	AdditionalArtifacts: []*artifact{
//   		artifact,
//   	},
//   	BashOptions: jsii.String("bashOptions"),
//   	Environment: &BuildEnvironment{
//   		BuildImage: buildImage,
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: bucket,
//   			ObjectKey: jsii.String("objectKey"),
//   		},
//   		ComputeType: awscdk.Aws_codebuild.ComputeType_SMALL,
//   		EnvironmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		Privileged: jsii.Boolean(false),
//   	},
//   	EnvironmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	RolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	RunOrder: jsii.Number(123),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	SubnetSelection: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   	},
//   	UseOutputs: map[string]*stackOutput{
//   		"useOutputsKey": stackOutput,
//   	},
//   	Vpc: vpc,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type ShellScriptAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The CodeBuild Project's principal.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	GrantPrincipal() awsiam.IPrincipal
	// Project generated to run the shell script in.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Project() awscodebuild.IProject
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for ShellScriptAction
type jsiiProxy_ShellScriptAction struct {
	internal.Type__awscodepipelineIAction
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_ShellScriptAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellScriptAction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellScriptAction) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewShellScriptAction(props *ShellScriptActionProps) ShellScriptAction {
	_init_.Initialize()

	if err := validateNewShellScriptActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ShellScriptAction{}

	_jsii_.Create(
		"monocdk.pipelines.ShellScriptAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewShellScriptAction_Override(s ShellScriptAction, props *ShellScriptActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ShellScriptAction",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_ShellScriptAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ShellScriptAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := s.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

