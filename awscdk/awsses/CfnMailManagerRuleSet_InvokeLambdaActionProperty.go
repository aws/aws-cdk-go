package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invokeLambdaActionProperty := &InvokeLambdaActionProperty{
//   	FunctionArn: jsii.String("functionArn"),
//   	InvocationType: jsii.String("invocationType"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	RetryTimeMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html
//
type CfnMailManagerRuleSet_InvokeLambdaActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-functionarn
	//
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-invocationtype
	//
	InvocationType *string `field:"required" json:"invocationType" yaml:"invocationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-retrytimeminutes
	//
	RetryTimeMinutes *float64 `field:"optional" json:"retryTimeMinutes" yaml:"retryTimeMinutes"`
}

