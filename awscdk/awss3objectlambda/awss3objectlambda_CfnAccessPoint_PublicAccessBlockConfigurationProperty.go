package awss3objectlambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &publicAccessBlockConfigurationProperty{
//   	blockPublicAcls: jsii.Boolean(false),
//   	blockPublicPolicy: jsii.Boolean(false),
//   	ignorePublicAcls: jsii.Boolean(false),
//   	restrictPublicBuckets: jsii.Boolean(false),
//   }
//
type CfnAccessPoint_PublicAccessBlockConfigurationProperty struct {
	// `CfnAccessPoint.PublicAccessBlockConfigurationProperty.BlockPublicAcls`.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// `CfnAccessPoint.PublicAccessBlockConfigurationProperty.BlockPublicPolicy`.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// `CfnAccessPoint.PublicAccessBlockConfigurationProperty.IgnorePublicAcls`.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// `CfnAccessPoint.PublicAccessBlockConfigurationProperty.RestrictPublicBuckets`.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

