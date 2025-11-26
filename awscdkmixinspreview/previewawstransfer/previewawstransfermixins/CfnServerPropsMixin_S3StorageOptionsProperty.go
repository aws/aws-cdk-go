package previewawstransfermixins


// The Amazon S3 storage options that are configured for your server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3StorageOptionsProperty := &S3StorageOptionsProperty{
//   	DirectoryListingOptimization: jsii.String("directoryListingOptimization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-s3storageoptions.html
//
type CfnServerPropsMixin_S3StorageOptionsProperty struct {
	// Specifies whether or not performance for your Amazon S3 directories is optimized.
	//
	// - If using the console, this is enabled by default.
	// - If using the API or CLI, this is disabled by default.
	//
	// By default, home directory mappings have a `TYPE` of `DIRECTORY` . If you enable this option, you would then need to explicitly set the `HomeDirectoryMapEntry` `Type` to `FILE` if you want a mapping to have a file target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-s3storageoptions.html#cfn-transfer-server-s3storageoptions-directorylistingoptimization
	//
	DirectoryListingOptimization *string `field:"optional" json:"directoryListingOptimization" yaml:"directoryListingOptimization"`
}

