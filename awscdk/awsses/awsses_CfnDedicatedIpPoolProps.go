package awsses


// Properties for defining a `CfnDedicatedIpPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPoolProps := &cfnDedicatedIpPoolProps{
//   	poolName: jsii.String("poolName"),
//   	scalingMode: jsii.String("scalingMode"),
//   }
//
type CfnDedicatedIpPoolProps struct {
	// `AWS::SES::DedicatedIpPool.PoolName`.
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// `AWS::SES::DedicatedIpPool.ScalingMode`.
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
}

