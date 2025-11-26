package previewawscleanroomsmixins


// The settings for client-side encryption for cryptographic computing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataEncryptionMetadataProperty := &DataEncryptionMetadataProperty{
//   	AllowCleartext: jsii.Boolean(false),
//   	AllowDuplicates: jsii.Boolean(false),
//   	AllowJoinsOnColumnsWithDifferentNames: jsii.Boolean(false),
//   	PreserveNulls: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html
//
type CfnCollaborationPropsMixin_DataEncryptionMetadataProperty struct {
	// Indicates whether encrypted tables can contain cleartext data ( `TRUE` ) or are to cryptographically process every column ( `FALSE` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html#cfn-cleanrooms-collaboration-dataencryptionmetadata-allowcleartext
	//
	AllowCleartext interface{} `field:"optional" json:"allowCleartext" yaml:"allowCleartext"`
	// Indicates whether Fingerprint columns can contain duplicate entries ( `TRUE` ) or are to contain only non-repeated values ( `FALSE` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html#cfn-cleanrooms-collaboration-dataencryptionmetadata-allowduplicates
	//
	AllowDuplicates interface{} `field:"optional" json:"allowDuplicates" yaml:"allowDuplicates"`
	// Indicates whether Fingerprint columns can be joined on any other Fingerprint column with a different name ( `TRUE` ) or can only be joined on Fingerprint columns of the same name ( `FALSE` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html#cfn-cleanrooms-collaboration-dataencryptionmetadata-allowjoinsoncolumnswithdifferentnames
	//
	AllowJoinsOnColumnsWithDifferentNames interface{} `field:"optional" json:"allowJoinsOnColumnsWithDifferentNames" yaml:"allowJoinsOnColumnsWithDifferentNames"`
	// Indicates whether NULL values are to be copied as NULL to encrypted tables ( `TRUE` ) or cryptographically processed ( `FALSE` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html#cfn-cleanrooms-collaboration-dataencryptionmetadata-preservenulls
	//
	PreserveNulls interface{} `field:"optional" json:"preserveNulls" yaml:"preserveNulls"`
}

