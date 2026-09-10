package awsiot


// Describes a rule.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
//
type CfnTopicRule_TopicRulePayloadProperty struct {
	// The actions associated with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The SQL statement used to query the topic.
	//
	// For more information, see [AWS IoT SQL Reference](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) in the *AWS IoT Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
	//
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	//
	// The default value is 2015-10-08.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
	//
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The description of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The action to take when an error occurs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-erroraction
	//
	ErrorAction interface{} `field:"optional" json:"errorAction" yaml:"errorAction"`
	// Specifies whether the rule is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
	//
	RuleDisabled interface{} `field:"optional" json:"ruleDisabled" yaml:"ruleDisabled"`
}

