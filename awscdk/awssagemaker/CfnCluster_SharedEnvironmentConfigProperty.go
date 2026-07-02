package awssagemaker


// The shared environment configuration for restricted instance groups that use cluster-level shared FSx Lustre storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharedEnvironmentConfigProperty := &SharedEnvironmentConfigProperty{
//   	FSxLustreDeletionPolicy: jsii.String("fSxLustreDeletionPolicy"),
//
//   	// the properties below are optional
//   	FSxLustreConfig: &FSxLustreConfigProperty{
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		SizeInGiB: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-sharedenvironmentconfig.html
//
type CfnCluster_SharedEnvironmentConfigProperty struct {
	// The deletion policy for the shared FSx Lustre file system.
	//
	// Keep retains the FSx when RIGs are deleted. DeleteIfNotUsed deletes the FSx when no RIGs reference it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-sharedenvironmentconfig.html#cfn-sagemaker-cluster-sharedenvironmentconfig-fsxlustredeletionpolicy
	//
	FSxLustreDeletionPolicy *string `field:"required" json:"fSxLustreDeletionPolicy" yaml:"fSxLustreDeletionPolicy"`
	// Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-sharedenvironmentconfig.html#cfn-sagemaker-cluster-sharedenvironmentconfig-fsxlustreconfig
	//
	FSxLustreConfig interface{} `field:"optional" json:"fSxLustreConfig" yaml:"fSxLustreConfig"`
}

