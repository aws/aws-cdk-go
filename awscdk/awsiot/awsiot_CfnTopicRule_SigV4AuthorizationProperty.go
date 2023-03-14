package awsiot


// For more information, see [Signature Version 4 signing process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sigV4AuthorizationProperty := &sigV4AuthorizationProperty{
//   	roleArn: jsii.String("roleArn"),
//   	serviceName: jsii.String("serviceName"),
//   	signingRegion: jsii.String("signingRegion"),
//   }
//
type CfnTopicRule_SigV4AuthorizationProperty struct {
	// The ARN of the signing role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The service name to use while signing with Sig V4.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The signing region.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
}

