package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options for configuring CDN Authorization Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var secret Secret
//
//   cdnAuthConfiguration := &CdnAuthConfiguration{
//   	Secrets: []ISecret{
//   		secret,
//   	},
//
//   	// the properties below are optional
//   	Role: role,
//   }
//
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/cdn-auth.html
//
// Experimental.
type CdnAuthConfiguration struct {
	// Secrets to use for CDN authorization.
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

