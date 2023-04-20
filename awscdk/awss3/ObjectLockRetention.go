package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The default retention settings for an S3 Object Lock configuration.
//
// Example:
//   // Configure for governance mode with a duration of 7 years
//   // Configure for governance mode with a duration of 7 years
//   s3.NewBucket(this, jsii.String("Bucket1"), &BucketProps{
//   	ObjectLockDefaultRetention: s3.ObjectLockRetention_Governance(awscdk.Duration_Days(jsii.Number(7 * 365))),
//   })
//
//   // Configure for compliance mode with a duration of 1 year
//   // Configure for compliance mode with a duration of 1 year
//   s3.NewBucket(this, jsii.String("Bucket2"), &BucketProps{
//   	ObjectLockDefaultRetention: s3.ObjectLockRetention_Compliance(awscdk.Duration_*Days(jsii.Number(365))),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html
//
type ObjectLockRetention interface {
	// The default period for which objects should be retained.
	Duration() awscdk.Duration
	// The retention mode to use for the object lock configuration.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-retention-modes
	//
	Mode() ObjectLockMode
}

// The jsii proxy struct for ObjectLockRetention
type jsiiProxy_ObjectLockRetention struct {
	_ byte // padding
}

func (j *jsiiProxy_ObjectLockRetention) Duration() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectLockRetention) Mode() ObjectLockMode {
	var returns ObjectLockMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Configure for Compliance retention for a specified duration.
//
// When an object is locked in compliance mode, its retention mode can't be changed, and
// its retention period can't be shortened. Compliance mode helps ensure that an object
// version can't be overwritten or deleted for the duration of the retention period.
//
// Returns: the ObjectLockRetention configuration.
func ObjectLockRetention_Compliance(duration awscdk.Duration) ObjectLockRetention {
	_init_.Initialize()

	if err := validateObjectLockRetention_ComplianceParameters(duration); err != nil {
		panic(err)
	}
	var returns ObjectLockRetention

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ObjectLockRetention",
		"compliance",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Configure for Governance retention for a specified duration.
//
// With governance mode, you protect objects against being deleted by most users, but you can
// still grant some users permission to alter the retention settings or delete the object if
// necessary. You can also use governance mode to test retention-period settings before
// creating a compliance-mode retention period.
//
// Returns: the ObjectLockRetention configuration.
func ObjectLockRetention_Governance(duration awscdk.Duration) ObjectLockRetention {
	_init_.Initialize()

	if err := validateObjectLockRetention_GovernanceParameters(duration); err != nil {
		panic(err)
	}
	var returns ObjectLockRetention

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ObjectLockRetention",
		"governance",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

