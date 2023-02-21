// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Possible values for a resource's Removal Policy.
//
// The removal policy controls what happens to the resource if it stops being
// managed by CloudFormation. This can happen in one of three situations:
//
// - The resource is removed from the template, so CloudFormation stops managing it;
// - A change to the resource is made that requires it to be replaced, so CloudFormation stops
//    managing it;
// - The stack is deleted, so CloudFormation stops managing all resources in it.
//
// The Removal Policy applies to all above cases.
//
// Many stateful resources in the AWS Construct Library will accept a
// `removalPolicy` as a property, typically defaulting it to `RETAIN`.
//
// If the AWS Construct Library resource does not accept a `removalPolicy`
// argument, you can always configure it by using the escape hatch mechanism,
// as shown in the following example:
//
// ```ts
// declare const bucket: s3.Bucket;
//
// const cfnBucket = bucket.node.findChild('Resource') as CfnResource;
// cfnBucket.applyRemovalPolicy(RemovalPolicy.DESTROY);
// ```.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var api graphqlApi
//
//
//   user := iam.NewUser(this, jsii.String("User"))
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_2_3(),
//   	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
//   	FineGrainedAccessControl: &AdvancedSecurityOptions{
//   		MasterUserArn: user.UserArn,
//   	},
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EnforceHttps: jsii.Boolean(true),
//   })
//   ds := api.AddOpenSearchDataSource(jsii.String("ds"), domain)
//
//   ds.CreateResolver(jsii.String("QueryGetTestsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getTests"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jSON.stringify(map[string]interface{}{
//   		"version": jsii.String("2017-02-28"),
//   		"operation": jsii.String("GET"),
//   		"path": jsii.String("/id/post/_search"),
//   		"params": map[string]map[string]interface{}{
//   			"headers": map[string]interface{}{
//   			},
//   			"queryString": map[string]interface{}{
//   			},
//   			"body": map[string]*f64{
//   				"from": jsii.Number(0),
//   				"size": jsii.Number(50),
//   			},
//   		},
//   	})),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`[
//   	    #foreach($entry in $context.result.hits.hits)
//   	    #if( $velocityCount > 1 ) , #end
//   	    $utils.toJson($entry.get("_source"))
//   	    #end
//   	  ]`)),
//   })
//
type RemovalPolicy string

const (
	// This is the default removal policy.
	//
	// It means that when the resource is
	// removed from the app, it will be physically destroyed.
	RemovalPolicy_DESTROY RemovalPolicy = "DESTROY"
	// This uses the 'Retain' DeletionPolicy, which will cause the resource to be retained in the account, but orphaned from the stack.
	RemovalPolicy_RETAIN RemovalPolicy = "RETAIN"
	// This retention policy deletes the resource, but saves a snapshot of its data before deleting, so that it can be re-created later.
	//
	// Only available for some stateful resources,
	// like databases, EC2 volumes, etc.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	RemovalPolicy_SNAPSHOT RemovalPolicy = "SNAPSHOT"
)

