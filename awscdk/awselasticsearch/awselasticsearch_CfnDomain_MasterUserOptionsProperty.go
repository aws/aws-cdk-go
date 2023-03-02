package awselasticsearch


// Specifies information about the master user. Required if you enabled the internal user database.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   masterUserOptionsProperty := &masterUserOptionsProperty{
//   	masterUserArn: jsii.String("masterUserArn"),
//   	masterUserName: jsii.String("masterUserName"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   }
//
type CfnDomain_MasterUserOptionsProperty struct {
	// ARN for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is false in `AdvancedSecurityOptions` .
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` .
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Password for the master user.
	//
	// Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions` .
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}

