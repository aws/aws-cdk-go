package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Domain name configuration for AppSync.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myDomainName := "api.example.com"
//   certificate := acm.NewCertificate(this, jsii.String("cert"), &CertificateProps{
//   	DomainName: myDomainName,
//   })
//
//   apiKeyProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
//   }
//
//   api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
//   	ApiName: jsii.String("Api"),
//   	OwnerContact: jsii.String("OwnerContact"),
//   	AuthorizationConfig: &EventApiAuthConfig{
//   		AuthProviders: []appSyncAuthProvider{
//   			apiKeyProvider,
//   		},
//   		ConnectionAuthModeTypes: []appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   	},
//   	// Custom Domain Settings
//   	DomainName: &AppSyncDomainOptions{
//   		Certificate: *Certificate,
//   		DomainName: myDomainName,
//   	},
//   })
//
//   // You can get custom HTTP/Realtime endpoint
//   // You can get custom HTTP/Realtime endpoint
//   awscdk.NewCfnOutput(this, jsii.String("AWS AppSync Events HTTP endpoint"), &CfnOutputProps{
//   	Value: api.customHttpEndpoint,
//   })
//   awscdk.NewCfnOutput(this, jsii.String("AWS AppSync Events Realtime endpoint"), &CfnOutputProps{
//   	Value: api.customRealtimeEndpoint,
//   })
//
type AppSyncDomainOptions struct {
	// The certificate to use with the domain name.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The actual domain name.
	//
	// For example, `api.example.com`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

