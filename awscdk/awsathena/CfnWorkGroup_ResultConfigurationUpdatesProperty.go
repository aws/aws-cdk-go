package awsathena


// The result configuration information about the queries in this workgroup that will be updated.
//
// Includes the updated results location and an updated option for encrypting query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resultConfigurationUpdatesProperty := &ResultConfigurationUpdatesProperty{
//   	AclConfiguration: &AclConfigurationProperty{
//   		S3AclOption: jsii.String("s3AclOption"),
//   	},
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionOption: jsii.String("encryptionOption"),
//
//   		// the properties below are optional
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	OutputLocation: jsii.String("outputLocation"),
//   	RemoveAclConfiguration: jsii.Boolean(false),
//   	RemoveEncryptionConfiguration: jsii.Boolean(false),
//   	RemoveExpectedBucketOwner: jsii.Boolean(false),
//   	RemoveOutputLocation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html
//
type CfnWorkGroup_ResultConfigurationUpdatesProperty struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-aclconfiguration
	//
	AclConfiguration interface{} `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The AWS account ID of the owner of S3 bucket where query results are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/.
	//
	// To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-outputlocation
	//
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeaclconfiguration
	//
	RemoveAclConfiguration interface{} `field:"optional" json:"removeAclConfiguration" yaml:"removeAclConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeencryptionconfiguration
	//
	RemoveEncryptionConfiguration interface{} `field:"optional" json:"removeEncryptionConfiguration" yaml:"removeEncryptionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeexpectedbucketowner
	//
	RemoveExpectedBucketOwner interface{} `field:"optional" json:"removeExpectedBucketOwner" yaml:"removeExpectedBucketOwner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-resultconfigurationupdates.html#cfn-athena-workgroup-resultconfigurationupdates-removeoutputlocation
	//
	RemoveOutputLocation interface{} `field:"optional" json:"removeOutputLocation" yaml:"removeOutputLocation"`
}

