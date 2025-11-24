package mixinsawspcaconnectorad


// Properties for CfnServicePrincipalNamePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServicePrincipalNameMixinProps := &CfnServicePrincipalNameMixinProps{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	DirectoryRegistrationArn: jsii.String("directoryRegistrationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-serviceprincipalname.html
//
type CfnServicePrincipalNameMixinProps struct {
	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector.html](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-serviceprincipalname.html#cfn-pcaconnectorad-serviceprincipalname-connectorarn
	//
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
	// The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-serviceprincipalname.html#cfn-pcaconnectorad-serviceprincipalname-directoryregistrationarn
	//
	DirectoryRegistrationArn *string `field:"optional" json:"directoryRegistrationArn" yaml:"directoryRegistrationArn"`
}

