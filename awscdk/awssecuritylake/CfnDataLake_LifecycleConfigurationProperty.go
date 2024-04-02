package awssecuritylake


// Provides lifecycle details of Amazon Security Lake object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	Expiration: &ExpirationProperty{
//   		Days: jsii.Number(123),
//   	},
//   	Transitions: []interface{}{
//   		&TransitionsProperty{
//   			Days: jsii.Number(123),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html
//
type CfnDataLake_LifecycleConfigurationProperty struct {
	// Provides data expiration details of Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-expiration
	//
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// Provides data storage transition details of Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-transitions
	//
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

