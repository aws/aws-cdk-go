package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Username and password combination.
//
// Example:
//   // Build a data source for AppSync to access the database.
//   var api graphqlApi
//   // Create username and password secret for DB Cluster
//   secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
//   	Username: jsii.String("clusteradmin"),
//   })
//
//   // The VPC to place the cluster in
//   vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))
//
//   // Create the serverless cluster, provide all values needed to customise the database.
//   cluster := rds.NewDatabaseCluster(this, jsii.String("AuroraClusterV2"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_15_5(),
//   	}),
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	ClusterIdentifier: jsii.String("db-endpoint-test"),
//   	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
//   	ServerlessV2MinCapacity: jsii.Number(2),
//   	ServerlessV2MaxCapacity: jsii.Number(10),
//   	Vpc: Vpc,
//   	DefaultDatabaseName: jsii.String("demos"),
//   	EnableDataApi: jsii.Boolean(true),
//   })
//   rdsDS := api.AddRdsDataSourceV2(jsii.String("rds"), cluster, secret, jsii.String("demos"))
//
//   // Set up a resolver for an RDS query.
//   rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemosRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "SELECT * FROM demos"
//   	    ]
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
//   	  `)),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemoRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "INSERT INTO demos VALUES (:id, :version)",
//   	      "SELECT * WHERE id = :id"
//   	    ],
//   	    "variableMap": {
//   	      ":id": $util.toJson($util.autoId()),
//   	      ":version": $util.toJson($ctx.args.version)
//   	    }
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
//   	  `)),
//   })
//
type Credentials interface {
	// KMS encryption key to encrypt the generated secret.
	// Default: - default master key.
	//
	EncryptionKey() awskms.IKey
	// The characters to exclude from the generated password.
	//
	// Only used if `password` has not been set.
	// Default: - the DatabaseSecret default exclude character set (" %+~`#$&*()|[]{}:;<>?!'/@\"\\")
	//
	ExcludeCharacters() *string
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Default: - a Secrets Manager generated password.
	//
	Password() awscdk.SecretValue
	// A list of regions where to replicate the generated secret.
	// Default: - Secret is not replicated.
	//
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	// Secret used to instantiate this Login.
	// Default: - none.
	//
	Secret() awssecretsmanager.ISecret
	// The name to use for the Secret if a new Secret is to be generated in SecretsManager for these Credentials.
	// Default: - A name is generated by CloudFormation.
	//
	SecretName() *string
	// Username.
	Username() *string
	// Whether the username should be referenced as a string and not as a dynamic reference to the username in the secret.
	// Default: false.
	//
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


func NewCredentials_Override(c Credentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Credentials",
		nil, // no parameters
		c,
	)
}

// Creates Credentials with a password generated and stored in Secrets Manager.
func Credentials_FromGeneratedSecret(username *string, options *CredentialsBaseOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromGeneratedSecretParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Creates Credentials from a password.
//
// Do not put passwords in your CDK code directly.
func Credentials_FromPassword(username *string, password awscdk.SecretValue) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromPasswordParameters(username, password); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
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
//   ...
//   "username": <required: username>,
//   "password": <required: password>,
// }
// ```.
func Credentials_FromSecret(secret awssecretsmanager.ISecret, username *string) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromSecret",
		[]interface{}{secret, username},
		&returns,
	)

	return returns
}

// Creates Credentials for the given username, and optional password and key.
//
// If no password is provided, one will be generated and stored in Secrets Manager.
func Credentials_FromUsername(username *string, options *CredentialsFromUsernameOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromUsernameParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromUsername",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

