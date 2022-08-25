package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner"
)

// Construction properties for a Code Signing Config object.
//
// Example:
//   import signer "github.com/aws/aws-cdk-go/awscdk"
//
//
//   signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &signingProfileProps{
//   	platform: signer.platform_AWS_LAMBDA_SHA384_ECDSA(),
//   })
//
//   codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &codeSigningConfigProps{
//   	signingProfiles: []iSigningProfile{
//   		signingProfile,
//   	},
//   })
//
//   lambda.NewFunction(this, jsii.String("Function"), &functionProps{
//   	codeSigningConfig: codeSigningConfig,
//   	runtime: lambda.runtime_NODEJS_16_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
type CodeSigningConfigProps struct {
	// List of signing profiles that defines a trusted user who can sign a code package.
	SigningProfiles *[]awssigner.ISigningProfile `field:"required" json:"signingProfiles" yaml:"signingProfiles"`
	// Code signing configuration description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to Enforce, Lambda blocks the deployment request
	// if signature validation checks fail.
	// If you set the policy to Warn, Lambda allows the deployment and
	// creates a CloudWatch log.
	UntrustedArtifactOnDeployment UntrustedArtifactOnDeployment `field:"optional" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

