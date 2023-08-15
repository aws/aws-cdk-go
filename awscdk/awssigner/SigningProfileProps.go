package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a Signing Profile object.
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
type SigningProfileProps struct {
	// The Signing Platform available for signing profile.
	// See: https://docs.aws.amazon.com/signer/latest/developerguide/gs-platform.html
	//
	Platform Platform `field:"required" json:"platform" yaml:"platform"`
	// The validity period for signatures generated using this signing profile.
	// Default: - 135 months.
	//
	SignatureValidity awscdk.Duration `field:"optional" json:"signatureValidity" yaml:"signatureValidity"`
	// Physical name of this Signing Profile.
	// Default: - Assigned by CloudFormation (recommended).
	//
	SigningProfileName *string `field:"optional" json:"signingProfileName" yaml:"signingProfileName"`
}

