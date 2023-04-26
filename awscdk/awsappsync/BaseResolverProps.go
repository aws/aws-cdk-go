package awsappsync


// Basic properties for an AppSync resolver.
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
//   cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	ClusterIdentifier: jsii.String("db-endpoint-test"),
//   	DefaultDatabaseName: jsii.String("demos"),
//   })
//   rdsDS := api.AddRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))
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
type BaseResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The function code.
	Code Code `field:"optional" json:"code" yaml:"code"`
	// The maximum number of elements per batch, when using batch invoke.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

