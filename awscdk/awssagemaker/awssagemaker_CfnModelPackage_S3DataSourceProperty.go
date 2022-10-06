package awssagemaker


// Describes the S3 data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceProperty := &s3DataSourceProperty{
//   	s3DataType: jsii.String("s3DataType"),
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelPackage_S3DataSourceProperty struct {
	// If you choose `S3Prefix` , `S3Uri` identifies a key name prefix.
	//
	// SageMaker uses all objects that match the specified key name prefix for model training.
	//
	// If you choose `ManifestFile` , `S3Uri` identifies an object that is a manifest file containing a list of object keys that you want SageMaker to use for model training.
	//
	// If you choose `AugmentedManifestFile` , S3Uri identifies an object that is an augmented manifest file in JSON lines format. This file contains the data you want to use for model training. `AugmentedManifestFile` can only be used if the Channel's input mode is `Pipe` .
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Depending on the value specified for the `S3DataType` , identifies either a key name prefix or a manifest.
	//
	// For example:
	//
	// - A key name prefix might look like this: `s3://bucketname/exampleprefix`
	// - A manifest might look like this: `s3://bucketname/example.manifest`
	//
	// A manifest is an S3 object which is a JSON file consisting of an array of elements. The first element is a prefix which is followed by one or more suffixes. SageMaker appends the suffix elements to the prefix to get a full set of `S3Uri` . Note that the prefix must be a valid non-empty `S3Uri` that precludes users from specifying a manifest whose individual `S3Uri` is sourced from different S3 buckets.
	//
	// The following code example shows a valid manifest format:
	//
	// `[ {"prefix": "s3://customer_bucket/some/prefix/"},`
	//
	// `"relative/path/to/custdata-1",`
	//
	// `"relative/path/custdata-2",`
	//
	// `...`
	//
	// `"relative/path/custdata-N"`
	//
	// `]`
	//
	// This JSON is equivalent to the following `S3Uri` list:
	//
	// `s3://customer_bucket/some/prefix/relative/path/to/custdata-1`
	//
	// `s3://customer_bucket/some/prefix/relative/path/custdata-2`
	//
	// `...`
	//
	// `s3://customer_bucket/some/prefix/relative/path/custdata-N`
	//
	// The complete set of `S3Uri` in this manifest is the input data for the channel for this data source. The object that each `S3Uri` points to must be readable by the IAM role that SageMaker uses to perform tasks on your behalf.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

