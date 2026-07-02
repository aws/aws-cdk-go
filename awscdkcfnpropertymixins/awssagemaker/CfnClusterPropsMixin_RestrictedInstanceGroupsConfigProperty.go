package awssagemaker


// The cluster-level configuration for restricted instance groups, including shared environment settings for inter-RIG communication and FSx Lustre sharing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   restrictedInstanceGroupsConfigProperty := &RestrictedInstanceGroupsConfigProperty{
//   	SharedEnvironmentConfig: &SharedEnvironmentConfigProperty{
//   		FSxLustreConfig: &FSxLustreConfigProperty{
//   			PerUnitStorageThroughput: jsii.Number(123),
//   			SizeInGiB: jsii.Number(123),
//   		},
//   		FSxLustreDeletionPolicy: jsii.String("fSxLustreDeletionPolicy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-restrictedinstancegroupsconfig.html
//
type CfnClusterPropsMixin_RestrictedInstanceGroupsConfigProperty struct {
	// The shared environment configuration for restricted instance groups that use cluster-level shared FSx Lustre storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-restrictedinstancegroupsconfig.html#cfn-sagemaker-cluster-restrictedinstancegroupsconfig-sharedenvironmentconfig
	//
	SharedEnvironmentConfig interface{} `field:"optional" json:"sharedEnvironmentConfig" yaml:"sharedEnvironmentConfig"`
}

