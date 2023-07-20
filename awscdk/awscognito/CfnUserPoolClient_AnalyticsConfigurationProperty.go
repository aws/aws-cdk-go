package awscognito


// The Amazon Pinpoint analytics configuration necessary to collect metrics for a user pool.
//
// > In Regions where Amazon Pinpoint isn't available, user pools only support sending events to Amazon Pinpoint projects in us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyticsConfigurationProperty := &AnalyticsConfigurationProperty{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ApplicationId: jsii.String("applicationId"),
//   	ExternalId: jsii.String("externalId"),
//   	RoleArn: jsii.String("roleArn"),
//   	UserDataShared: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html
//
type CfnUserPoolClient_AnalyticsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon Pinpoint project.
	//
	// You can use the Amazon Pinpoint project for integration with the chosen user pool client. Amazon Cognito publishes events to the Amazon Pinpoint project that the app ARN declares.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html#cfn-cognito-userpoolclient-analyticsconfiguration-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The application ID for an Amazon Pinpoint application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html#cfn-cognito-userpoolclient-analyticsconfiguration-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The external ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html#cfn-cognito-userpoolclient-analyticsconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The ARN of an AWS Identity and Access Management role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html#cfn-cognito-userpoolclient-analyticsconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// If `UserDataShared` is `true` , Amazon Cognito includes user data in the events that it publishes to Amazon Pinpoint analytics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-analyticsconfiguration.html#cfn-cognito-userpoolclient-analyticsconfiguration-userdatashared
	//
	UserDataShared interface{} `field:"optional" json:"userDataShared" yaml:"userDataShared"`
}

