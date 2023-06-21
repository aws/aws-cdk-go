package awscleanrooms


// The settings for client-side encryption for cryptographic computing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataEncryptionMetadataProperty := &DataEncryptionMetadataProperty{
//   	AllowCleartext: jsii.Boolean(false),
//   	AllowDuplicates: jsii.Boolean(false),
//   	AllowJoinsOnColumnsWithDifferentNames: jsii.Boolean(false),
//   	PreserveNulls: jsii.Boolean(false),
//   }
//
type CfnCollaboration_DataEncryptionMetadataProperty struct {
	// Indicates whether encrypted tables can contain cleartext data (true) or are to cryptographically process every column (false).
	AllowCleartext interface{} `field:"required" json:"allowCleartext" yaml:"allowCleartext"`
	// Indicates whether Fingerprint columns can contain duplicate entries (true) or are to contain only non-repeated values (false).
	AllowDuplicates interface{} `field:"required" json:"allowDuplicates" yaml:"allowDuplicates"`
	// Indicates whether Fingerprint columns can be joined on any other Fingerprint column with a different name (true) or can only be joined on Fingerprint columns of the same name (false).
	AllowJoinsOnColumnsWithDifferentNames interface{} `field:"required" json:"allowJoinsOnColumnsWithDifferentNames" yaml:"allowJoinsOnColumnsWithDifferentNames"`
	// Indicates whether NULL values are to be copied as NULL to encrypted tables (true) or cryptographically processed (false).
	PreserveNulls interface{} `field:"required" json:"preserveNulls" yaml:"preserveNulls"`
}

