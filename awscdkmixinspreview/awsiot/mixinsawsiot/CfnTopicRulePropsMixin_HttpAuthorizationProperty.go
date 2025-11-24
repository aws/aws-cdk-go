package mixinsawsiot


// The authorization method used to send messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpAuthorizationProperty := &HttpAuthorizationProperty{
//   	Sigv4: &SigV4AuthorizationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		ServiceName: jsii.String("serviceName"),
//   		SigningRegion: jsii.String("signingRegion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpauthorization.html
//
type CfnTopicRulePropsMixin_HttpAuthorizationProperty struct {
	// Use Sig V4 authorization.
	//
	// For more information, see [Signature Version 4 Signing Process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-httpauthorization.html#cfn-iot-topicrule-httpauthorization-sigv4
	//
	Sigv4 interface{} `field:"optional" json:"sigv4" yaml:"sigv4"`
}

