package awssagemaker


// A collection of settings that configure the `RStudioServerPro` Domain-level app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProDomainSettingsProperty := &RStudioServerProDomainSettingsProperty{
//   	DomainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   	// the properties below are optional
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	RStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   	RStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   }
//
type CfnDomain_RStudioServerProDomainSettingsProperty struct {
	// The ARN of the execution role for the `RStudioServerPro` Domain-level app.
	DomainExecutionRoleArn *string `field:"required" json:"domainExecutionRoleArn" yaml:"domainExecutionRoleArn"`
	// A collection that defines the default `InstanceType` , `SageMakerImageArn` , and `SageMakerImageVersionArn` for the Domain.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A URL pointing to an RStudio Connect server.
	RStudioConnectUrl *string `field:"optional" json:"rStudioConnectUrl" yaml:"rStudioConnectUrl"`
	// A URL pointing to an RStudio Package Manager server.
	RStudioPackageManagerUrl *string `field:"optional" json:"rStudioPackageManagerUrl" yaml:"rStudioPackageManagerUrl"`
}

