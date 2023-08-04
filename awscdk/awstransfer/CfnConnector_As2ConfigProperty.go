package awstransfer


// A structure that contains the parameters for an AS2 connector object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   as2ConfigProperty := &As2ConfigProperty{
//   	Compression: jsii.String("compression"),
//   	EncryptionAlgorithm: jsii.String("encryptionAlgorithm"),
//   	LocalProfileId: jsii.String("localProfileId"),
//   	MdnResponse: jsii.String("mdnResponse"),
//   	MdnSigningAlgorithm: jsii.String("mdnSigningAlgorithm"),
//   	MessageSubject: jsii.String("messageSubject"),
//   	PartnerProfileId: jsii.String("partnerProfileId"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html
//
type CfnConnector_As2ConfigProperty struct {
	// Specifies whether the AS2 file is compressed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-compression
	//
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The algorithm that is used to encrypt the file.
	//
	// > You can only specify `NONE` if the URL for your connector uses HTTPS. This ensures that no traffic is sent in clear text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-encryptionalgorithm
	//
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// A unique identifier for the AS2 local profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-localprofileid
	//
	LocalProfileId *string `field:"optional" json:"localProfileId" yaml:"localProfileId"`
	// Used for outbound requests (from an AWS Transfer Family server to a partner AS2 server) to determine whether the partner response for transfers is synchronous or asynchronous.
	//
	// Specify either of the following values:
	//
	// - `SYNC` : The system expects a synchronous MDN response, confirming that the file was transferred successfully (or not).
	// - `NONE` : Specifies that no MDN response is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-mdnresponse
	//
	MdnResponse *string `field:"optional" json:"mdnResponse" yaml:"mdnResponse"`
	// The signing algorithm for the MDN response.
	//
	// > If set to DEFAULT (or not set at all), the value for `SigningAlgorithm` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-mdnsigningalgorithm
	//
	MdnSigningAlgorithm *string `field:"optional" json:"mdnSigningAlgorithm" yaml:"mdnSigningAlgorithm"`
	// Used as the `Subject` HTTP header attribute in AS2 messages that are being sent with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-messagesubject
	//
	MessageSubject *string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
	// A unique identifier for the partner profile for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-partnerprofileid
	//
	PartnerProfileId *string `field:"optional" json:"partnerProfileId" yaml:"partnerProfileId"`
	// The algorithm that is used to sign the AS2 messages sent with the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-signingalgorithm
	//
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

