package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Collection of grant methods for a Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucketRef IBucketRef
//
//   bucketGrants := awscdk.Aws_s3.BucketGrants_FromBucket(bucketRef)
//
type BucketGrants interface {
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	Delete(grantee awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//     const grant = bucket.grantPublicAccess();
	//     grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	//
	// Note that if this `IBucket` refers to an existing bucket, possibly not
	// managed by CloudFormation, this method will have no effect, since it's
	// impossible to modify the policy of an existing bucket.
	PublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	Put(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grants s3:PutObjectAcl and s3:PutObjectVersionAcl permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	PutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	Read(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant read and write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	ReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant replication permission to a principal. This method allows the principal to perform replication operations on this bucket.
	//
	// Note that when calling this function for source or destination buckets that support KMS encryption,
	// you need to specify the KMS key for encryption and the KMS key for decryption, respectively.
	ReplicationPermission(identity awsiam.IGrantable, props *GrantReplicationPermissionProps) awsiam.Grant
	// Grant write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	Write(identity awsiam.IGrantable, objectsKeyPattern interface{}, allowedActionPatterns *[]*string) awsiam.Grant
}

// The jsii proxy struct for BucketGrants
type jsiiProxy_BucketGrants struct {
	_ byte // padding
}

// Creates grants for an IBucketRef.
func BucketGrants_FromBucket(bucket interfacesawss3.IBucketRef) BucketGrants {
	_init_.Initialize()

	if err := validateBucketGrants_FromBucketParameters(bucket); err != nil {
		panic(err)
	}
	var returns BucketGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketGrants",
		"fromBucket",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) Delete(grantee awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateDeleteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"delete",
		[]interface{}{grantee, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) PublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant {
	args := []interface{}{keyPrefix}
	for _, a := range allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"publicAccess",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) Put(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validatePutParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"put",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) PutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant {
	if err := b.validatePutAclParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"putAcl",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) Read(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateReadParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"read",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) ReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateReadWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"readWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) ReplicationPermission(identity awsiam.IGrantable, props *GrantReplicationPermissionProps) awsiam.Grant {
	if err := b.validateReplicationPermissionParameters(identity, props); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"replicationPermission",
		[]interface{}{identity, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketGrants) Write(identity awsiam.IGrantable, objectsKeyPattern interface{}, allowedActionPatterns *[]*string) awsiam.Grant {
	if err := b.validateWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"write",
		[]interface{}{identity, objectsKeyPattern, allowedActionPatterns},
		&returns,
	)

	return returns
}

