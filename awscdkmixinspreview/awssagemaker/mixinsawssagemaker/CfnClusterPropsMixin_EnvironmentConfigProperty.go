package mixinsawssagemaker


// The configuration for the restricted instance groups (RIG) environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentConfigProperty := &EnvironmentConfigProperty{
//   	FSxLustreConfig: &FSxLustreConfigProperty{
//   		PerUnitStorageThroughput: jsii.Number(123),
//   		SizeInGiB: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-environmentconfig.html
//
type CfnClusterPropsMixin_EnvironmentConfigProperty struct {
	// Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-environmentconfig.html#cfn-sagemaker-cluster-environmentconfig-fsxlustreconfig
	//
	FSxLustreConfig interface{} `field:"optional" json:"fSxLustreConfig" yaml:"fSxLustreConfig"`
}

