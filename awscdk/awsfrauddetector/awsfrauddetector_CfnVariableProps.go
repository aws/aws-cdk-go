package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVariable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVariableProps := &cfnVariableProps{
//   	dataSource: jsii.String("dataSource"),
//   	dataType: jsii.String("dataType"),
//   	defaultValue: jsii.String("defaultValue"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	variableType: jsii.String("variableType"),
//   }
//
type CfnVariableProps struct {
	// The data source of the variable.
	//
	// Valid values: `EVENT | EXTERNAL_MODEL_SCORE`
	//
	// When defining a variable within a detector, you can only use the `EVENT` value for DataSource when the *Inline* property is set to true. If the *Inline* property is set false, you can use either `EVENT` or `MODEL_SCORE` for DataSource.
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The data type of the variable.
	//
	// Valid data types: `STRING | INTEGER | BOOLEAN | FLOAT`.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The default value of the variable.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// The name of the variable.
	//
	// Pattern: `^[0-9a-z_-]+$`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the variable.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the variable. For more information see [Variable types](https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types) .
	//
	// Valid Values: `AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT`.
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

