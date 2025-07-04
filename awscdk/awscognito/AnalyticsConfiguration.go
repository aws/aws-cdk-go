package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint"
)

// The settings for Amazon Pinpoint analytics configuration.
//
// With an analytics configuration, your application can collect user-activity metrics for user notifications with an Amazon Pinpoint campaign.
// Amazon Pinpoint isn't available in all AWS Regions.
// For a list of available Regions, see Amazon Cognito and Amazon Pinpoint Region availability: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-pinpoint-integration.html#cognito-user-pools-find-region-mappings.
//
// Example:
//   import pinpoint "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPool userPool
//   var pinpointApp cfnApp
//   var pinpointRole role
//
//
//   cognito.NewUserPoolClient(this, jsii.String("Client"), &UserPoolClientProps{
//   	UserPool: UserPool,
//   	Analytics: &AnalyticsConfiguration{
//   		// Your Pinpoint project
//   		Application: pinpointApp,
//
//   		// Whether to include user data in analytics events
//   		ShareUserData: jsii.Boolean(true),
//   	},
//   })
//
type AnalyticsConfiguration struct {
	// The Amazon Pinpoint project that you want to connect to your user pool app client.
	//
	// Amazon Cognito publishes events to the Amazon Pinpoint project.
	// You can also configure your application to pass an endpoint ID in the `AnalyticsMetadata` parameter of sign-in operations.
	// The endpoint ID is information about the destination for push notifications.
	// Default: - no configuration, you need to specify either `application` or all of `applicationId`, `externalId`, and `role`.
	//
	Application awspinpoint.CfnApp `field:"optional" json:"application" yaml:"application"`
	// Your Amazon Pinpoint project ID.
	// Default: - no configuration, you need to specify either this property along with `externalId` and `role` or `application`.
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The external ID of the role that Amazon Cognito assumes to send analytics data to Amazon Pinpoint.
	//
	// More info here: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	// Default: - no configuration, you need to specify either this property along with `applicationId` and `role` or `application`.
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The IAM role that has the permissions required for Amazon Cognito to publish events to Amazon Pinpoint analytics.
	// Default: - no configuration, you need to specify either this property along with `applicationId` and `externalId` or `application`.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// If `true`, Amazon Cognito includes user data in the events that it publishes to Amazon Pinpoint analytics.
	// Default: - false.
	//
	ShareUserData *bool `field:"optional" json:"shareUserData" yaml:"shareUserData"`
}

