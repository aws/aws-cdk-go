package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   invokeLambdaActionProperty := &InvokeLambdaActionProperty{
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	FunctionArn: jsii.String("functionArn"),
//   	InvocationType: jsii.String("invocationType"),
//   	RetryTimeMinutes: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html
//
type CfnMailManagerRuleSetPropsMixin_InvokeLambdaActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-functionarn
	//
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-invocationtype
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-retrytimeminutes
	//
	RetryTimeMinutes *float64 `field:"optional" json:"retryTimeMinutes" yaml:"retryTimeMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-invokelambdaaction.html#cfn-ses-mailmanagerruleset-invokelambdaaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

