package awss3


// The transition default minimum object size for lifecycle.
//
// Example:
//   s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	TransitionDefaultMinimumObjectSize: s3.TransitionDefaultMinimumObjectSize_VARIES_BY_STORAGE_CLASS,
//   	LifecycleRules: []lifecycleRule{
//   		&lifecycleRule{
//   			Transitions: []transition{
//   				&transition{
//   					StorageClass: s3.StorageClass_DEEP_ARCHIVE(),
//   					TransitionAfter: awscdk.Duration_Days(jsii.Number(30)),
//   				},
//   			},
//   		},
//   		&lifecycleRule{
//   			ObjectSizeLessThan: jsii.Number(300000),
//   			ObjectSizeGreaterThan: jsii.Number(200000),
//   			Transitions: []*transition{
//   				&transition{
//   					StorageClass: s3.StorageClass_ONE_ZONE_INFREQUENT_ACCESS(),
//   					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),
//   				},
//   			},
//   		},
//   	},
//   })
//
type TransitionDefaultMinimumObjectSize string

const (
	// Objects smaller than 128 KB will not transition to any storage class by default.
	TransitionDefaultMinimumObjectSize_ALL_STORAGE_CLASSES_128_K TransitionDefaultMinimumObjectSize = "ALL_STORAGE_CLASSES_128_K"
	// Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes.
	//
	// By default, all other storage classes will prevent transitions smaller than 128 KB.
	TransitionDefaultMinimumObjectSize_VARIES_BY_STORAGE_CLASS TransitionDefaultMinimumObjectSize = "VARIES_BY_STORAGE_CLASS"
)

