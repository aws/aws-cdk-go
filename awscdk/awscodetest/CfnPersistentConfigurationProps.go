package awscodetest


// Properties for defining a `CfnPersistentConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPersistentConfigurationProps := &CfnPersistentConfigurationProps{
//   	ResultsRoleArn: jsii.String("resultsRoleArn"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-persistentconfiguration.html
//
type CfnPersistentConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-persistentconfiguration.html#cfn-codetest-persistentconfiguration-resultsrolearn
	//
	ResultsRoleArn *string `field:"required" json:"resultsRoleArn" yaml:"resultsRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-persistentconfiguration.html#cfn-codetest-persistentconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-persistentconfiguration.html#cfn-codetest-persistentconfiguration-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codetest-persistentconfiguration.html#cfn-codetest-persistentconfiguration-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

