package awss3


// The configuration options of the S3 Access Grants location.
//
// It contains the `S3SubPrefix` field. The grant scope, the data to which you are granting access, is the result of appending the `Subprefix` field to the scope of the registered location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantsLocationConfigurationProperty := &AccessGrantsLocationConfigurationProperty{
//   	S3SubPrefix: jsii.String("s3SubPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-accessgrantslocationconfiguration.html
//
type CfnAccessGrant_AccessGrantsLocationConfigurationProperty struct {
	// The `S3SubPrefix` is appended to the location scope creating the grant scope.
	//
	// Use this field to narrow the scope of the grant to a subset of the location scope. This field is required if the location scope is the default location `s3://` because you cannot create a grant for all of your S3 data in the Region and must narrow the scope. For example, if the location scope is the default location `s3://` , the `S3SubPrefx` can be a `<bucket-name>/*` , so the full grant scope path would be `s3://<bucket-name>/*` . Or the `S3SubPrefx` can be `<bucket-name>/<prefix-name>*` , so the full grant scope path would be `s3://<bucket-name>/<prefix-name>*` .
	//
	// If the `S3SubPrefix` includes a prefix, append the wildcard character `*` after the prefix to indicate that you want to include all object key names in the bucket that start with that prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-accessgrantslocationconfiguration.html#cfn-s3-accessgrant-accessgrantslocationconfiguration-s3subprefix
	//
	S3SubPrefix *string `field:"required" json:"s3SubPrefix" yaml:"s3SubPrefix"`
}

