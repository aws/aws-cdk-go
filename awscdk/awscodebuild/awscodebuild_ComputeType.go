package awscodebuild


// Build machine compute type.
//
// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//
//   pipelines.NewCodeBuildStep(jsii.String("Synth"), &codeBuildStepProps{
//   	// ...standard ShellStep props...
//   	commands: []*string{
//   	},
//   	env: map[string]interface{}{
//   	},
//
//   	// If you are using a CodeBuildStep explicitly, set the 'cdk.out' directory
//   	// to be the synth step's output.
//   	primaryOutputDirectory: jsii.String("cdk.out"),
//
//   	// Control the name of the project
//   	projectName: jsii.String("MyProject"),
//
//   	// Control parts of the BuildSpec other than the regular 'build' and 'install' commands
//   	partialBuildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//
//   	// Control the build environment
//   	buildEnvironment: &buildEnvironment{
//   		computeType: codebuild.computeType_LARGE,
//   	},
//   	timeout: awscdk.Duration.minutes(jsii.Number(90)),
//
//   	// Control Elastic Network Interface creation
//   	vpc: vpc,
//   	subnetSelection: &subnetSelection{
//   		subnetType: ec2.subnetType_PRIVATE_WITH_NAT,
//   	},
//   	securityGroups: []iSecurityGroup{
//   		mySecurityGroup,
//   	},
//
//   	// Additional policy statements for the execution role
//   	rolePolicyStatements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   		}),
//   	},
//   })
//
// Experimental.
type ComputeType string

const (
	// Experimental.
	ComputeType_SMALL ComputeType = "SMALL"
	// Experimental.
	ComputeType_MEDIUM ComputeType = "MEDIUM"
	// Experimental.
	ComputeType_LARGE ComputeType = "LARGE"
	// Experimental.
	ComputeType_X2_LARGE ComputeType = "X2_LARGE"
)

