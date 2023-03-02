package awsiot


// The authorization method used to send messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpAuthorizationProperty := &httpAuthorizationProperty{
//   	sigv4: &sigV4AuthorizationProperty{
//   		roleArn: jsii.String("roleArn"),
//   		serviceName: jsii.String("serviceName"),
//   		signingRegion: jsii.String("signingRegion"),
//   	},
//   }
//
type CfnTopicRule_HttpAuthorizationProperty struct {
	// Use Sig V4 authorization.
	//
	// For more information, see [Signature Version 4 Signing Process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
	Sigv4 interface{} `field:"optional" json:"sigv4" yaml:"sigv4"`
}

