package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options for configuring CDN Authorization Configuration.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var channel Channel
//   var mySecret ISecret
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("OriginEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ts(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   	CdnAuth: &CdnAuthConfiguration{
//   		Secrets: []ISecret{
//   			mySecret,
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/cdn-auth.html
//
// Experimental.
type CdnAuthConfiguration struct {
	// Secrets to use for CDN authorization.
	//
	// Each secret must be a JSON object with a `MediaPackageV2CDNIdentifier` key whose
	// value is the CDN-Identifier header value. See the
	// {@link https://docs.aws.amazon.com/mediapackage/latest/userguide/cdn-auth-setup.html MediaPackage CDN authorization docs}.
	// Experimental.
	Secrets *[]awssecretsmanager.ISecret `field:"required" json:"secrets" yaml:"secrets"`
	// Role to use for reading the secrets.
	//
	// If not provided, a role will be created automatically with the required permissions
	// (secretsmanager:GetSecretValue, secretsmanager:DescribeSecret, secretsmanager:BatchGetSecretValue,
	// and kms:Decrypt if the secret uses a customer-managed KMS key).
	// Default: - a new role is created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

