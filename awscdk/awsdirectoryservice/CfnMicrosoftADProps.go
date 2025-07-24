package awsdirectoryservice


// Properties for defining a `CfnMicrosoftAD`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMicrosoftADProps := &CfnMicrosoftADProps{
//   	Name: jsii.String("name"),
//   	Password: jsii.String("password"),
//   	VpcSettings: &VpcSettingsProperty{
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//
//   	// the properties below are optional
//   	CreateAlias: jsii.Boolean(false),
//   	Edition: jsii.String("edition"),
//   	EnableSso: jsii.Boolean(false),
//   	ShortName: jsii.String("shortName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
//
type CfnMicrosoftADProps struct {
	// The fully qualified domain name for the AWS Managed Microsoft AD directory, such as `corp.example.com` . This name will resolve inside your VPC only. It does not need to be publicly resolvable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The password for the default administrative user named `Admin` .
	//
	// If you need to change the password for the administrator account, see the [ResetUserPassword](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_ResetUserPassword.html) API call in the *AWS Directory Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the VPC settings of the Microsoft AD directory server in AWS .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-vpcsettings
	//
	VpcSettings interface{} `field:"required" json:"vpcSettings" yaml:"vpcSettings"`
	// Specifies an alias for a directory and assigns the alias to the directory.
	//
	// The alias is used to construct the access URL for the directory, such as `http://<alias>.awsapps.com` . By default, AWS CloudFormation does not create an alias.
	//
	// > After an alias has been created, it cannot be deleted or reused, so this operation should only be used when absolutely necessary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-createalias
	//
	CreateAlias interface{} `field:"optional" json:"createAlias" yaml:"createAlias"`
	// AWS Managed Microsoft AD is available in two editions: `Standard` and `Enterprise` .
	//
	// `Enterprise` is the default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-edition
	//
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Whether to enable single sign-on for a Microsoft Active Directory in AWS .
	//
	// Single sign-on allows users in your directory to access certain AWS services from a computer joined to the directory without having to enter their credentials separately. If you don't specify a value, AWS CloudFormation disables single sign-on by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-enablesso
	//
	EnableSso interface{} `field:"optional" json:"enableSso" yaml:"enableSso"`
	// The NetBIOS name for your domain, such as `CORP` .
	//
	// If you don't specify a NetBIOS name, it will default to the first part of your directory DNS. For example, `CORP` for the directory DNS `corp.example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-shortname
	//
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
}

