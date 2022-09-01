package cloudassemblyschema


// Artifact properties for CloudFormation stacks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
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
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	TemplateFile *string `field:"required" json:"templateFile" yaml:"templateFile"`
	// The role that needs to be assumed to deploy the stack.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// Only used if `requiresBootstrapStackVersion` is set.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//    deployment time so the stack version can be looked up from the stack
	//    outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//    we won't need to look it up.
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	CloudFormationExecutionRoleArn *string `field:"optional" json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	LookupRole *BootstrapRole `field:"optional" json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion *float64 `field:"optional" json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// The name to use for the CloudFormation stack.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	StackTemplateAssetObjectUrl *string `field:"optional" json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
	// Values for CloudFormation stack tags that should be passed when the stack is deployed.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Whether this stack should be validated by the CLI after synthesis.
	ValidateOnSynth *bool `field:"optional" json:"validateOnSynth" yaml:"validateOnSynth"`
}

