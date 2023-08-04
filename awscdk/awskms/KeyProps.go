package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a KMS Key object.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   key := kms.NewKey(this, jsii.String("EnvVarEncryptKey"), &KeyProps{
//   	Description: jsii.String("sample key"),
//   })
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   	ActionName: jsii.String("S3Deploy"),
//   	Bucket: targetBucket,
//   	Input: sourceOutput,
//   	EncryptionKey: key,
//   })
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type KeyProps struct {
	// A list of principals to add as key administrators to the key policy.
	//
	// Key administrators have permissions to manage the key (e.g., change permissions, revoke), but do not have permissions
	// to use the key in cryptographic operations (e.g., encrypt, decrypt).
	//
	// These principals will be added to the default key policy (if none specified), or to the specified policy (if provided).
	// Default: [].
	//
	Admins *[]awsiam.IPrincipal `field:"optional" json:"admins" yaml:"admins"`
	// Initial alias to add to the key.
	//
	// More aliases can be added later by calling `addAlias`.
	// Default: - No alias is added for the key.
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A description of the key.
	//
	// Use a description that helps your users decide
	// whether the key is appropriate for a particular task.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the key is available for use.
	// Default: - Key is enabled.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether AWS KMS rotates the key.
	// Default: false.
	//
	EnableKeyRotation *bool `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// The cryptographic configuration of the key. The valid value depends on usage of the key.
	//
	// IMPORTANT: If you change this property of an existing key, the existing key is scheduled for deletion
	// and a new key is created with the specified value.
	// Default: KeySpec.SYMMETRIC_DEFAULT
	//
	KeySpec KeySpec `field:"optional" json:"keySpec" yaml:"keySpec"`
	// The cryptographic operations for which the key can be used.
	//
	// IMPORTANT: If you change this property of an existing key, the existing key is scheduled for deletion
	// and a new key is created with the specified value.
	// Default: KeyUsage.ENCRYPT_DECRYPT
	//
	KeyUsage KeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack.
	//
	// When you remove a customer master key (CMK) from a CloudFormation stack, AWS KMS schedules the CMK for deletion
	// and starts the mandatory waiting period. The PendingWindowInDays property determines the length of waiting period.
	// During the waiting period, the key state of CMK is Pending Deletion, which prevents the CMK from being used in
	// cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the CMK.
	//
	// Enter a value between 7 and 30 days.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
	//
	// Default: - 30 days.
	//
	PendingWindow awscdk.Duration `field:"optional" json:"pendingWindow" yaml:"pendingWindow"`
	// Custom policy document to attach to the KMS key.
	//
	// NOTE - If the `@aws-cdk/aws-kms:defaultKeyPolicies` feature flag is set (the default for new projects),
	// this policy will *override* the default key policy and become the only key policy for the key. If the
	// feature flag is not set, this policy will be appended to the default key policy.
	// Default: - A policy document with permissions for the account root to
	// administer the key will be created.
	//
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// Whether the encryption key should be retained when it is removed from the Stack.
	//
	// This is useful when one wants to
	// retain access to data that was encrypted with a key that is being retired.
	// Default: RemovalPolicy.Retain
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

