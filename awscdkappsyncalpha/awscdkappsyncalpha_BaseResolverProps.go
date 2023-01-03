// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Basic properties for an AppSync resolver.
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
//   rdsDS.createResolver(jsii.String("QueryGetDemosRdsResolver"), &baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getDemosRds"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"SELECT * FROM demos\"\n    ]\n  }\n  ")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])\n  ")),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.createResolver(jsii.String("MutationAddDemoRdsResolver"), &baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("addDemoRds"),
//   	requestMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n  {\n    \"version\": \"2018-05-29\",\n    \"statements\": [\n      \"INSERT INTO demos VALUES (:id, :version)\",\n      \"SELECT * WHERE id = :id\"\n    ],\n    \"variableMap\": {\n      \":id\": $util.toJson($util.autoId()),\n      \":version\": $util.toJson($ctx.args.version)\n    }\n  }\n  ")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("\n    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])\n  ")),
//   })
//
// Experimental.
type BaseResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	// Experimental.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The maximum number of elements per batch, when using batch invoke.
	// Experimental.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
}

