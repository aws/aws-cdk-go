package awsiotwireless


// Information about a Sidewalk account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sidewalkAccountInfoProperty := &sidewalkAccountInfoProperty{
//   	appServerPrivateKey: jsii.String("appServerPrivateKey"),
//   }
//
type CfnPartnerAccount_SidewalkAccountInfoProperty struct {
	// The Sidewalk application server private key.
	//
	// The application server private key is a secret key, which you should handle in a similar way as you would an application password. You can protect the application server private key by storing the value in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	AppServerPrivateKey *string `field:"required" json:"appServerPrivateKey" yaml:"appServerPrivateKey"`
}

