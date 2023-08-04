package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new Aurora Serverless Cluster.
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
type ServerlessClusterProps struct {
	// What kind of database to start.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	// Default: Duration.days(1)
	//
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	// Default: - true.
	//
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Credentials for the administrative user.
	// Default: - A username of 'admin' and SecretsManager-generated password.
	//
	Credentials Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Default: - Database is not created in cluster.
	//
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Default: - true if removalPolicy is RETAIN, false otherwise.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	// Default: false.
	//
	EnableDataApi *bool `field:"optional" json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	// Default: - no parameter group.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Default: - RemovalPolicy.SNAPSHOT (remove the cluster and instances, but retain a snapshot of the data)
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	// Default: - Serverless cluster is automatically paused after 5 minutes of being idle.
	// minimum capacity: 2 ACU
	// maximum capacity: 16 ACU.
	//
	Scaling *ServerlessScalingOptions `field:"optional" json:"scaling" yaml:"scaling"`
	// Security group.
	// Default: - a new security group is created if `vpc` was provided.
	// If the `vpc` property was not provided, no VPC security groups will be associated with the DB cluster.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The KMS key for storage encryption.
	// Default: - the default master key will be used for storage encryption.
	//
	StorageEncryptionKey awskms.IKey `field:"optional" json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	// Default: - a new subnet group is created if `vpc` was provided.
	// If the `vpc` property was not provided, no subnet group will be associated with the DB cluster.
	//
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// The VPC that this Aurora Serverless cluster has been created in.
	// Default: - the default VPC in the account and region will be used.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the instances within the VPC.
	//
	// If provided, the `vpc` property must also be specified.
	// Default: - the VPC default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

