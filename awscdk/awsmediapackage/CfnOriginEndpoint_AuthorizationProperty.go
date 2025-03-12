package awsmediapackage


// Parameters for enabling CDN authorization on the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationProperty := &AuthorizationProperty{
//   	CdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   	SecretsRoleArn: jsii.String("secretsRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-authorization.html
//
type CfnOriginEndpoint_AuthorizationProperty struct {
	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that your Content Delivery Network (CDN) uses for authorization to access your endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-authorization.html#cfn-mediapackage-originendpoint-authorization-cdnidentifiersecret
	//
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows AWS Elemental MediaPackage to communicate with AWS Secrets Manager .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-authorization.html#cfn-mediapackage-originendpoint-authorization-secretsrolearn
	//
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

