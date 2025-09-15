package awsapplicationautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a ScalableTableAttribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   baseScalableAttributeProps := &BaseScalableAttributeProps{
//   	Dimension: jsii.String("dimension"),
//   	MaxCapacity: jsii.Number(123),
//   	ResourceId: jsii.String("resourceId"),
//   	Role: role,
//   	ServiceNamespace: awscdk.Aws_applicationautoscaling.ServiceNamespace_ECS,
//
//   	// the properties below are optional
//   	MinCapacity: jsii.Number(123),
//   }
//
type BaseScalableAttributeProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Default: 1.
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Scalable dimension of the attribute.
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// Resource ID of the attribute.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to use for scaling.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Service namespace of the scalable attribute.
	ServiceNamespace ServiceNamespace `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

