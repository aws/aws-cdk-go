package awscognito


// The Amazon Pinpoint analytics configuration necessary to collect metrics for a user pool.
//
// > In Regions where Amazon Pinpointisn't available, user pools only support sending events to Amazon Pinpoint projects in us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyticsConfigurationProperty := &analyticsConfigurationProperty{
//   	applicationArn: jsii.String("applicationArn"),
//   	applicationId: jsii.String("applicationId"),
//   	externalId: jsii.String("externalId"),
//   	roleArn: jsii.String("roleArn"),
//   	userDataShared: jsii.Boolean(false),
//   }
//
type CfnUserPoolClient_AnalyticsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon Pinpoint project.
	//
	// You can use the Amazon Pinpoint project for integration with the chosen user pool client. Amazon Cognito publishes events to the Amazon Pinpoint project that the app ARN declares.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The application ID for an Amazon Pinpoint application.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The external ID.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The ARN of an AWS Identity and Access Management role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// If `UserDataShared` is `true` , Amazon Cognito includes user data in the events that it publishes to Amazon Pinpoint analytics.
	UserDataShared interface{} `field:"optional" json:"userDataShared" yaml:"userDataShared"`
}

