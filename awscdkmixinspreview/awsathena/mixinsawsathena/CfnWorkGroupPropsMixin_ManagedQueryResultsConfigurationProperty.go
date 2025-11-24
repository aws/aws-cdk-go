package mixinsawsathena


// The configuration for storing results in Athena owned storage, which includes whether this feature is enabled;
//
// whether encryption configuration, if any, is used for encrypting query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnWorkGroupPropsMixin_ManagedQueryResultsConfigurationProperty struct {
	// If set to true, allows you to store query results in Athena owned storage.
	//
	// If set to false, workgroup member stores query results in location specified under `ResultConfiguration$OutputLocation` . The default is false. A workgroup cannot have the `ResultConfiguration$OutputLocation` parameter when you set this field to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedqueryresultsconfiguration.html#cfn-athena-workgroup-managedqueryresultsconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If you encrypt query and calculation results in Athena owned storage, this field indicates the encryption option (for example, SSE_KMS or CSE_KMS) and key information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedqueryresultsconfiguration.html#cfn-athena-workgroup-managedqueryresultsconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

