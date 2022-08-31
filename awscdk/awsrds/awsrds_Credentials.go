package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Username and password combination.
//
// Example:
//   // Build a data source for AppSync to access the database.
//   var api graphqlApi
//   // Create username and password secret for DB Cluster
//   secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &databaseSecretProps{
//   	username: jsii.String("clusteradmin"),
//   })
//
//   // The VPC to place the cluster in
//   vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))
//
//   // Create the serverless cluster, provide all values needed to customise the database.
//   cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &serverlessClusterProps{
//   	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
//   	vpc: vpc,
//   	credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	clusterIdentifier: jsii.String("db-endpoint-test"),
//   	defaultDatabaseName: jsii.String("demos"),
//   })
//   rdsDS := api.addRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))
//
//   // Set up a resolver for an RDS query.
//   rdsDS.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemosRds"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"SELECT * FROM demos\"\n    ]\n  }\n  ")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])\n  ")),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("addDemoRds"),
//   	requestMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"INSERT INTO demos VALUES (:id, :version)\",\n      \"SELECT * WHERE id = :id\"\n    ],\n    \"variableMap\": {\n      \":id\": $util.toJson($util.autoId()),\n      \":version\": $util.toJson($ctx.args.version)\n    }\n  }\n  ")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])\n  ")),
//   })
//
// Experimental.
type Credentials interface {
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The characters to exclude from the generated password.
	//
	// Only used if {@link password} has not been set.
	// Experimental.
	ExcludeCharacters() *string
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	Password() awscdk.SecretValue
	// A list of regions where to replicate the generated secret.
	// Experimental.
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	// Secret used to instantiate this Login.
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// The name to use for the Secret if a new Secret is to be generated in SecretsManager for these Credentials.
	// Experimental.
	SecretName() *string
	// Username.
	// Experimental.
	Username() *string
	// Whether the username should be referenced as a string and not as a dynamic reference to the username in the secret.
	// Experimental.
	UsernameAsString() *bool
}

// The jsii proxy struct for Credentials
type jsiiProxy_Credentials struct {
	_ byte // padding
}

func (j *jsiiProxy_Credentials) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion {
	var returns *[]*awssecretsmanager.ReplicaRegion
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) UsernameAsString() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"usernameAsString",
		&returns,
	)
	return returns
}


// Experimental.
func NewCredentials_Override(c Credentials) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.Credentials",
		nil, // no parameters
		c,
	)
}

// Creates Credentials with a password generated and stored in Secrets Manager.
// Experimental.
func Credentials_FromGeneratedSecret(username *string, options *CredentialsBaseOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromGeneratedSecretParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.Credentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Creates Credentials from a password.
//
// Do not put passwords in your CDK code directly.
// Experimental.
func Credentials_FromPassword(username *string, password awscdk.SecretValue) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromPasswordParameters(username, password); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.Credentials",
		"fromPassword",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Creates Credentials from an existing Secrets Manager ``Secret`` (or ``DatabaseSecret``).
//
// The Secret must be a JSON string with a ``username`` and ``password`` field:
// ```
// {
//    ...
//    "username": <required: username>,
//    "password": <required: password>,
// }
// ```.
// Experimental.
func Credentials_FromSecret(secret awssecretsmanager.ISecret, username *string) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.Credentials",
		"fromSecret",
		[]interface{}{secret, username},
		&returns,
	)

	return returns
}

// Creates Credentials for the given username, and optional password and key.
//
// If no password is provided, one will be generated and stored in Secrets Manager.
// Experimental.
func Credentials_FromUsername(username *string, options *CredentialsFromUsernameOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromUsernameParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.Credentials",
		"fromUsername",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

