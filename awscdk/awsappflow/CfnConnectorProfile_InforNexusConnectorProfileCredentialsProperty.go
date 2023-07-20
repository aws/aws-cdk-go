package awsappflow


// The connector-specific profile credentials required by Infor Nexus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inforNexusConnectorProfileCredentialsProperty := &InforNexusConnectorProfileCredentialsProperty{
//   	AccessKeyId: jsii.String("accessKeyId"),
//   	Datakey: jsii.String("datakey"),
//   	SecretAccessKey: jsii.String("secretAccessKey"),
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.html
//
type CfnConnectorProfile_InforNexusConnectorProfileCredentialsProperty struct {
	// The Access Key portion of the credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.html#cfn-appflow-connectorprofile-infornexusconnectorprofilecredentials-accesskeyid
	//
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The encryption keys used to encrypt data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.html#cfn-appflow-connectorprofile-infornexusconnectorprofilecredentials-datakey
	//
	Datakey *string `field:"required" json:"datakey" yaml:"datakey"`
	// The secret key used to sign requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.html#cfn-appflow-connectorprofile-infornexusconnectorprofilecredentials-secretaccesskey
	//
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The identifier for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.html#cfn-appflow-connectorprofile-infornexusconnectorprofilecredentials-userid
	//
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

