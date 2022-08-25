package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
type SigningProfileProps struct {
	// The Signing Platform available for signing profile.
	// See: https://docs.aws.amazon.com/signer/latest/developerguide/gs-platform.html
	//
	Platform Platform `field:"required" json:"platform" yaml:"platform"`
	// The validity period for signatures generated using this signing profile.
	SignatureValidity awscdk.Duration `field:"optional" json:"signatureValidity" yaml:"signatureValidity"`
	// Physical name of this Signing Profile.
	SigningProfileName *string `field:"optional" json:"signingProfileName" yaml:"signingProfileName"`
}

