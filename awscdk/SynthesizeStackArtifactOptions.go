package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

// Stack artifact options.
//
// A subset of `cxschema.AwsCloudFormationStackProperties` of optional settings that need to be
// configurable by synthesizers, plus `additionalDependencies`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   synthesizeStackArtifactOptions := &SynthesizeStackArtifactOptions{
//   	AdditionalDependencies: []*string{
//   		jsii.String("additionalDependencies"),
//   	},
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	LookupRole: &BootstrapRole{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		RequiresBootstrapStackVersion: jsii.Number(123),
//   	},
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	RequiresBootstrapStackVersion: jsii.Number(123),
//   	StackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   }
//
// Experimental.
type SynthesizeStackArtifactOptions struct {
	// Identifiers of additional dependencies.
	// Experimental.
	AdditionalDependencies *[]*string `field:"optional" json:"additionalDependencies" yaml:"additionalDependencies"`
	// The role that needs to be assumed to deploy the stack.
	// Experimental.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The externalID to use with the assumeRoleArn.
	// Experimental.
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
	// Experimental.
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// The role that is passed to CloudFormation to execute the change set.
	// Experimental.
	CloudFormationExecutionRoleArn *string `field:"optional" json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The role to use to look up values from the target AWS account.
	// Experimental.
	LookupRole *cloudassemblyschema.BootstrapRole `field:"optional" json:"lookupRole" yaml:"lookupRole"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion *float64 `field:"optional" json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
	// If the stack template has already been included in the asset manifest, its asset URL.
	// Experimental.
	StackTemplateAssetObjectUrl *string `field:"optional" json:"stackTemplateAssetObjectUrl" yaml:"stackTemplateAssetObjectUrl"`
}

