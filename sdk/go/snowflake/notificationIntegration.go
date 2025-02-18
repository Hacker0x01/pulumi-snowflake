// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-snowflake/sdk/go/snowflake"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := snowflake.NewNotificationIntegration(ctx, "integration", &snowflake.NotificationIntegrationArgs{
// 			AwsSnsRoleArn:               pulumi.String("..."),
// 			AwsSnsTopicArn:              pulumi.String("..."),
// 			AwsSqsArn:                   pulumi.String("..."),
// 			AwsSqsRoleArn:               pulumi.String("..."),
// 			AzureStorageQueuePrimaryUri: pulumi.String("..."),
// 			AzureTenantId:               pulumi.String("..."),
// 			Comment:                     pulumi.String("A notification integration."),
// 			Direction:                   pulumi.String("OUTBOUND"),
// 			Enabled:                     pulumi.Bool(true),
// 			NotificationProvider:        pulumi.String("AWS_SNS"),
// 			Type:                        pulumi.String("QUEUE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import snowflake:index/notificationIntegration:NotificationIntegration example name
// ```
type NotificationIntegration struct {
	pulumi.CustomResourceState

	// The external ID that Snowflake will use when assuming the AWS role
	AwsSnsExternalId pulumi.StringOutput `pulumi:"awsSnsExternalId"`
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSnsIamUserArn pulumi.StringOutput `pulumi:"awsSnsIamUserArn"`
	// AWS IAM role ARN for notification integration to assume
	AwsSnsRoleArn pulumi.StringPtrOutput `pulumi:"awsSnsRoleArn"`
	// AWS SNS Topic ARN for notification integration to connect to
	AwsSnsTopicArn pulumi.StringPtrOutput `pulumi:"awsSnsTopicArn"`
	// AWS SQS queue ARN for notification integration to connect to
	AwsSqsArn pulumi.StringPtrOutput `pulumi:"awsSqsArn"`
	// The external ID that Snowflake will use when assuming the AWS role
	AwsSqsExternalId pulumi.StringOutput `pulumi:"awsSqsExternalId"`
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSqsIamUserArn pulumi.StringOutput `pulumi:"awsSqsIamUserArn"`
	// AWS IAM role ARN for notification integration to assume
	AwsSqsRoleArn pulumi.StringPtrOutput `pulumi:"awsSqsRoleArn"`
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications
	AzureStorageQueuePrimaryUri pulumi.StringPtrOutput `pulumi:"azureStorageQueuePrimaryUri"`
	// The ID of the Azure Active Directory tenant used for identity management
	AzureTenantId pulumi.StringPtrOutput `pulumi:"azureTenantId"`
	// A comment for the integration
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Date and time when the notification integration was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	Enabled   pulumi.BoolPtrOutput   `pulumi:"enabled"`
	// The GCP service account identifier that Snowflake will use when assuming the GCP role
	GcpPubsubServiceAccount pulumi.StringOutput `pulumi:"gcpPubsubServiceAccount"`
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	GcpPubsubSubscriptionName pulumi.StringPtrOutput `pulumi:"gcpPubsubSubscriptionName"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	// The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS*SQS, AWS*SNS)
	NotificationProvider pulumi.StringPtrOutput `pulumi:"notificationProvider"`
	// A type of integration
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewNotificationIntegration registers a new resource with the given unique name, arguments, and options.
func NewNotificationIntegration(ctx *pulumi.Context,
	name string, args *NotificationIntegrationArgs, opts ...pulumi.ResourceOption) (*NotificationIntegration, error) {
	if args == nil {
		args = &NotificationIntegrationArgs{}
	}

	var resource NotificationIntegration
	err := ctx.RegisterResource("snowflake:index/notificationIntegration:NotificationIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationIntegration gets an existing NotificationIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationIntegrationState, opts ...pulumi.ResourceOption) (*NotificationIntegration, error) {
	var resource NotificationIntegration
	err := ctx.ReadResource("snowflake:index/notificationIntegration:NotificationIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationIntegration resources.
type notificationIntegrationState struct {
	// The external ID that Snowflake will use when assuming the AWS role
	AwsSnsExternalId *string `pulumi:"awsSnsExternalId"`
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSnsIamUserArn *string `pulumi:"awsSnsIamUserArn"`
	// AWS IAM role ARN for notification integration to assume
	AwsSnsRoleArn *string `pulumi:"awsSnsRoleArn"`
	// AWS SNS Topic ARN for notification integration to connect to
	AwsSnsTopicArn *string `pulumi:"awsSnsTopicArn"`
	// AWS SQS queue ARN for notification integration to connect to
	AwsSqsArn *string `pulumi:"awsSqsArn"`
	// The external ID that Snowflake will use when assuming the AWS role
	AwsSqsExternalId *string `pulumi:"awsSqsExternalId"`
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSqsIamUserArn *string `pulumi:"awsSqsIamUserArn"`
	// AWS IAM role ARN for notification integration to assume
	AwsSqsRoleArn *string `pulumi:"awsSqsRoleArn"`
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications
	AzureStorageQueuePrimaryUri *string `pulumi:"azureStorageQueuePrimaryUri"`
	// The ID of the Azure Active Directory tenant used for identity management
	AzureTenantId *string `pulumi:"azureTenantId"`
	// A comment for the integration
	Comment *string `pulumi:"comment"`
	// Date and time when the notification integration was created.
	CreatedOn *string `pulumi:"createdOn"`
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
	Direction *string `pulumi:"direction"`
	Enabled   *bool   `pulumi:"enabled"`
	// The GCP service account identifier that Snowflake will use when assuming the GCP role
	GcpPubsubServiceAccount *string `pulumi:"gcpPubsubServiceAccount"`
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	GcpPubsubSubscriptionName *string `pulumi:"gcpPubsubSubscriptionName"`
	Name                      *string `pulumi:"name"`
	// The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS*SQS, AWS*SNS)
	NotificationProvider *string `pulumi:"notificationProvider"`
	// A type of integration
	Type *string `pulumi:"type"`
}

type NotificationIntegrationState struct {
	// The external ID that Snowflake will use when assuming the AWS role
	AwsSnsExternalId pulumi.StringPtrInput
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSnsIamUserArn pulumi.StringPtrInput
	// AWS IAM role ARN for notification integration to assume
	AwsSnsRoleArn pulumi.StringPtrInput
	// AWS SNS Topic ARN for notification integration to connect to
	AwsSnsTopicArn pulumi.StringPtrInput
	// AWS SQS queue ARN for notification integration to connect to
	AwsSqsArn pulumi.StringPtrInput
	// The external ID that Snowflake will use when assuming the AWS role
	AwsSqsExternalId pulumi.StringPtrInput
	// The Snowflake user that will attempt to assume the AWS role.
	AwsSqsIamUserArn pulumi.StringPtrInput
	// AWS IAM role ARN for notification integration to assume
	AwsSqsRoleArn pulumi.StringPtrInput
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications
	AzureStorageQueuePrimaryUri pulumi.StringPtrInput
	// The ID of the Azure Active Directory tenant used for identity management
	AzureTenantId pulumi.StringPtrInput
	// A comment for the integration
	Comment pulumi.StringPtrInput
	// Date and time when the notification integration was created.
	CreatedOn pulumi.StringPtrInput
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
	Direction pulumi.StringPtrInput
	Enabled   pulumi.BoolPtrInput
	// The GCP service account identifier that Snowflake will use when assuming the GCP role
	GcpPubsubServiceAccount pulumi.StringPtrInput
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	GcpPubsubSubscriptionName pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	// The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS*SQS, AWS*SNS)
	NotificationProvider pulumi.StringPtrInput
	// A type of integration
	Type pulumi.StringPtrInput
}

func (NotificationIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationIntegrationState)(nil)).Elem()
}

type notificationIntegrationArgs struct {
	// AWS IAM role ARN for notification integration to assume
	AwsSnsRoleArn *string `pulumi:"awsSnsRoleArn"`
	// AWS SNS Topic ARN for notification integration to connect to
	AwsSnsTopicArn *string `pulumi:"awsSnsTopicArn"`
	// AWS SQS queue ARN for notification integration to connect to
	AwsSqsArn *string `pulumi:"awsSqsArn"`
	// AWS IAM role ARN for notification integration to assume
	AwsSqsRoleArn *string `pulumi:"awsSqsRoleArn"`
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications
	AzureStorageQueuePrimaryUri *string `pulumi:"azureStorageQueuePrimaryUri"`
	// The ID of the Azure Active Directory tenant used for identity management
	AzureTenantId *string `pulumi:"azureTenantId"`
	// A comment for the integration
	Comment *string `pulumi:"comment"`
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
	Direction *string `pulumi:"direction"`
	Enabled   *bool   `pulumi:"enabled"`
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	GcpPubsubSubscriptionName *string `pulumi:"gcpPubsubSubscriptionName"`
	Name                      *string `pulumi:"name"`
	// The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS*SQS, AWS*SNS)
	NotificationProvider *string `pulumi:"notificationProvider"`
	// A type of integration
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a NotificationIntegration resource.
type NotificationIntegrationArgs struct {
	// AWS IAM role ARN for notification integration to assume
	AwsSnsRoleArn pulumi.StringPtrInput
	// AWS SNS Topic ARN for notification integration to connect to
	AwsSnsTopicArn pulumi.StringPtrInput
	// AWS SQS queue ARN for notification integration to connect to
	AwsSqsArn pulumi.StringPtrInput
	// AWS IAM role ARN for notification integration to assume
	AwsSqsRoleArn pulumi.StringPtrInput
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications
	AzureStorageQueuePrimaryUri pulumi.StringPtrInput
	// The ID of the Azure Active Directory tenant used for identity management
	AzureTenantId pulumi.StringPtrInput
	// A comment for the integration
	Comment pulumi.StringPtrInput
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
	Direction pulumi.StringPtrInput
	Enabled   pulumi.BoolPtrInput
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	GcpPubsubSubscriptionName pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	// The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS*SQS, AWS*SNS)
	NotificationProvider pulumi.StringPtrInput
	// A type of integration
	Type pulumi.StringPtrInput
}

func (NotificationIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationIntegrationArgs)(nil)).Elem()
}

type NotificationIntegrationInput interface {
	pulumi.Input

	ToNotificationIntegrationOutput() NotificationIntegrationOutput
	ToNotificationIntegrationOutputWithContext(ctx context.Context) NotificationIntegrationOutput
}

func (*NotificationIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationIntegration)(nil)).Elem()
}

func (i *NotificationIntegration) ToNotificationIntegrationOutput() NotificationIntegrationOutput {
	return i.ToNotificationIntegrationOutputWithContext(context.Background())
}

func (i *NotificationIntegration) ToNotificationIntegrationOutputWithContext(ctx context.Context) NotificationIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationIntegrationOutput)
}

// NotificationIntegrationArrayInput is an input type that accepts NotificationIntegrationArray and NotificationIntegrationArrayOutput values.
// You can construct a concrete instance of `NotificationIntegrationArrayInput` via:
//
//          NotificationIntegrationArray{ NotificationIntegrationArgs{...} }
type NotificationIntegrationArrayInput interface {
	pulumi.Input

	ToNotificationIntegrationArrayOutput() NotificationIntegrationArrayOutput
	ToNotificationIntegrationArrayOutputWithContext(context.Context) NotificationIntegrationArrayOutput
}

type NotificationIntegrationArray []NotificationIntegrationInput

func (NotificationIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationIntegration)(nil)).Elem()
}

func (i NotificationIntegrationArray) ToNotificationIntegrationArrayOutput() NotificationIntegrationArrayOutput {
	return i.ToNotificationIntegrationArrayOutputWithContext(context.Background())
}

func (i NotificationIntegrationArray) ToNotificationIntegrationArrayOutputWithContext(ctx context.Context) NotificationIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationIntegrationArrayOutput)
}

// NotificationIntegrationMapInput is an input type that accepts NotificationIntegrationMap and NotificationIntegrationMapOutput values.
// You can construct a concrete instance of `NotificationIntegrationMapInput` via:
//
//          NotificationIntegrationMap{ "key": NotificationIntegrationArgs{...} }
type NotificationIntegrationMapInput interface {
	pulumi.Input

	ToNotificationIntegrationMapOutput() NotificationIntegrationMapOutput
	ToNotificationIntegrationMapOutputWithContext(context.Context) NotificationIntegrationMapOutput
}

type NotificationIntegrationMap map[string]NotificationIntegrationInput

func (NotificationIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationIntegration)(nil)).Elem()
}

func (i NotificationIntegrationMap) ToNotificationIntegrationMapOutput() NotificationIntegrationMapOutput {
	return i.ToNotificationIntegrationMapOutputWithContext(context.Background())
}

func (i NotificationIntegrationMap) ToNotificationIntegrationMapOutputWithContext(ctx context.Context) NotificationIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationIntegrationMapOutput)
}

type NotificationIntegrationOutput struct{ *pulumi.OutputState }

func (NotificationIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationIntegration)(nil)).Elem()
}

func (o NotificationIntegrationOutput) ToNotificationIntegrationOutput() NotificationIntegrationOutput {
	return o
}

func (o NotificationIntegrationOutput) ToNotificationIntegrationOutputWithContext(ctx context.Context) NotificationIntegrationOutput {
	return o
}

type NotificationIntegrationArrayOutput struct{ *pulumi.OutputState }

func (NotificationIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationIntegration)(nil)).Elem()
}

func (o NotificationIntegrationArrayOutput) ToNotificationIntegrationArrayOutput() NotificationIntegrationArrayOutput {
	return o
}

func (o NotificationIntegrationArrayOutput) ToNotificationIntegrationArrayOutputWithContext(ctx context.Context) NotificationIntegrationArrayOutput {
	return o
}

func (o NotificationIntegrationArrayOutput) Index(i pulumi.IntInput) NotificationIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotificationIntegration {
		return vs[0].([]*NotificationIntegration)[vs[1].(int)]
	}).(NotificationIntegrationOutput)
}

type NotificationIntegrationMapOutput struct{ *pulumi.OutputState }

func (NotificationIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationIntegration)(nil)).Elem()
}

func (o NotificationIntegrationMapOutput) ToNotificationIntegrationMapOutput() NotificationIntegrationMapOutput {
	return o
}

func (o NotificationIntegrationMapOutput) ToNotificationIntegrationMapOutputWithContext(ctx context.Context) NotificationIntegrationMapOutput {
	return o
}

func (o NotificationIntegrationMapOutput) MapIndex(k pulumi.StringInput) NotificationIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotificationIntegration {
		return vs[0].(map[string]*NotificationIntegration)[vs[1].(string)]
	}).(NotificationIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationIntegrationInput)(nil)).Elem(), &NotificationIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationIntegrationArrayInput)(nil)).Elem(), NotificationIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationIntegrationMapInput)(nil)).Elem(), NotificationIntegrationMap{})
	pulumi.RegisterOutputType(NotificationIntegrationOutput{})
	pulumi.RegisterOutputType(NotificationIntegrationArrayOutput{})
	pulumi.RegisterOutputType(NotificationIntegrationMapOutput{})
}
