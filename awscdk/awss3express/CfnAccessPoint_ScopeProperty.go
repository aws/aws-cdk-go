package awss3express


// You can use the access point scope to restrict access to specific prefixes, API operations, or a combination of both.
//
// For more information, see [Manage the scope of your access points for directory buckets.](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets-manage-scope.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeProperty := &ScopeProperty{
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	Prefixes: []*string{
//   		jsii.String("prefixes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-accesspoint-scope.html
//
type CfnAccessPoint_ScopeProperty struct {
	// You can include one or more API operations as permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-accesspoint-scope.html#cfn-s3express-accesspoint-scope-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// You can specify any amount of prefixes, but the total length of characters of all prefixes must be less than 256 bytes in size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-accesspoint-scope.html#cfn-s3express-accesspoint-scope-prefixes
	//
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

