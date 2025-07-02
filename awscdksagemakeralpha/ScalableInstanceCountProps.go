package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The properties of a scalable attribute representing task count.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   scalableInstanceCountProps := &ScalableInstanceCountProps{
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
// Experimental.
type ScalableInstanceCountProps struct {
	// Maximum capacity to scale to.
	// Experimental.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Default: 1.
	//
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Scalable dimension of the attribute.
	// Experimental.
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// Resource ID of the attribute.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to use for scaling.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Service namespace of the scalable attribute.
	// Experimental.
	ServiceNamespace awsapplicationautoscaling.ServiceNamespace `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

