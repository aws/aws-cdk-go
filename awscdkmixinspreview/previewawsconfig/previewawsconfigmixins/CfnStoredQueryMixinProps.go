package previewawsconfigmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStoredQueryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStoredQueryMixinProps := &CfnStoredQueryMixinProps{
//   	QueryDescription: jsii.String("queryDescription"),
//   	QueryExpression: jsii.String("queryExpression"),
//   	QueryName: jsii.String("queryName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-storedquery.html
//
type CfnStoredQueryMixinProps struct {
	// A unique description for the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-storedquery.html#cfn-config-storedquery-querydescription
	//
	QueryDescription *string `field:"optional" json:"queryDescription" yaml:"queryDescription"`
	// The expression of the query.
	//
	// For example, `SELECT resourceId, resourceType, supplementaryConfiguration.BucketVersioningConfiguration.status WHERE resourceType = 'AWS::S3::Bucket' AND supplementaryConfiguration.BucketVersioningConfiguration.status = 'Off'.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-storedquery.html#cfn-config-storedquery-queryexpression
	//
	QueryExpression *string `field:"optional" json:"queryExpression" yaml:"queryExpression"`
	// The name of the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-storedquery.html#cfn-config-storedquery-queryname
	//
	QueryName *string `field:"optional" json:"queryName" yaml:"queryName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-storedquery.html#cfn-config-storedquery-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

