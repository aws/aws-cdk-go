package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Construction properties for a DatabaseSecret.
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
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	// Experimental.
	MasterSecret awssecretsmanager.ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// Whether to replace this secret when the criteria for the password change.
	//
	// This is achieved by overriding the logical id of the AWS::SecretsManager::Secret
	// with a hash of the options that influence the password generation. This
	// way a new secret will be created when the password is regenerated and the
	// cluster or instance consuming this secret will have its credentials updated.
	// Experimental.
	ReplaceOnPasswordCriteriaChanges *bool `field:"optional" json:"replaceOnPasswordCriteriaChanges" yaml:"replaceOnPasswordCriteriaChanges"`
	// A list of regions where to replicate this secret.
	// Experimental.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	// Experimental.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

