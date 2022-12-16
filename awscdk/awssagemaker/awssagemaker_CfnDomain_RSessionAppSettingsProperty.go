package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rSessionAppSettingsProperty := &rSessionAppSettingsProperty{
//   	customImages: []interface{}{
//   		&customImageProperty{
//   			appImageConfigName: jsii.String("appImageConfigName"),
//   			imageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			imageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnDomain_RSessionAppSettingsProperty struct {
	// `CfnDomain.RSessionAppSettingsProperty.CustomImages`.
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// `CfnDomain.RSessionAppSettingsProperty.DefaultResourceSpec`.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

