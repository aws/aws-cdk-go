package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   basicAuthConnectionMetadataProperty := &BasicAuthConnectionMetadataProperty{
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html
//
type CfnActionConnectorPropsMixin_BasicAuthConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"optional" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

