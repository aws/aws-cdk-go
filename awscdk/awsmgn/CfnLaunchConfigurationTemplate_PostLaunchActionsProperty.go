package awsmgn


// Post launch actions to execute on the Test or Cutover instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postLaunchActionsProperty := &PostLaunchActionsProperty{
//   	CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	Deployment: jsii.String("deployment"),
//   	S3LogBucket: jsii.String("s3LogBucket"),
//   	S3OutputKeyPrefix: jsii.String("s3OutputKeyPrefix"),
//   	SsmDocuments: []interface{}{
//   		&SsmDocumentProperty{
//   			ActionName: jsii.String("actionName"),
//   			SsmDocumentName: jsii.String("ssmDocumentName"),
//
//   			// the properties below are optional
//   			ExternalParameters: map[string]interface{}{
//   				"externalParametersKey": &SsmExternalParameterProperty{
//   					"dynamicPath": jsii.String("dynamicPath"),
//   				},
//   			},
//   			MustSucceedForCutover: jsii.Boolean(false),
//   			Parameters: map[string]interface{}{
//   				"parametersKey": []interface{}{
//   					&SsmParameterStoreParameterProperty{
//   						"parameterName": jsii.String("parameterName"),
//   						"parameterType": jsii.String("parameterType"),
//   					},
//   				},
//   			},
//   			TimeoutSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html
//
type CfnLaunchConfigurationTemplate_PostLaunchActionsProperty struct {
	// AWS Systems Manager Command's CloudWatch log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions-cloudwatchloggroupname
	//
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Deployment type in which AWS Systems Manager Documents will be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions-deployment
	//
	Deployment *string `field:"optional" json:"deployment" yaml:"deployment"`
	// AWS Systems Manager Command's logs S3 log bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions-s3logbucket
	//
	S3LogBucket *string `field:"optional" json:"s3LogBucket" yaml:"s3LogBucket"`
	// AWS Systems Manager Command's logs S3 output key prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions-s3outputkeyprefix
	//
	S3OutputKeyPrefix *string `field:"optional" json:"s3OutputKeyPrefix" yaml:"s3OutputKeyPrefix"`
	// AWS Systems Manager Documents to execute, in order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-postlaunchactions.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions-ssmdocuments
	//
	SsmDocuments interface{} `field:"optional" json:"ssmDocuments" yaml:"ssmDocuments"`
}

