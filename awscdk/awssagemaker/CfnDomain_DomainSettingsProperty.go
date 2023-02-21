package awssagemaker


// A collection of settings that apply to the `SageMaker Domain` .
//
// These settings are specified through the `CreateDomain` API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSettingsProperty := &DomainSettingsProperty{
//   	RStudioServerProDomainSettings: &RStudioServerProDomainSettingsProperty{
//   		DomainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   		// the properties below are optional
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		RStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   		RStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type CfnDomain_DomainSettingsProperty struct {
	// A collection of settings that configure the `RStudioServerPro` Domain-level app.
	RStudioServerProDomainSettings interface{} `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the `Domain` uses for communication between Domain-level apps and user apps.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

