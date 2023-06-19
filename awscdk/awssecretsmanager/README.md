# AWS Secrets Manager Construct Library

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
```

## Create a new Secret in a Stack

In order to have SecretsManager generate a new secret value automatically,
you can get started with the following:

```go
// Default secret
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
// Using the default secret
// Using the default secret
iam.NewUser(this, jsii.String("User"), &UserProps{
	Password: secret.secretValue,
})
// Templated secret
templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &SecretProps{
	GenerateSecretString: &SecretStringGenerator{
		SecretStringTemplate: jSON.stringify(map[string]*string{
			"username": jsii.String("user"),
		}),
		GenerateStringKey: jsii.String("password"),
	},
})
// Using the templated secret
// Using the templated secret
iam.NewUser(this, jsii.String("OtherUser"), &UserProps{
	UserName: templatedSecret.secretValueFromJson(jsii.String("username")).ToString(),
	Password: templatedSecret.secretValueFromJson(jsii.String("password")),
})
```

If you need to use a pre-existing secret, the recommended way is to manually
provision the secret in *AWS SecretsManager* and use the `Secret.fromSecretArn`
or `Secret.fromSecretAttributes` method to make it available in your CDK Application:

```go
var encryptionKey key

secret := secretsmanager.Secret_FromSecretAttributes(this, jsii.String("ImportedSecret"), &SecretAttributes{
	SecretArn: jsii.String("arn:aws:secretsmanager:<region>:<account-id-number>:secret:<secret-name>-<random-6-characters>"),
	// If the secret is encrypted using a KMS-hosted CMK, either import or reference that key:
	EncryptionKey: EncryptionKey,
})
```

SecretsManager secret values can only be used in select set of properties. For the
list of properties, see [the CloudFormation Dynamic References documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html).

A secret can set `RemovalPolicy`. If it set to `RETAIN`, that removing a secret will fail.

## Grant permission to use the secret to a role

You must grant permission to a resource for that resource to be allowed to
use a secret. This can be achieved with the `Secret.grantRead` and/or `Secret.grantWrite`
method, depending on your need:

```go
role := iam.NewRole(this, jsii.String("SomeRole"), &RoleProps{
	AssumedBy: iam.NewAccountRootPrincipal(),
})
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
secret.grantRead(role)
secret.grantWrite(role)
```

If, as in the following example, your secret was created with a KMS key:

```go
var role role

key := kms.NewKey(this, jsii.String("KMS"))
secret := secretsmanager.NewSecret(this, jsii.String("Secret"), &SecretProps{
	EncryptionKey: key,
})
secret.grantRead(role)
secret.grantWrite(role)
```

then `Secret.grantRead` and `Secret.grantWrite` will also grant the role the
relevant encrypt and decrypt permissions to the KMS key through the
SecretsManager service principal.

The principal is automatically added to Secret resource policy and KMS Key policy for cross account access:

```go
otherAccount := iam.NewAccountPrincipal(jsii.String("1234"))
key := kms.NewKey(this, jsii.String("KMS"))
secret := secretsmanager.NewSecret(this, jsii.String("Secret"), &SecretProps{
	EncryptionKey: key,
})
secret.grantRead(otherAccount)
```

## Rotating a Secret

### Using a Custom Lambda Function

A rotation schedule can be added to a Secret using a custom Lambda function:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn function

secret := secretsmanager.NewSecret(this, jsii.String("Secret"))

secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
	RotationLambda: fn,
	AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(15)),
})
```

Note: The required permissions for Lambda to call SecretsManager and the other way round are automatically granted based on [AWS Documentation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions.html) as long as the Lambda is not imported.

See [Overview of the Lambda Rotation Function](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-lambda-function-overview.html) on how to implement a Lambda Rotation Function.

### Using a Hosted Lambda Function

Use the `hostedRotation` prop to rotate a secret with a hosted Lambda function:

```go
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))

secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
	HostedRotation: secretsmanager.HostedRotation_MysqlSingleUser(),
})
```

Hosted rotation is available for secrets representing credentials for MySQL, PostgreSQL, Oracle,
MariaDB, SQLServer, Redshift and MongoDB (both for the single and multi user schemes).

