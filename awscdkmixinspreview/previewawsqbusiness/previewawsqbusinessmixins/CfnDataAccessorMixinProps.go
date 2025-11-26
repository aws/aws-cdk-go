package previewawsqbusinessmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataAccessorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributeFilterProperty_ AttributeFilterProperty
//
//   cfnDataAccessorMixinProps := &CfnDataAccessorMixinProps{
//   	ActionConfigurations: []interface{}{
//   		&ActionConfigurationProperty{
//   			Action: jsii.String("action"),
//   			FilterConfiguration: &ActionFilterConfigurationProperty{
//   				DocumentAttributeFilter: &AttributeFilterProperty{
//   					AndAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   					ContainsAll: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					ContainsAny: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					EqualsTo: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					NotFilter: attributeFilterProperty_,
//   					OrAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	AuthenticationDetail: &DataAccessorAuthenticationDetailProperty{
//   		AuthenticationConfiguration: &DataAccessorAuthenticationConfigurationProperty{
//   			IdcTrustedTokenIssuerConfiguration: &DataAccessorIdcTrustedTokenIssuerConfigurationProperty{
//   				IdcTrustedTokenIssuerArn: jsii.String("idcTrustedTokenIssuerArn"),
//   			},
//   		},
//   		AuthenticationType: jsii.String("authenticationType"),
//   		ExternalIds: []*string{
//   			jsii.String("externalIds"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Principal: jsii.String("principal"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html
//
type CfnDataAccessorMixinProps struct {
	// A list of action configurations specifying the allowed actions and any associated filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-actionconfigurations
	//
	ActionConfigurations interface{} `field:"optional" json:"actionConfigurations" yaml:"actionConfigurations"`
	// The unique identifier of the Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The authentication configuration details for the data accessor.
	//
	// This specifies how the ISV authenticates when accessing data through this data accessor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-authenticationdetail
	//
	AuthenticationDetail interface{} `field:"optional" json:"authenticationDetail" yaml:"authenticationDetail"`
	// The friendly name of the data accessor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The Amazon Resource Name (ARN) of the IAM role for the ISV associated with this data accessor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// The tags to associate with the data accessor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

