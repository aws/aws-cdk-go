package awss3vectors


// The metadata configuration for the vector index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	NonFilterableMetadataKeys: []*string{
//   		jsii.String("nonFilterableMetadataKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-index-metadataconfiguration.html
//
type CfnIndex_MetadataConfigurationProperty struct {
	// Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval.
	//
	// Unlike default metadata keys, these keys cannot be used as query filters. Non-filterable metadata keys can be retrieved but cannot be searched, queried, or filtered. You can access non-filterable metadata keys of your vectors after finding the vectors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-index-metadataconfiguration.html#cfn-s3vectors-index-metadataconfiguration-nonfilterablemetadatakeys
	//
	NonFilterableMetadataKeys *[]*string `field:"optional" json:"nonFilterableMetadataKeys" yaml:"nonFilterableMetadataKeys"`
}

