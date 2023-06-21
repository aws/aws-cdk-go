package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Storage class to move an object to.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	LifecycleRules: []lifecycleRule{
//   		&lifecycleRule{
//   			AbortIncompleteMultipartUploadAfter: awscdk.Duration_Minutes(jsii.Number(30)),
//   			Enabled: jsii.Boolean(false),
//   			Expiration: awscdk.Duration_Days(jsii.Number(30)),
//   			ExpirationDate: NewDate(),
//   			ExpiredObjectDeleteMarker: jsii.Boolean(false),
//   			Id: jsii.String("id"),
//   			NoncurrentVersionExpiration: awscdk.Duration_*Days(jsii.Number(30)),
//
//   			// the properties below are optional
//   			NoncurrentVersionsToRetain: jsii.Number(123),
//   			NoncurrentVersionTransitions: []noncurrentVersionTransition{
//   				&noncurrentVersionTransition{
//   					StorageClass: s3.StorageClass_GLACIER(),
//   					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),
//
//   					// the properties below are optional
//   					NoncurrentVersionsToRetain: jsii.Number(123),
//   				},
//   			},
//   			ObjectSizeGreaterThan: jsii.Number(500),
//   			Prefix: jsii.String("prefix"),
//   			ObjectSizeLessThan: jsii.Number(10000),
//   			Transitions: []transition{
//   				&transition{
//   					StorageClass: s3.StorageClass_GLACIER(),
//
//   					// the properties below are optional
//   					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),
//   					TransitionDate: NewDate(),
//   				},
//   			},
//   		},
//   	},
//   })
//
type StorageClass interface {
	Value() *string
	ToString() *string
}

// The jsii proxy struct for StorageClass
type jsiiProxy_StorageClass struct {
	_ byte // padding
}

func (j *jsiiProxy_StorageClass) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewStorageClass(value *string) StorageClass {
	_init_.Initialize()

	if err := validateNewStorageClassParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageClass{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.StorageClass",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewStorageClass_Override(s StorageClass, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.StorageClass",
		[]interface{}{value},
		s,
	)
}

func StorageClass_DEEP_ARCHIVE() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"DEEP_ARCHIVE",
		&returns,
	)
	return returns
}

func StorageClass_GLACIER() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"GLACIER",
		&returns,
	)
	return returns
}

func StorageClass_GLACIER_INSTANT_RETRIEVAL() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"GLACIER_INSTANT_RETRIEVAL",
		&returns,
	)
	return returns
}

func StorageClass_INFREQUENT_ACCESS() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"INFREQUENT_ACCESS",
		&returns,
	)
	return returns
}

func StorageClass_INTELLIGENT_TIERING() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"INTELLIGENT_TIERING",
		&returns,
	)
	return returns
}

func StorageClass_ONE_ZONE_INFREQUENT_ACCESS() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"ONE_ZONE_INFREQUENT_ACCESS",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageClass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

