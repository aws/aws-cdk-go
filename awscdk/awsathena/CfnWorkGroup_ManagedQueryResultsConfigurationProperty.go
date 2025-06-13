package awsathena


// The configuration for the managed query results and encryption option.
//
// ResultConfiguration and ManagedQueryResultsConfiguration cannot be set at the same time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedQueryResultsConfigurationProperty := &ManagedQueryResultsConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedqueryresultsconfiguration.html
//
type CfnWorkGroup_ManagedQueryResultsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedqueryresultsconfiguration.html#cfn-athena-workgroup-managedqueryresultsconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the encryption configuration for Athena Managed Storage.
	//
	// If not setting this field, Managed Storage will encrypt the query results with Athena's encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedqueryresultsconfiguration.html#cfn-athena-workgroup-managedqueryresultsconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

