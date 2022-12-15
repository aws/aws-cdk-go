package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Cache options for CodeBuild Project.
//
// A cache can store reusable pieces of your build environment and use them across multiple builds.
//
// Example:
//   var myCachingBucket bucket
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	source: codebuild.source.bitBucket(&bitBucketSourceProps{
//   		owner: jsii.String("awslabs"),
//   		repo: jsii.String("aws-cdk"),
//   	}),
//
//   	cache: codebuild.cache.bucket(myCachingBucket),
//
//   	// BuildSpec with a 'cache' section necessary for S3 caching. This can
//   	// also come from 'buildspec.yml' in your source.
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("..."),
//   				},
//   			},
//   		},
//   		"cache": map[string][]*string{
//   			"paths": []*string{
//   				jsii.String("/root/cachedir/**/*"),
//   			},
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-caching.html
//
// Experimental.
type Cache interface {
}

// The jsii proxy struct for Cache
type jsiiProxy_Cache struct {
	_ byte // padding
}

// Experimental.
func NewCache_Override(c Cache) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codebuild.Cache",
		nil, // no parameters
		c,
	)
}

// Create an S3 caching strategy.
// Experimental.
func Cache_Bucket(bucket awss3.IBucket, options *BucketCacheOptions) Cache {
	_init_.Initialize()

	if err := validateCache_BucketParameters(bucket, options); err != nil {
		panic(err)
	}
	var returns Cache

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.Cache",
		"bucket",
		[]interface{}{bucket, options},
		&returns,
	)

	return returns
}

// Create a local caching strategy.
// Experimental.
func Cache_Local(modes ...LocalCacheMode) Cache {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range modes {
		args = append(args, a)
	}

	var returns Cache

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.Cache",
		"local",
		args,
		&returns,
	)

	return returns
}

// Experimental.
func Cache_None() Cache {
	_init_.Initialize()

	var returns Cache

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.Cache",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

