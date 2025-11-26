package previewawssesmixins


// The action to deliver incoming emails to an Amazon Q Business application for indexing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deliverToQBusinessActionProperty := &DeliverToQBusinessActionProperty{
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	ApplicationId: jsii.String("applicationId"),
//   	IndexId: jsii.String("indexId"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.html
//
type CfnMailManagerRuleSetPropsMixin_DeliverToQBusinessActionProperty struct {
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the specified application has been deleted or the role lacks necessary permissions to call the `qbusiness:BatchPutDocument` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.html#cfn-ses-mailmanagerruleset-delivertoqbusinessaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// The unique identifier of the Amazon Q Business application instance where the email content will be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.html#cfn-ses-mailmanagerruleset-delivertoqbusinessaction-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The identifier of the knowledge base index within the Amazon Q Business application where the email content will be stored and indexed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.html#cfn-ses-mailmanagerruleset-delivertoqbusinessaction-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// The Amazon Resource Name (ARN) of the IAM Role to use while delivering to Amazon Q Business.
	//
	// This role must have access to the `qbusiness:BatchPutDocument` API for the given application and index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.html#cfn-ses-mailmanagerruleset-delivertoqbusinessaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

