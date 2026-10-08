package awsredshift


// A union structure that defines the scope of Amazon Redshift service integrations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   redshiftScopeUnionProperty := &RedshiftScopeUnionProperty{
//   	Connect: &ConnectProperty{
//   		Authorization: jsii.String("authorization"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-redshiftscopeunion.html
//
type CfnRedshiftIdcApplicationPropsMixin_RedshiftScopeUnionProperty struct {
	// The Amazon Redshift connect integration scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-redshiftscopeunion.html#cfn-redshift-redshiftidcapplication-redshiftscopeunion-connect
	//
	Connect interface{} `field:"optional" json:"connect" yaml:"connect"`
}

