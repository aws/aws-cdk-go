package awselasticloadbalancingv2


// Properties for defining a `CfnListener`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerProps := &CfnListenerProps{
//   	DefaultActions: []interface{}{
//   		&ActionProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   				UserPoolArn: jsii.String("userPoolArn"),
//   				UserPoolClientId: jsii.String("userPoolClientId"),
//   				UserPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.String("sessionTimeout"),
//   			},
//   			AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				ClientId: jsii.String("clientId"),
//   				Issuer: jsii.String("issuer"),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   				UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				ClientSecret: jsii.String("clientSecret"),
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.String("sessionTimeout"),
//   				UseExistingClientSecret: jsii.Boolean(false),
//   			},
//   			FixedResponseConfig: &FixedResponseConfigProperty{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				ContentType: jsii.String("contentType"),
//   				MessageBody: jsii.String("messageBody"),
//   			},
//   			ForwardConfig: &ForwardConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupTupleProperty{
//   						TargetGroupArn: jsii.String("targetGroupArn"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   				TargetGroupStickinessConfig: &TargetGroupStickinessConfigProperty{
//   					DurationSeconds: jsii.Number(123),
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   			Order: jsii.Number(123),
//   			RedirectConfig: &RedirectConfigProperty{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				Host: jsii.String("host"),
//   				Path: jsii.String("path"),
//   				Port: jsii.String("port"),
//   				Protocol: jsii.String("protocol"),
//   				Query: jsii.String("query"),
//   			},
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//
//   	// the properties below are optional
//   	AlpnPolicy: []*string{
//   		jsii.String("alpnPolicy"),
//   	},
//   	Certificates: []interface{}{
//   		&CertificateProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	MutualAuthentication: &MutualAuthenticationProperty{
//   		IgnoreClientCertificateExpiry: jsii.Boolean(false),
//   		Mode: jsii.String("mode"),
//   		TrustStoreArn: jsii.String("trustStoreArn"),
//   	},
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SslPolicy: jsii.String("sslPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html
//
type CfnListenerProps struct {
	// The actions for the default rule. You cannot define a condition for a default rule.
	//
	// To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-defaultactions
	//
	DefaultActions interface{} `field:"required" json:"defaultActions" yaml:"defaultActions"`
	// The Amazon Resource Name (ARN) of the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-loadbalancerarn
	//
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-alpnpolicy
	//
	AlpnPolicy *[]*string `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// The default SSL server certificate for a secure listener.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//
	// To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-certificates
	//
	Certificates interface{} `field:"optional" json:"certificates" yaml:"certificates"`
	// The mutual authentication configuration information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-mutualauthentication
	//
	MutualAuthentication interface{} `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// The port on which the load balancer is listening.
	//
	// You cannot specify a port for a Gateway Load Balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//
	// Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-sslpolicy
	//
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

