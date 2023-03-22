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
//   domainSettingsProperty := &domainSettingsProperty{
//   	rStudioServerProDomainSettings: &rStudioServerProDomainSettingsProperty{
//   		domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   		// the properties below are optional
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   		rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   	},
//   	securityGroupIds: []*string{
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