When deployed in a VPC, the hosted rotation implements `ec2.IConnectable`:

```go
var myVpc vpc
var dbConnections connections
var secret secret


myHostedRotation := secretsmanager.HostedRotation_MysqlSingleUser(&SingleUserHostedRotationOptions{
	Vpc: myVpc,
})
secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
	HostedRotation: myHostedRotation,
})
dbConnections.AllowDefaultPortFrom(myHostedRotation)
```

See also [Automating secret creation in AWS CloudFormation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/integrating_cloudformation.html).

## Rotating database credentials

Define a `SecretRotation` to rotate database credentials:

```go
var mySecret secret
var myDatabase iConnectable
var myVpc vpc


secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &SecretRotationProps{
	Application: secretsmanager.SecretRotationApplication_MYSQL_ROTATION_SINGLE_USER(),
	 // MySQL single user scheme
	Secret: mySecret,
	Target: myDatabase,
	 // a Connectable
	Vpc: myVpc,
	 // The VPC where the secret rotation application will be deployed
	ExcludeCharacters: jsii.String(" %+:;{}"),
})
```

The secret must be a JSON string with the following format:

```json
{
  "engine": "<required: database engine>",
  "host": "<required: instance host name>",
  "username": "<required: username>",
  "password": "<required: password>",
  "dbname": "<optional: database name>",
  "port": "<optional: if not specified, default port will be used>",
  "masterarn": "<required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>"
}
```

For the multi user scheme, a `masterSecret` must be specified:

```go
var myUserSecret secret
var myMasterSecret secret
var myDatabase iConnectable
var myVpc vpc


secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &SecretRotationProps{
	Application: secretsmanager.SecretRotationApplication_MYSQL_ROTATION_MULTI_USER(),
	Secret: myUserSecret,
	 // The secret that will be rotated
	MasterSecret: myMasterSecret,
	 // The secret used for the rotation
	Target: myDatabase,
	Vpc: myVpc,
})
```

See also [aws-rds](https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk/aws-rds/README.md) where
credentials generation and rotation is integrated.

## Importing Secrets

Existing secrets can be imported by ARN, name, and other attributes (including the KMS key used to encrypt the secret).
Secrets imported by name should use the short-form of the name (without the SecretsManager-provided suffx);
the secret name must exist in the same account and region as the stack.
Importing by name makes it easier to reference secrets created in different regions, each with their own suffix and ARN.

```go
secretCompleteArn := "arn:aws:secretsmanager:eu-west-1:111111111111:secret:MySecret-f3gDy9"
secretPartialArn := "arn:aws:secretsmanager:eu-west-1:111111111111:secret:MySecret" // No Secrets Manager suffix
encryptionKey := kms.Key_FromKeyArn(this, jsii.String("MyEncKey"), jsii.String("arn:aws:kms:eu-west-1:111111111111:key/21c4b39b-fde2-4273-9ac0-d9bb5c0d0030"))
mySecretFromCompleteArn := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("SecretFromCompleteArn"), secretCompleteArn)
mySecretFromPartialArn := secretsmanager.Secret_FromSecretPartialArn(this, jsii.String("SecretFromPartialArn"), secretPartialArn)
mySecretFromName := secretsmanager.Secret_FromSecretNameV2(this, jsii.String("SecretFromName"), jsii.String("MySecret"))
mySecretFromAttrs := secretsmanager.Secret_FromSecretAttributes(this, jsii.String("SecretFromAttributes"), &SecretAttributes{
	SecretCompleteArn: jsii.String(SecretCompleteArn),
	EncryptionKey: EncryptionKey,
})
```

## Replicating secrets

Secrets can be replicated to multiple regions by specifying `replicaRegions`:

```go
var myKey key

secretsmanager.NewSecret(this, jsii.String("Secret"), &SecretProps{
	ReplicaRegions: []replicaRegion{
		&replicaRegion{
			Region: jsii.String("eu-west-1"),
		},
		&replicaRegion{
			Region: jsii.String("eu-central-1"),
			EncryptionKey: myKey,
		},
	},
})
```

Alternatively, use `addReplicaRegion()`:

```go
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
secret.AddReplicaRegion(jsii.String("eu-west-1"))
```
