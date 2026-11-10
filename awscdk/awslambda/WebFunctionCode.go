package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Code for a `WebFunction`.
//
// WebFunction always reads its code from S3 at runtime, on every sandbox start, using the
// Lambda service principal. The bucket holding the code must therefore have S3 versioning
// enabled — enforced at synth time for buckets defined in the same app — and grant
// `lambda.amazonaws.com` read access to the code object. The construct adds that grant
// automatically on owned buckets (best-effort — imported buckets emit a warning instead).
//
// Bucket encryption compatibility:
//
// - `BucketEncryption.S3_MANAGED` (SSE-S3, AES256) — works, no key policy needed
// - `BucketEncryption.KMS` or `BucketEncryption.DSSE` with a customer managed key —
//   works; `kms:Decrypt` is granted to the Lambda service principal automatically
//   (best-effort for imported keys)
// - `BucketEncryption.KMS_MANAGED` (SSE-KMS with `aws/s3`) and
//   `BucketEncryption.DSSE_MANAGED` — do not work: an AWS-managed key policy cannot
//   be modified to authorize the Lambda service principal. For a bucket defined in
//   this app this fails at synthesis; an imported bucket exposes no encryption
//   settings, so there it surfaces as a read failure at runtime
//
// Use one of the static factory methods to create an instance.
//
// Unlike `lambda.Code`, there is deliberately no plain `fromAsset()`: each asset factory
// names the bucket the code is staged to, so the choice is explicit at the call site.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("PerRegionApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_PerRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		}),
//   	},
//   })
//
type WebFunctionCode interface {
	// Binds the code to a construct and returns the S3 location.
	//
	// Returns: The S3 location configuration for the code.
	Bind(scope constructs.Construct, options *WebFunctionCodeBindOptions) *WebFunctionCodeConfig
	// Called after the CloudFormation resource has been created, to let the code object bind to it.
	//
	// A no-op for every current code source.
	//
	// `lambda.Code` has the same hook, and its `AssetCode` uses it to attach asset metadata
	// (`aws:asset:path` etc.) for tooling. That is deliberately not done here: of the consumers
	// of that metadata, SAM CLI local invoke does not support WebFunction, and
	// `cdk deploy --hotswap` would be semantically wrong — it performs an in-place code
	// update, whereas WebFunction's model is "create a new revision, then shift traffic".
	// An override can be added once a consumer is verified to handle WebFunction correctly.
	BindToResource(_resource awscdk.CfnResource)
}

// The jsii proxy struct for WebFunctionCode
type jsiiProxy_WebFunctionCode struct {
	_ byte // padding
}

func NewWebFunctionCode_Override(w WebFunctionCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.WebFunctionCode",
		nil, // no parameters
		w,
	)
}

// Create WebFunction code from a local file or directory, staged into a bucket that CDK creates and configures for you.
//
// One bucket is created per stack and shared by every WebFunction in it that uses this
// factory. It is configured with everything WebFunction requires: versioning, AES256
// (SSE-S3) encryption, and a bucket policy granting `lambda.amazonaws.com` read access.
//
// The bucket is removed and emptied along with the stack that owns it. If a revision whose
// code it holds is retained, the bucket is retained as well, with a warning.
//
// Not supported in an environment bootstrapped with a customer managed KMS key; see the README.
func WebFunctionCode_FromAssetToManagedBucket(path *string, options *WebFunctionManagedBucketAssetOptions) WebFunctionCode {
	_init_.Initialize()

	if err := validateWebFunctionCode_FromAssetToManagedBucketParameters(path, options); err != nil {
		panic(err)
	}
	var returns WebFunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionCode",
		"fromAssetToManagedBucket",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create WebFunction code from a local file or directory, staged into a bucket you own.
//
// You own the bucket, so CDK only makes the narrowest change it needs: it adds a resource
// policy statement granting `lambda.amazonaws.com` read access to this function's code, and
// `kms:Decrypt` on the bucket's key if it is encrypted with a customer managed key. Both
// are best-effort — on an imported bucket or key they are no-ops and a warning is emitted
// instead, since the policy cannot be modified from this app.
//
// You are responsible for the bucket settings CDK cannot change:
//
// - S3 versioning must be enabled. This is validated at synth time for a bucket defined in
//   the same app; an imported bucket is only caught at deployment time, where WebFunction
//   rejects a bucket that reports no object version.
// - Encryption must not be `BucketEncryption.KMS_MANAGED` (SSE-KMS with `aws/s3`) or
//   `BucketEncryption.DSSE_MANAGED`, whose key policy cannot be modified to authorize the
//   Lambda service principal. SSE-S3 and a customer managed key both work.
//
// Not supported in an environment bootstrapped with a customer managed KMS key; see the README.
func WebFunctionCode_FromAssetToProvidedBucket(bucket awss3.IBucket, path *string, options *WebFunctionProvidedBucketAssetOptions) WebFunctionCode {
	_init_.Initialize()

	if err := validateWebFunctionCode_FromAssetToProvidedBucketParameters(bucket, path, options); err != nil {
		panic(err)
	}
	var returns WebFunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionCode",
		"fromAssetToProvidedBucket",
		[]interface{}{bucket, path, options},
		&returns,
	)

	return returns
}

// Create WebFunction code from an S3 bucket.
//
// This method configures the bucket on a best-effort basis so it satisfies
// WebFunction's runtime requirements:
//
// - It adds a resource policy statement granting `lambda.amazonaws.com`
//   `s3:GetObject` and `s3:GetObjectVersion`, scoped to the given object key,
//   to the deploying account and to the revisions of the bound WebFunction.
//   For buckets imported by attributes (e.g.
//   `Bucket.fromBucketName`) the policy cannot be modified from this app;
//   ensure the bucket already carries an equivalent statement.
// - This method does not change the bucket's versioning setting, so it is
//   validated instead: a bucket defined in this app without
//   `versioned: true` fails at synthesis. Imported buckets cannot be
//   inspected, so for those the requirement is only enforced at deployment
//   time, where WebFunction rejects a bucket that reports no object version.
// - Encryption with an AWS-managed KMS key is asserted against for the same
//   reason: a bucket defined in this app using `BucketEncryption.KMS_MANAGED`
//   or `BucketEncryption.DSSE_MANAGED` fails at synthesis, since that key's
//   policy cannot be modified to authorize the service principal. Imported
//   buckets cannot be inspected and are assumed to be compatible.
func WebFunctionCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) WebFunctionCode {
	_init_.Initialize()

	if err := validateWebFunctionCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns WebFunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebFunctionCode) Bind(scope constructs.Construct, options *WebFunctionCodeBindOptions) *WebFunctionCodeConfig {
	if err := w.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *WebFunctionCodeConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebFunctionCode) BindToResource(_resource awscdk.CfnResource) {
	if err := w.validateBindToResourceParameters(_resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"bindToResource",
		[]interface{}{_resource},
	)
}

