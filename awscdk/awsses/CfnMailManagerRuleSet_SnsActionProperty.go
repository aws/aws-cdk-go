package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsActionProperty := &SnsActionProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	TopicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	Encoding: jsii.String("encoding"),
//   	PayloadType: jsii.String("payloadType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html
//
type CfnMailManagerRuleSet_SnsActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-encoding
	//
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-payloadtype
	//
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
}

