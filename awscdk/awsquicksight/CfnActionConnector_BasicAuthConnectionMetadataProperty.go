package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthConnectionMetadataProperty := &BasicAuthConnectionMetadataProperty{
//   	BaseEndpoint: jsii.String("baseEndpoint"),
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html
//
type CfnActionConnector_BasicAuthConnectionMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-baseendpoint
	//
	BaseEndpoint *string `field:"required" json:"baseEndpoint" yaml:"baseEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.html#cfn-quicksight-actionconnector-basicauthconnectionmetadata-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
}

