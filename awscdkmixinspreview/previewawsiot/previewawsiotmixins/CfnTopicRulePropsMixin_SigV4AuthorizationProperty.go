package previewawsiotmixins


// For more information, see [Signature Version 4 signing process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sigV4AuthorizationProperty := &SigV4AuthorizationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	ServiceName: jsii.String("serviceName"),
//   	SigningRegion: jsii.String("signingRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sigv4authorization.html
//
type CfnTopicRulePropsMixin_SigV4AuthorizationProperty struct {
	// The ARN of the signing role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sigv4authorization.html#cfn-iot-topicrule-sigv4authorization-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The service name to use while signing with Sig V4.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sigv4authorization.html#cfn-iot-topicrule-sigv4authorization-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The signing region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sigv4authorization.html#cfn-iot-topicrule-sigv4authorization-signingregion
	//
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
}

