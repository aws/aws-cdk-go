package awscodebuild


// Example:
//   var sourceBucket Bucket
//   var myCachingBucket Bucket
//
//
//   codebuild.NewProject(this, jsii.String("ProjectA"), &ProjectProps{
//   	Source: codebuild.Source_S3(&S3SourceProps{
//   		Bucket: sourceBucket,
//   		Path: jsii.String("path/to/source-a.zip"),
//   	}),
//   	// configure the same bucket and path prefix
//   	Cache: codebuild.Cache_Bucket(myCachingBucket, &BucketCacheOptions{
//   		Prefix: jsii.String("cache"),
//   		// use the same cache namespace
//   		CacheNamespace: jsii.String("cache-namespace"),
//   	}),
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("..."),
//   				},
//   			},
//   		},
//   		// specify the same cache key and paths
//   		"cache": map[string]interface{}{
//   			"key": jsii.String("unique-key"),
//   			"paths": []*string{
//   				jsii.String("/root/cachedir/**/*"),
//   			},
//   		},
//   	}),
//   })
//
//   codebuild.NewProject(this, jsii.String("ProjectB"), &ProjectProps{
//   	Source: codebuild.Source_*S3(&S3SourceProps{
//   		Bucket: sourceBucket,
//   		Path: jsii.String("path/to/source-b.zip"),
//   	}),
//   	// configure the same bucket and path prefix
//   	Cache: codebuild.Cache_*Bucket(myCachingBucket, &BucketCacheOptions{
//   		Prefix: jsii.String("cache"),
//   		// use the same cache namespace
//   		CacheNamespace: jsii.String("cache-namespace"),
//   	}),
//   	BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("..."),
//   				},
//   			},
//   		},
//   		// specify the same cache key and paths
//   		"cache": map[string]interface{}{
//   			"key": jsii.String("unique-key"),
//   			"paths": []*string{
//   				jsii.String("/root/cachedir/**/*"),
//   			},
//   		},
//   	}),
//   })
//
type BucketCacheOptions struct {
	// Defines the scope of the cache.
	//
	// You can use this namespace to share a cache across multiple projects.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/caching-s3.html#caching-s3-sharing
	//
	// Default: undefined - No cache namespace, which means that the cache is not shared across multiple projects.
	//
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The prefix to use to store the cache in the bucket.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

