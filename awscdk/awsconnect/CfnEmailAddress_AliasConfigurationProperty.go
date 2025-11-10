package awsconnect


// Configuration information of an email alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasConfigurationProperty := &AliasConfigurationProperty{
//   	EmailAddressArn: jsii.String("emailAddressArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-emailaddress-aliasconfiguration.html
//
type CfnEmailAddress_AliasConfigurationProperty struct {
	// The identifier of the email address alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-emailaddress-aliasconfiguration.html#cfn-connect-emailaddress-aliasconfiguration-emailaddressarn
	//
	EmailAddressArn *string `field:"required" json:"emailAddressArn" yaml:"emailAddressArn"`
}

