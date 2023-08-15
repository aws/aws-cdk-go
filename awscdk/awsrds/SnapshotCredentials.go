package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Credentials to update the password for a ``DatabaseInstanceFromSnapshot``.
//
// Example:
//   var vpc vpc
//
//   engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   	Version: rds.PostgresEngineVersion_VER_15_2(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("InstanceFromSnapshotWithCustomizedSecret"), &DatabaseInstanceFromSnapshotProps{
//   	Engine: Engine,
//   	Vpc: Vpc,
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   	Credentials: rds.SnapshotCredentials_FromGeneratedSecret(jsii.String("username"), &SnapshotCredentialsFromGeneratedPasswordOptions{
//   		EncryptionKey: myKey,
//   		ExcludeCharacters: jsii.String("!&*^#@()"),
//   		ReplicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-2"),
//   			},
//   		},
//   	}),
//   })
//
type SnapshotCredentials interface {
	// KMS encryption key to encrypt the generated secret.
	// Default: - default master key.
	//
	EncryptionKey() awskms.IKey
	// The characters to exclude from the generated password.
	//
	// Only used if `generatePassword` if true.
	// Default: - the DatabaseSecret default exclude character set (" %+~`#$&*()|[]{}:;<>?!'/@\"\\")
	//
	ExcludeCharacters() *string
	// Whether a new password should be generated.
	GeneratePassword() *bool
	// The master user password.
	//
	// Do not put passwords in your CDK code directly.
	// Default: - the existing password from the snapshot.
	//
	Password() awscdk.SecretValue
	// Whether to replace the generated secret when the criteria for the password change.
	// Default: false.
	//
	ReplaceOnPasswordCriteriaChanges() *bool
	// A list of regions where to replicate the generated secret.
	// Default: - Secret is not replicated.
	//
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	// Secret used to instantiate this Login.
	// Default: - none.
	//
	Secret() awssecretsmanager.ISecret
	// The master user name.
	//
	// Must be the **current** master user name of the snapshot.
	// It is not possible to change the master user name of a RDS instance.
	// Default: - the existing username from the snapshot.
	//
	Username() *string
}

// The jsii proxy struct for SnapshotCredentials
type jsiiProxy_SnapshotCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_SnapshotCredentials) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) GeneratePassword() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"generatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ReplaceOnPasswordCriteriaChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceOnPasswordCriteriaChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion {
	var returns *[]*awssecretsmanager.ReplicaRegion
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnapshotCredentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


func NewSnapshotCredentials_Override(s SnapshotCredentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		nil, // no parameters
		s,
	)
}

// Generate a new password for the snapshot, using the existing username and an optional encryption key.
//
// Note - The username must match the existing master username of the snapshot.
//
// NOTE: use `fromGeneratedSecret()` for new Clusters and Instances. Switching from
// `fromGeneratedPassword()` to `fromGeneratedSecret()` for already deployed Clusters
// or Instances will update their master password.
func SnapshotCredentials_FromGeneratedPassword(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) SnapshotCredentials {
	_init_.Initialize()

	if err := validateSnapshotCredentials_FromGeneratedPasswordParameters(username, options); err != nil {
		panic(err)
	}
	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromGeneratedPassword",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Generate a new password for the snapshot, using the existing username and an optional encryption key.
//
// The new credentials are stored in Secrets Manager.
//
// Note - The username must match the existing master username of the snapshot.
func SnapshotCredentials_FromGeneratedSecret(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) SnapshotCredentials {
	_init_.Initialize()

	if err := validateSnapshotCredentials_FromGeneratedSecretParameters(username, options); err != nil {
		panic(err)
	}
	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Update the snapshot login with an existing password.
func SnapshotCredentials_FromPassword(password awscdk.SecretValue) SnapshotCredentials {
	_init_.Initialize()

	if err := validateSnapshotCredentials_FromPasswordParameters(password); err != nil {
		panic(err)
	}
	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromPassword",
		[]interface{}{password},
		&returns,
	)

	return returns
}

// Update the snapshot login with an existing password from a Secret.
//
// The Secret must be a JSON string with a ``password`` field:
// ```
// {
//   ...
//   "password": <required: password>,
// }
// ```.
func SnapshotCredentials_FromSecret(secret awssecretsmanager.ISecret) SnapshotCredentials {
	_init_.Initialize()

	if err := validateSnapshotCredentials_FromSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns SnapshotCredentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.SnapshotCredentials",
		"fromSecret",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

