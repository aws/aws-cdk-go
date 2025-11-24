package mixinsawsappsync


// Describes properties used to specify configurations related to a source API.
//
// This is a property of the `AWS:AppSync:SourceApiAssociation` type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceApiAssociationConfigProperty := &SourceApiAssociationConfigProperty{
//   	MergeType: jsii.String("mergeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-sourceapiassociation-sourceapiassociationconfig.html
//
type CfnSourceApiAssociationPropsMixin_SourceApiAssociationConfigProperty struct {
	// The property that indicates which merging option is enabled in the source API association.
	//
	// Valid merge types are `MANUAL_MERGE` (default) and `AUTO_MERGE` . Manual merges are the default behavior and require the user to trigger any changes from the source APIs to the merged API manually. Auto merges subscribe the merged API to the changes performed on the source APIs so that any change in the source APIs are also made to the merged API. Auto merges use `MergedApiExecutionRoleArn` to perform merge operations.
	//
	// The following values are valid:
	//
	// `MANUAL_MERGE | AUTO_MERGE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-sourceapiassociation-sourceapiassociationconfig.html#cfn-appsync-sourceapiassociation-sourceapiassociationconfig-mergetype
	//
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

