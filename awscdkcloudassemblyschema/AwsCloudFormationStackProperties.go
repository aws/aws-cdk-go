package awscdkcloudassemblyschema


// Artifact properties for CloudFormation stacks.
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	TemplateFile *string `field:"required" json:"templateFile" yaml:"templateFile"`
	// The role that needs to be assumed to deploy the stack.
	// Default: - No role is assumed (current credentials are used).
	//
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	// Default: - No external ID.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// Only used if `requiresBootstrapStackVersion` is set.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//   deployment time so the stack version can be looked up from the stack
	//   outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//   we won't need to look it up.
	// Default: - Bootstrap stack version number looked up.
	//
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	// Default: - No role is passed (currently assumed role/credentials are used).
	//
	CloudFormationExecutionRoleArn *string `field:"optional" json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	// Default: - No role is assumed (current credentials are used).
	//
	LookupRole *BootstrapRole `field:"optional" json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	// Default: - No parameters.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	// Default: - No bootstrap stack required.
	//
	RequiresBootstrapStackVersion *float64 `field:"optional" json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// The name to use for the CloudFormation stack.
	// Default: - name derived from artifact ID.
	//
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	// Default: - Not uploaded yet, upload just before deploying.
	//
	StackTemplateAssetObjectUrl *string `field:"optional" json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
	// Values for CloudFormation stack tags that should be passed when the stack is deployed.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Default: false.
	//
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Whether this stack should be validated by the CLI after synthesis.
	// Default: - false.
	//
	ValidateOnSynth *bool `field:"optional" json:"validateOnSynth" yaml:"validateOnSynth"`
}

