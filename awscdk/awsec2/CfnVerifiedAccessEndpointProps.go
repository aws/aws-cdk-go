package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVerifiedAccessEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVerifiedAccessEndpointProps := &CfnVerifiedAccessEndpointProps{
//   	ApplicationDomain: jsii.String("applicationDomain"),
//   	AttachmentType: jsii.String("attachmentType"),
//   	DomainCertificateArn: jsii.String("domainCertificateArn"),
//   	EndpointDomainPrefix: jsii.String("endpointDomainPrefix"),
//   	EndpointType: jsii.String("endpointType"),
//   	VerifiedAccessGroupId: jsii.String("verifiedAccessGroupId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LoadBalancerOptions: &LoadBalancerOptionsProperty{
//   		LoadBalancerArn: jsii.String("loadBalancerArn"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	NetworkInterfaceOptions: &NetworkInterfaceOptionsProperty{
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyEnabled: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SseSpecification: &SseSpecificationProperty{
//   		CustomerManagedKeyEnabled: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html
//
type CfnVerifiedAccessEndpointProps struct {
	// The DNS name for users to reach your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-applicationdomain
	//
	ApplicationDomain *string `field:"required" json:"applicationDomain" yaml:"applicationDomain"`
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-attachmenttype
	//
	AttachmentType *string `field:"required" json:"attachmentType" yaml:"attachmentType"`
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-domaincertificatearn
	//
	DomainCertificateArn *string `field:"required" json:"domainCertificateArn" yaml:"domainCertificateArn"`
	// A custom identifier that is prepended to the DNS name that is generated for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-endpointdomainprefix
	//
	EndpointDomainPrefix *string `field:"required" json:"endpointDomainPrefix" yaml:"endpointDomainPrefix"`
	// The type of AWS Verified Access endpoint.
	//
	// Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-endpointtype
	//
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The ID of the AWS Verified Access group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-verifiedaccessgroupid
	//
	VerifiedAccessGroupId *string `field:"required" json:"verifiedAccessGroupId" yaml:"verifiedAccessGroupId"`
	// A description for the AWS Verified Access endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The load balancer details if creating the AWS Verified Access endpoint as `load-balancer` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-loadbalanceroptions
	//
	LoadBalancerOptions interface{} `field:"optional" json:"loadBalancerOptions" yaml:"loadBalancerOptions"`
	// The options for network-interface type endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-networkinterfaceoptions
	//
	NetworkInterfaceOptions interface{} `field:"optional" json:"networkInterfaceOptions" yaml:"networkInterfaceOptions"`
	// The Verified Access policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The status of the Verified Access policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-policyenabled
	//
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
	// The IDs of the security groups for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The options for additional server side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessendpoint.html#cfn-ec2-verifiedaccessendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

