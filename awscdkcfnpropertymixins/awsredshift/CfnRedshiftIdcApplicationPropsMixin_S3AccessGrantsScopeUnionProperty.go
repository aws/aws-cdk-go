package awsredshift


// A list of scopes set up for S3 Access Grants integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3AccessGrantsScopeUnionProperty := &S3AccessGrantsScopeUnionProperty{
//   	ReadWriteAccess: &ReadWriteAccessProperty{
//   		Authorization: jsii.String("authorization"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-s3accessgrantsscopeunion.html
//
type CfnRedshiftIdcApplicationPropsMixin_S3AccessGrantsScopeUnionProperty struct {
	// The S3 Access Grants scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-s3accessgrantsscopeunion.html#cfn-redshift-redshiftidcapplication-s3accessgrantsscopeunion-readwriteaccess
	//
	ReadWriteAccess interface{} `field:"optional" json:"readWriteAccess" yaml:"readWriteAccess"`
}

