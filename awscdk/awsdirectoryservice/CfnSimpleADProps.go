package awsdirectoryservice


// Properties for defining a `CfnSimpleAD`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimpleADProps := &CfnSimpleADProps{
//   	Name: jsii.String("name"),
//   	Size: jsii.String("size"),
//   	VpcSettings: &VpcSettingsProperty{
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//
//   	// the properties below are optional
//   	CreateAlias: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnableSso: jsii.Boolean(false),
//   	Password: jsii.String("password"),
//   	ShortName: jsii.String("shortName"),
//   }
//
type CfnSimpleADProps struct {
	// The fully qualified name for the directory, such as `corp.example.com` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size of the directory.
	//
	// For valid values, see [CreateDirectory](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateDirectory.html) in the *AWS Directory Service API Reference* .
	Size *string `field:"required" json:"size" yaml:"size"`
	// A [DirectoryVpcSettings](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_DirectoryVpcSettings.html) object that contains additional information for the operation.
	VpcSettings interface{} `field:"required" json:"vpcSettings" yaml:"vpcSettings"`
	// If set to `true` , specifies an alias for a directory and assigns the alias to the directory.
	//
	// The alias is used to construct the access URL for the directory, such as `http://<alias>.awsapps.com` . By default, this property is set to `false` .
	//
	// > After an alias has been created, it cannot be deleted or reused, so this operation should only be used when absolutely necessary.
	CreateAlias interface{} `field:"optional" json:"createAlias" yaml:"createAlias"`
	// A description for the directory.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable single sign-on for a directory.
	//
	// If you don't specify a value, AWS CloudFormation disables single sign-on by default.
	EnableSso interface{} `field:"optional" json:"enableSso" yaml:"enableSso"`
	// The password for the directory administrator.
	//
	// The directory creation process creates a directory administrator account with the user name `Administrator` and this password.
	//
	// If you need to change the password for the administrator account, see the [ResetUserPassword](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_ResetUserPassword.html) API call in the *AWS Directory Service API Reference* .
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The NetBIOS name of the directory, such as `CORP` .
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
}

