package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   synchronizationConfigurationProperty := &SynchronizationConfigurationProperty{
//   	ExpirationDataRules: []interface{}{
//   		&ExpirationDataRuleProperty{
//   			DaysAfterLastAccess: jsii.Number(123),
//   		},
//   	},
//   	ImportDataRules: []interface{}{
//   		&ImportDataRuleProperty{
//   			Prefix: jsii.String("prefix"),
//   			SizeLessThan: jsii.Number(123),
//   			Trigger: jsii.String("trigger"),
//   		},
//   	},
//   	LatestVersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-synchronizationconfiguration.html
//
type CfnFileSystemPropsMixin_SynchronizationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-synchronizationconfiguration.html#cfn-s3files-filesystem-synchronizationconfiguration-expirationdatarules
	//
	ExpirationDataRules interface{} `field:"optional" json:"expirationDataRules" yaml:"expirationDataRules"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-synchronizationconfiguration.html#cfn-s3files-filesystem-synchronizationconfiguration-importdatarules
	//
	ImportDataRules interface{} `field:"optional" json:"importDataRules" yaml:"importDataRules"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-synchronizationconfiguration.html#cfn-s3files-filesystem-synchronizationconfiguration-latestversionnumber
	//
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
}

