package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a Signing Profile object.
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
// Experimental.
type SigningProfileProps struct {
	// The Signing Platform available for signing profile.
	// See: https://docs.aws.amazon.com/signer/latest/developerguide/gs-platform.html
	//
	// Experimental.
	Platform Platform `field:"required" json:"platform" yaml:"platform"`
	// The validity period for signatures generated using this signing profile.
	// Experimental.
	SignatureValidity awscdk.Duration `field:"optional" json:"signatureValidity" yaml:"signatureValidity"`
	// Physical name of this Signing Profile.
	// Experimental.
	SigningProfileName *string `field:"optional" json:"signingProfileName" yaml:"signingProfileName"`
}

