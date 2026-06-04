package awscdkelasticachealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Supported cache engines together with available versions.
//
// Named instances cover the versions currently available on ElastiCache Serverless.
// To target a version that is not yet represented by a named instance, use
// `CacheEngine.of(engineType, majorEngineVersion)`.
//
// Example:
//   var vpc Vpc
//
//
//   serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
//   	Engine: elasticache.CacheEngine_VALKEY_LATEST(),
//   	Backup: &BackupSettings{
//   		// set a backup name before deleting a cache
//   		BackupNameBeforeDeletion: jsii.String("my-final-backup-name"),
//   	},
//   	Vpc: Vpc,
//   })
//
// See: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html
//
// Experimental.
type CacheEngine interface {
	// The engine type, for example `'valkey'`, `'redis'`, or `'memcached'`.
	//
	// Maps directly to the `Engine` property of `AWS::ElastiCache::ServerlessCache`.
	// Experimental.
	EngineType() *string
	// The major engine version, for example `'9'` or `'1.6'`. Maps directly to the `MajorEngineVersion` property of `AWS::ElastiCache::ServerlessCache`. When `undefined`, the service selects the latest major version automatically.
	// Experimental.
	MajorEngineVersion() *string
	// Returns a string representation of this cache engine, for logging and error messages.
	//
	// The format is `engineType_majorEngineVersion` when a
	// major version is set, or just `engineType` otherwise (for example,
	// `'valkey_8'`, `'memcached_1.6'`, `'redis'`).
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CacheEngine
type jsiiProxy_CacheEngine struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheEngine) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheEngine) MajorEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorEngineVersion",
		&returns,
	)
	return returns
}


// Create a new `CacheEngine` with an arbitrary engine type and major version.
//
// Use this for engine/version combinations that are not yet represented by a
// named static member.
// Experimental.
func CacheEngine_Of(engineType *string, majorEngineVersion *string) CacheEngine {
	_init_.Initialize()

	if err := validateCacheEngine_OfParameters(engineType); err != nil {
		panic(err)
	}
	var returns CacheEngine

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"of",
		[]interface{}{engineType, majorEngineVersion},
		&returns,
	)

	return returns
}

func CacheEngine_MEMCACHED_1_6() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"MEMCACHED_1_6",
		&returns,
	)
	return returns
}

func CacheEngine_MEMCACHED_LATEST() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"MEMCACHED_LATEST",
		&returns,
	)
	return returns
}

func CacheEngine_REDIS_7() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"REDIS_7",
		&returns,
	)
	return returns
}

func CacheEngine_REDIS_LATEST() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"REDIS_LATEST",
		&returns,
	)
	return returns
}

func CacheEngine_VALKEY_7() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"VALKEY_7",
		&returns,
	)
	return returns
}

func CacheEngine_VALKEY_8() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"VALKEY_8",
		&returns,
	)
	return returns
}

func CacheEngine_VALKEY_9() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"VALKEY_9",
		&returns,
	)
	return returns
}

func CacheEngine_VALKEY_LATEST() CacheEngine {
	_init_.Initialize()
	var returns CacheEngine
	_jsii_.StaticGet(
		"@aws-cdk/aws-elasticache-alpha.CacheEngine",
		"VALKEY_LATEST",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CacheEngine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

