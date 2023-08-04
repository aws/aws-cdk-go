package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner"
)

// Construction properties for a Code Signing Config object.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &SigningProfileProps{
//   	Platform: signer.Platform_AWS_LAMBDA_SHA384_ECDSA(),
//   })
//
//   codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &CodeSigningConfigProps{
//   	SigningProfiles: []iSigningProfile{
//   		signingProfile,
//   	},
//   })
//
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	CodeSigningConfig: CodeSigningConfig,
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
type CodeSigningConfigProps struct {
	// List of signing profiles that defines a trusted user who can sign a code package.
	SigningProfiles *[]awssigner.ISigningProfile `field:"required" json:"signingProfiles" yaml:"signingProfiles"`
	// Code signing configuration description.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to Enforce, Lambda blocks the deployment request
	// if signature validation checks fail.
	// If you set the policy to Warn, Lambda allows the deployment and
	// creates a CloudWatch log.
	// Default: UntrustedArtifactOnDeployment.WARN
	//
	UntrustedArtifactOnDeployment UntrustedArtifactOnDeployment `field:"optional" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

