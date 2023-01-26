# AWS Secrets Manager Construct Library

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
```

## Create a new Secret in a Stack

To have SecretsManager generate a new secret value automatically,
follow this example:

```go
var vpc vpc


// Simple secret
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
// Using the secret
instance1 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance1"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine_POSTGRES(),
	credentials: rds.credentials.fromSecret(secret),
	vpc: vpc,
})
// Templated secret with username and password fields
templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &secretProps{
	generateSecretString: &secretStringGenerator{
		secretStringTemplate: jSON.stringify(map[string]*string{
			"username": jsii.String("postgres"),
		}),
		generateStringKey: jsii.String("password"),
	},
})
// Using the templated secret as credentials
instance2 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance2"), &databaseInstanceProps{
	engine: rds.*databaseInstanceEngine_POSTGRES(),
	credentials: map[string]interface{}{
		"username": templatedSecret.secretValueFromJson(jsii.String("username")).toString(),
		"password": templatedSecret.secretValueFromJson(jsii.String("password")),
	},
	vpc: vpc,
})
```

If you need to use a pre-existing secret, the recommended way is to manually
provision the secret in *AWS SecretsManager* and use the `Secret.fromSecretArn`
or `Secret.fromSecretAttributes` method to make it available in your CDK Application:

```go
var encryptionKey key

secret := secretsmanager.secret.fromSecretAttributes(this, jsii.String("ImportedSecret"), &secretAttributes{
	secretArn: jsii.String("arn:aws:secretsmanager:<region>:<account-id-number>:secret:<secret-name>-<random-6-characters>"),
	// If the secret is encrypted using a KMS-hosted CMK, either import or reference that key:
	encryptionKey: encryptionKey,
})
```

SecretsManager secret values can only be used in select set of properties. For the
list of properties, see [the CloudFormation Dynamic References documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html).

A secret can set `RemovalPolicy`. If it set to `RETAIN`, removing that secret will fail.

## Grant permission to use the secret to a role

You must grant permission to a resource for that resource to be allowed to
use a secret. This can be achieved with the `Secret.grantRead` and/or `Secret.grantWrite`
method, depending on your need:

```go
role := iam.NewRole(this, jsii.String("SomeRole"), &roleProps{
	assumedBy: iam.NewAccountRootPrincipal(),
})
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
secret.grantRead(role)
secret.grantWrite(role)
```

If, as in the following example, your secret was created with a KMS key:

```go
var role role

