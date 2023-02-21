package awssagemaker


// A collection of settings that apply to users of Amazon SageMaker Studio.
//
// These settings are specified when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called, and as `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called.
//
// `SecurityGroups` is aggregated when specified in both calls. For all other settings in `UserSettings` , the values specified in `CreateUserProfile` take precedence over those specified in `CreateDomain` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingsProperty := &userSettingsProperty{
//   	executionRole: jsii.String("executionRole"),
//   	jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   		customImages: []interface{}{
//   			&customImageProperty{
//   				appImageConfigName: jsii.String("appImageConfigName"),
//   				imageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				imageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   		accessStatus: jsii.String("accessStatus"),
//   		userGroup: jsii.String("userGroup"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	sharingSettings: &sharingSettingsProperty{
//   		notebookOutputOption: jsii.String("notebookOutputOption"),
//   		s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   		s3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   }
//
type CfnUserProfile_UserSettingsProperty struct {
	// The execution role for the user.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The Jupyter server's app settings.
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A collection of settings that configure user interaction with the `RStudioServerPro` app.
	RStudioServerProAppSettings interface{} `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	//
	// Optional when the `CreateDomain.AppNetworkAccessType` parameter is set to `PublicInternetOnly` .
	//
	// Required when the `CreateDomain.AppNetworkAccessType` parameter is set to `VpcOnly` .
	//
	// Amazon SageMaker adds a security group to allow NFS traffic from SageMaker Studio. Therefore, the number of security groups that you can specify is one less than the maximum number shown.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies options for sharing SageMaker Studio notebooks.
	SharingSettings interface{} `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
}

