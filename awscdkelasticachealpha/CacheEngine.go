package awscdkelasticachealpha


// Supported cache engines together with available versions.
//
// Example:
//   var vpc Vpc
//
//
//   serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
//   	Engine: elasticache.CacheEngine_VALKEY_LATEST,
//   	Backup: &BackupSettings{
//   		// set a backup name before deleting a cache
//   		BackupNameBeforeDeletion: jsii.String("my-final-backup-name"),
//   	},
//   	Vpc: Vpc,
//   })
//
// Experimental.
type CacheEngine string

const (
	// Valkey engine, latest major version available, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_VALKEY_LATEST CacheEngine = "VALKEY_LATEST"
	// Valkey engine, major version 7, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_VALKEY_7 CacheEngine = "VALKEY_7"
	// Valkey engine, major version 8, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_VALKEY_8 CacheEngine = "VALKEY_8"
	// Redis engine, latest major version available, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_REDIS_LATEST CacheEngine = "REDIS_LATEST"
	// Redis engine, major version 7, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_REDIS_7 CacheEngine = "REDIS_7"
	// Memcached engine, latest major version available, minor version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_MEMCACHED_LATEST CacheEngine = "MEMCACHED_LATEST"
	// Memcached engine, minor version 1.6, patch version is selected automatically For more information about the features related to this version check: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html.
	// Experimental.
	CacheEngine_MEMCACHED_1_6 CacheEngine = "MEMCACHED_1_6"
)