key := kms.NewKey(this, jsii.String("KMS"))
secret := secretsmanager.NewSecret(this, jsii.String("Secret"), &secretProps{
	encryptionKey: key,
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
secret := secretsmanager.NewSecret(this, jsii.String("Secret"), &secretProps{
	encryptionKey: key,
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

secret.addRotationSchedule(jsii.String("RotationSchedule"), &rotationScheduleOptions{
	rotationLambda: fn,
	automaticallyAfter: awscdk.Duration.days(jsii.Number(15)),
})
```

Note: The required permissions for Lambda to call SecretsManager and the other way round are automatically granted based on [AWS Documentation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions.html) as long as the Lambda is not imported.

See [Overview of the Lambda Rotation Function](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-lambda-function-overview.html) on how to implement a Lambda Rotation Function.

### Using a Hosted Lambda Function

Use the `hostedRotation` prop to rotate a secret with a hosted Lambda function:

```go
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))

secret.addRotationSchedule(jsii.String("RotationSchedule"), &rotationScheduleOptions{
	hostedRotation: secretsmanager.hostedRotation.mysqlSingleUser(),
})
```

Hosted rotation is available for secrets representing credentials for MySQL, PostgreSQL, Oracle,
MariaDB, SQLServer, Redshift and MongoDB (both for the single and multi user schemes).

When deployed in a VPC, the hosted rotation implements `ec2.IConnectable`:

```go
var myVpc vpc
var dbConnections connections
var secret secret


myHostedRotation := secretsmanager.hostedRotation.mysqlSingleUser(&singleUserHostedRotationOptions{
	vpc: myVpc,
})
secret.addRotationSchedule(jsii.String("RotationSchedule"), &rotationScheduleOptions{
	hostedRotation: myHostedRotation,
})
dbConnections.allowDefaultPortFrom(myHostedRotation)
```

Use the `excludeCharacters` option to customize the characters excluded from
the generated password when it is rotated. By default, the rotation excludes
the same characters as the ones excluded for the secret. If none are defined
then the following set is used: `% +~`#$&*()|[]{}:;<>?!'/@"\`.

See also [Automating secret creation in AWS CloudFormation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/integrating_cloudformation.html).

## Rotating database credentials

Define a `SecretRotation` to rotate database credentials:

```go
var mySecret secret
var myDatabase iConnectable
var myVpc vpc


secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &secretRotationProps{
	application: secretsmanager.secretRotationApplication_MYSQL_ROTATION_SINGLE_USER(),
	 // MySQL single user scheme
	secret: mySecret,
	target: myDatabase,
	 // a Connectable
	vpc: myVpc,
	 // The VPC where the secret rotation application will be deployed
	excludeCharacters: jsii.String(" %+:;{}"),
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


secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &secretRotationProps{
	application: secretsmanager.secretRotationApplication_MYSQL_ROTATION_MULTI_USER(),
	secret: myUserSecret,
	 // The secret that will be rotated
	masterSecret: myMasterSecret,
	 // The secret used for the rotation
	target: myDatabase,
	vpc: myVpc,
})
```

See also [aws-rds](https://github.com/aws/aws-cdk/blob/main/packages/%40aws-cdk/aws-rds/README.md) where
credentials generation and rotation is integrated.

## Importing Secrets

Existing secrets can be imported by ARN, name, and other attributes (including the KMS key used to encrypt the secret).
Secrets imported by name should use the short-form of the name (without the SecretsManager-provided suffix);
the secret name must exist in the same account and region as the stack.
Importing by name makes it easier to reference secrets created in different regions, each with their own suffix and ARN.

```go
secretCompleteArn := "arn:aws:secretsmanager:eu-west-1:111111111111:secret:MySecret-f3gDy9"
secretPartialArn := "arn:aws:secretsmanager:eu-west-1:111111111111:secret:MySecret" // No Secrets Manager suffix
encryptionKey := kms.key.fromKeyArn(this, jsii.String("MyEncKey"), jsii.String("arn:aws:kms:eu-west-1:111111111111:key/21c4b39b-fde2-4273-9ac0-d9bb5c0d0030"))
mySecretFromCompleteArn := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("SecretFromCompleteArn"), secretCompleteArn)
mySecretFromPartialArn := secretsmanager.secret.fromSecretPartialArn(this, jsii.String("SecretFromPartialArn"), secretPartialArn)
mySecretFromName := secretsmanager.secret.fromSecretNameV2(this, jsii.String("SecretFromName"), jsii.String("MySecret"))
mySecretFromAttrs := secretsmanager.secret.fromSecretAttributes(this, jsii.String("SecretFromAttributes"), &secretAttributes{
	secretCompleteArn: jsii.String(secretCompleteArn),
	encryptionKey: encryptionKey,
})
```

## Replicating secrets

Secrets can be replicated to multiple regions by specifying `replicaRegions`:

```go
var myKey key

secretsmanager.NewSecret(this, jsii.String("Secret"), &secretProps{
	replicaRegions: []replicaRegion{
		&replicaRegion{
			region: jsii.String("eu-west-1"),
		},
		&replicaRegion{
			region: jsii.String("eu-central-1"),
			encryptionKey: myKey,
		},
	},
})
```

Alternatively, use `addReplicaRegion()`:

```go
secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
secret.addReplicaRegion(jsii.String("eu-west-1"))
```

## Creating JSON Secrets

Sometimes it is necessary to create a secret in SecretsManager that contains a JSON object.
For example:

```json
{
  "username": "myUsername",
  "database": "foo",
  "password": "mypassword"
}
```

In order to create this type of secret, use the `secretObjectValue` input prop.

```go
var stack stack
user := iam.NewUser(stack, jsii.String("User"))
accessKey := iam.NewAccessKey(stack, jsii.String("AccessKey"), &accessKeyProps{
	user: user,
})

secretsmanager.NewSecret(stack, jsii.String("Secret"), &secretProps{
	secretObjectValue: map[string]secretValue{
		"username": awscdk.SecretValue.unsafePlainText(user.userName),
		"database": awscdk.SecretValue.unsafePlainText(jsii.String("foo")),
		"password": accessKey.secretAccessKey,
	},
})
```

In this case both the `username` and `database` are not a `Secret` so `SecretValue.unsafePlainText` needs to be used.
This means that they will be rendered as plain text in the template, but in this case neither of those
are actual "secrets".
