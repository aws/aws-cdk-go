package awsredshift


// The Amazon Redshift connect integration scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectProperty := &ConnectProperty{
//   	Authorization: jsii.String("authorization"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-connect.html
//
type CfnRedshiftIdcApplicationPropsMixin_ConnectProperty struct {
	// Determines whether the Amazon Redshift connect integration is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-connect.html#cfn-redshift-redshiftidcapplication-connect-authorization
	//
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
}

