package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &cfnDomainProps{
//   	authMode: jsii.String("authMode"),
//   	defaultUserSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rSessionAppSettings: &rSessionAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   	domainName: jsii.String("domainName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	appNetworkAccessType: jsii.String("appNetworkAccessType"),
//   	appSecurityGroupManagement: jsii.String("appSecurityGroupManagement"),
//   	domainSettings: &domainSettingsProperty{
//   		rStudioServerProDomainSettings: &rStudioServerProDomainSettingsProperty{
//   			domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   			// the properties below are optional
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   			rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainProps struct {
	// The mode of authentication that members use to access the Domain.
	//
	// *Valid Values* : `SSO | IAM`.
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The default user settings.
	DefaultUserSettings interface{} `field:"required" json:"defaultUserSettings" yaml:"defaultUserSettings"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The VPC subnets that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Array members* : Minimum number of 1 item. Maximum number of 16 items.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly` .
	//
	// - `PublicInternetOnly` - Non-EFS traffic is through a VPC managed by Amazon SageMaker , which allows direct internet access
	// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets
	//
	// *Valid Values* : `PublicInternetOnly | VpcOnly`.
	AppNetworkAccessType *string `field:"optional" json:"appNetworkAccessType" yaml:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in `VpcOnly` mode.
	//
	// Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and `DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided.
	AppSecurityGroupManagement *string `field:"optional" json:"appSecurityGroupManagement" yaml:"appSecurityGroupManagement"`
	// A collection of settings that apply to the `SageMaker Domain` .
	//
	// These settings are specified through the `CreateDomain` API call.
	DomainSettings interface{} `field:"optional" json:"domainSettings" yaml:"domainSettings"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default.
	//
	// For more control, specify a customer managed CMK.
	//
	// *Length Constraints* : Maximum length of 2048.
	//
	// *Pattern* : `.*`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Tags to associated with the Domain.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Tags that you specify for the Domain are also added to all apps that are launched in the Domain.
	//
	// *Array members* : Minimum number of 0 items. Maximum number of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

