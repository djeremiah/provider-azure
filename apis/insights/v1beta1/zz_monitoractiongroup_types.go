/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AADAuthInitParameters struct {

	// The identifier URI for AAD auth.
	IdentifierURI *string `json:"identifierUri,omitempty" tf:"identifier_uri,omitempty"`

	// The webhook application object Id for AAD auth.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id for AAD auth.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AADAuthObservation struct {

	// The identifier URI for AAD auth.
	IdentifierURI *string `json:"identifierUri,omitempty" tf:"identifier_uri,omitempty"`

	// The webhook application object Id for AAD auth.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id for AAD auth.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AADAuthParameters struct {

	// The identifier URI for AAD auth.
	// +kubebuilder:validation:Optional
	IdentifierURI *string `json:"identifierUri,omitempty" tf:"identifier_uri,omitempty"`

	// The webhook application object Id for AAD auth.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id for AAD auth.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ArmRoleReceiverInitParameters struct {

	// The name of the ARM role receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The arm role id.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type ArmRoleReceiverObservation struct {

	// The name of the ARM role receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The arm role id.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type ArmRoleReceiverParameters struct {

	// The name of the ARM role receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The arm role id.
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type AutomationRunBookReceiverInitParameters struct {

	// The automation account ID which holds this runbook and authenticates to Azure resources.
	AutomationAccountID *string `json:"automationAccountId,omitempty" tf:"automation_account_id,omitempty"`

	// Indicates whether this instance is global runbook.
	IsGlobalRunBook *bool `json:"isGlobalRunbook,omitempty" tf:"is_global_runbook,omitempty"`

	// The name of the automation runbook receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name for this runbook.
	RunBookName *string `json:"runbookName,omitempty" tf:"runbook_name,omitempty"`

	// The URI where webhooks should be sent.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`

	// The resource id for webhook linked to this runbook.
	WebhookResourceID *string `json:"webhookResourceId,omitempty" tf:"webhook_resource_id,omitempty"`
}

type AutomationRunBookReceiverObservation struct {

	// The automation account ID which holds this runbook and authenticates to Azure resources.
	AutomationAccountID *string `json:"automationAccountId,omitempty" tf:"automation_account_id,omitempty"`

	// Indicates whether this instance is global runbook.
	IsGlobalRunBook *bool `json:"isGlobalRunbook,omitempty" tf:"is_global_runbook,omitempty"`

	// The name of the automation runbook receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name for this runbook.
	RunBookName *string `json:"runbookName,omitempty" tf:"runbook_name,omitempty"`

	// The URI where webhooks should be sent.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`

	// The resource id for webhook linked to this runbook.
	WebhookResourceID *string `json:"webhookResourceId,omitempty" tf:"webhook_resource_id,omitempty"`
}

type AutomationRunBookReceiverParameters struct {

	// The automation account ID which holds this runbook and authenticates to Azure resources.
	// +kubebuilder:validation:Optional
	AutomationAccountID *string `json:"automationAccountId,omitempty" tf:"automation_account_id,omitempty"`

	// Indicates whether this instance is global runbook.
	// +kubebuilder:validation:Optional
	IsGlobalRunBook *bool `json:"isGlobalRunbook,omitempty" tf:"is_global_runbook,omitempty"`

	// The name of the automation runbook receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name for this runbook.
	// +kubebuilder:validation:Optional
	RunBookName *string `json:"runbookName,omitempty" tf:"runbook_name,omitempty"`

	// The URI where webhooks should be sent.
	// +kubebuilder:validation:Optional
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`

	// The resource id for webhook linked to this runbook.
	// +kubebuilder:validation:Optional
	WebhookResourceID *string `json:"webhookResourceId,omitempty" tf:"webhook_resource_id,omitempty"`
}

type AzureAppPushReceiverInitParameters struct {

	// The email address of the user signed into the mobile app who will receive push notifications from this receiver.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the Azure app push receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AzureAppPushReceiverObservation struct {

	// The email address of the user signed into the mobile app who will receive push notifications from this receiver.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the Azure app push receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AzureAppPushReceiverParameters struct {

	// The email address of the user signed into the mobile app who will receive push notifications from this receiver.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the Azure app push receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AzureFunctionReceiverInitParameters struct {

	// The Azure resource ID of the function app.
	FunctionAppResourceID *string `json:"functionAppResourceId,omitempty" tf:"function_app_resource_id,omitempty"`

	// The function name in the function app.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// The HTTP trigger url where HTTP request sent to.
	HTTPTriggerURL *string `json:"httpTriggerUrl,omitempty" tf:"http_trigger_url,omitempty"`

	// The name of the Azure Function receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type AzureFunctionReceiverObservation struct {

	// The Azure resource ID of the function app.
	FunctionAppResourceID *string `json:"functionAppResourceId,omitempty" tf:"function_app_resource_id,omitempty"`

	// The function name in the function app.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// The HTTP trigger url where HTTP request sent to.
	HTTPTriggerURL *string `json:"httpTriggerUrl,omitempty" tf:"http_trigger_url,omitempty"`

	// The name of the Azure Function receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type AzureFunctionReceiverParameters struct {

	// The Azure resource ID of the function app.
	// +kubebuilder:validation:Optional
	FunctionAppResourceID *string `json:"functionAppResourceId,omitempty" tf:"function_app_resource_id,omitempty"`

	// The function name in the function app.
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// The HTTP trigger url where HTTP request sent to.
	// +kubebuilder:validation:Optional
	HTTPTriggerURL *string `json:"httpTriggerUrl,omitempty" tf:"http_trigger_url,omitempty"`

	// The name of the Azure Function receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EmailReceiverInitParameters struct {

	// The email address of this receiver.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EmailReceiverObservation struct {

	// The email address of this receiver.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EmailReceiverParameters struct {

	// The email address of this receiver.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EventHubReceiverInitParameters struct {

	// The resource ID of the respective Event Hub.
	EventHubID *string `json:"eventHubId,omitempty" tf:"event_hub_id,omitempty"`

	// The name of the specific Event Hub queue.
	EventHubName *string `json:"eventHubName,omitempty" tf:"event_hub_name,omitempty"`

	// The namespace name of the Event Hub.
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" tf:"event_hub_namespace,omitempty"`

	// The name of the EventHub Receiver, must be unique within action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID for the subscription containing this Event Hub. Default to the subscription ID of the Action Group.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Tenant ID for the subscription containing this Event Hub.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Indicates whether to use common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EventHubReceiverObservation struct {

	// The resource ID of the respective Event Hub.
	EventHubID *string `json:"eventHubId,omitempty" tf:"event_hub_id,omitempty"`

	// The name of the specific Event Hub queue.
	EventHubName *string `json:"eventHubName,omitempty" tf:"event_hub_name,omitempty"`

	// The namespace name of the Event Hub.
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" tf:"event_hub_namespace,omitempty"`

	// The name of the EventHub Receiver, must be unique within action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID for the subscription containing this Event Hub. Default to the subscription ID of the Action Group.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Tenant ID for the subscription containing this Event Hub.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Indicates whether to use common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EventHubReceiverParameters struct {

	// The resource ID of the respective Event Hub.
	// +kubebuilder:validation:Optional
	EventHubID *string `json:"eventHubId,omitempty" tf:"event_hub_id,omitempty"`

	// The name of the specific Event Hub queue.
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventHubName,omitempty" tf:"event_hub_name,omitempty"`

	// The namespace name of the Event Hub.
	// +kubebuilder:validation:Optional
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" tf:"event_hub_namespace,omitempty"`

	// The name of the EventHub Receiver, must be unique within action group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID for the subscription containing this Event Hub. Default to the subscription ID of the Action Group.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Tenant ID for the subscription containing this Event Hub.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Indicates whether to use common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type ItsmReceiverInitParameters struct {

	// The unique connection identifier of the ITSM connection.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The name of the ITSM receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region of the workspace.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
	TicketConfiguration *string `json:"ticketConfiguration,omitempty" tf:"ticket_configuration,omitempty"`

	// The Azure Log Analytics workspace ID where this connection is defined. Format is <subscription id>|<workspace id>, for example 00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type ItsmReceiverObservation struct {

	// The unique connection identifier of the ITSM connection.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The name of the ITSM receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region of the workspace.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
	TicketConfiguration *string `json:"ticketConfiguration,omitempty" tf:"ticket_configuration,omitempty"`

	// The Azure Log Analytics workspace ID where this connection is defined. Format is <subscription id>|<workspace id>, for example 00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type ItsmReceiverParameters struct {

	// The unique connection identifier of the ITSM connection.
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The name of the ITSM receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region of the workspace.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
	// +kubebuilder:validation:Optional
	TicketConfiguration *string `json:"ticketConfiguration,omitempty" tf:"ticket_configuration,omitempty"`

	// The Azure Log Analytics workspace ID where this connection is defined. Format is <subscription id>|<workspace id>, for example 00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000.
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type LogicAppReceiverInitParameters struct {

	// The callback url where HTTP request sent to.
	CallbackURL *string `json:"callbackUrl,omitempty" tf:"callback_url,omitempty"`

	// The name of the logic app receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Azure resource ID of the logic app.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type LogicAppReceiverObservation struct {

	// The callback url where HTTP request sent to.
	CallbackURL *string `json:"callbackUrl,omitempty" tf:"callback_url,omitempty"`

	// The name of the logic app receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Azure resource ID of the logic app.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type LogicAppReceiverParameters struct {

	// The callback url where HTTP request sent to.
	// +kubebuilder:validation:Optional
	CallbackURL *string `json:"callbackUrl,omitempty" tf:"callback_url,omitempty"`

	// The name of the logic app receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Azure resource ID of the logic app.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type MonitorActionGroupInitParameters struct {

	// One or more arm_role_receiver blocks as defined below.
	ArmRoleReceiver []ArmRoleReceiverInitParameters `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver,omitempty"`

	// One or more automation_runbook_receiver blocks as defined below.
	AutomationRunBookReceiver []AutomationRunBookReceiverInitParameters `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver,omitempty"`

	// One or more azure_app_push_receiver blocks as defined below.
	AzureAppPushReceiver []AzureAppPushReceiverInitParameters `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver,omitempty"`

	// One or more azure_function_receiver blocks as defined below.
	AzureFunctionReceiver []AzureFunctionReceiverInitParameters `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver,omitempty"`

	// One or more email_receiver blocks as defined below.
	EmailReceiver []EmailReceiverInitParameters `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`

	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more event_hub_receiver blocks as defined below.
	EventHubReceiver []EventHubReceiverInitParameters `json:"eventHubReceiver,omitempty" tf:"event_hub_receiver,omitempty"`

	// One or more itsm_receiver blocks as defined below.
	ItsmReceiver []ItsmReceiverInitParameters `json:"itsmReceiver,omitempty" tf:"itsm_receiver,omitempty"`

	// The Azure Region where the Action Group should exist. Changing this forces a new Action Group to be created. Defaults to global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more logic_app_receiver blocks as defined below.
	LogicAppReceiver []LogicAppReceiverInitParameters `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver,omitempty"`

	// One or more sms_receiver blocks as defined below.
	SMSReceiver []SMSReceiverInitParameters `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`

	// The short name of the action group. This will be used in SMS messages.
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more voice_receiver blocks as defined below.
	VoiceReceiver []VoiceReceiverInitParameters `json:"voiceReceiver,omitempty" tf:"voice_receiver,omitempty"`

	// One or more webhook_receiver blocks as defined below.
	WebhookReceiver []WebhookReceiverInitParameters `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type MonitorActionGroupObservation struct {

	// One or more arm_role_receiver blocks as defined below.
	ArmRoleReceiver []ArmRoleReceiverObservation `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver,omitempty"`

	// One or more automation_runbook_receiver blocks as defined below.
	AutomationRunBookReceiver []AutomationRunBookReceiverObservation `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver,omitempty"`

	// One or more azure_app_push_receiver blocks as defined below.
	AzureAppPushReceiver []AzureAppPushReceiverObservation `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver,omitempty"`

	// One or more azure_function_receiver blocks as defined below.
	AzureFunctionReceiver []AzureFunctionReceiverObservation `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver,omitempty"`

	// One or more email_receiver blocks as defined below.
	EmailReceiver []EmailReceiverObservation `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`

	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more event_hub_receiver blocks as defined below.
	EventHubReceiver []EventHubReceiverObservation `json:"eventHubReceiver,omitempty" tf:"event_hub_receiver,omitempty"`

	// The ID of the Action Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more itsm_receiver blocks as defined below.
	ItsmReceiver []ItsmReceiverObservation `json:"itsmReceiver,omitempty" tf:"itsm_receiver,omitempty"`

	// The Azure Region where the Action Group should exist. Changing this forces a new Action Group to be created. Defaults to global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more logic_app_receiver blocks as defined below.
	LogicAppReceiver []LogicAppReceiverObservation `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver,omitempty"`

	// The name of the resource group in which to create the Action Group instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// One or more sms_receiver blocks as defined below.
	SMSReceiver []SMSReceiverObservation `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`

	// The short name of the action group. This will be used in SMS messages.
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more voice_receiver blocks as defined below.
	VoiceReceiver []VoiceReceiverObservation `json:"voiceReceiver,omitempty" tf:"voice_receiver,omitempty"`

	// One or more webhook_receiver blocks as defined below.
	WebhookReceiver []WebhookReceiverObservation `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type MonitorActionGroupParameters struct {

	// One or more arm_role_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	ArmRoleReceiver []ArmRoleReceiverParameters `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver,omitempty"`

	// One or more automation_runbook_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AutomationRunBookReceiver []AutomationRunBookReceiverParameters `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver,omitempty"`

	// One or more azure_app_push_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AzureAppPushReceiver []AzureAppPushReceiverParameters `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver,omitempty"`

	// One or more azure_function_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AzureFunctionReceiver []AzureFunctionReceiverParameters `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver,omitempty"`

	// One or more email_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	EmailReceiver []EmailReceiverParameters `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`

	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more event_hub_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	EventHubReceiver []EventHubReceiverParameters `json:"eventHubReceiver,omitempty" tf:"event_hub_receiver,omitempty"`

	// One or more itsm_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	ItsmReceiver []ItsmReceiverParameters `json:"itsmReceiver,omitempty" tf:"itsm_receiver,omitempty"`

	// The Azure Region where the Action Group should exist. Changing this forces a new Action Group to be created. Defaults to global.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more logic_app_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	LogicAppReceiver []LogicAppReceiverParameters `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver,omitempty"`

	// The name of the resource group in which to create the Action Group instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more sms_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	SMSReceiver []SMSReceiverParameters `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`

	// The short name of the action group. This will be used in SMS messages.
	// +kubebuilder:validation:Optional
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more voice_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	VoiceReceiver []VoiceReceiverParameters `json:"voiceReceiver,omitempty" tf:"voice_receiver,omitempty"`

	// One or more webhook_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	WebhookReceiver []WebhookReceiverParameters `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type SMSReceiverInitParameters struct {

	// The country code of the SMS receiver.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the SMS receiver.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type SMSReceiverObservation struct {

	// The country code of the SMS receiver.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the SMS receiver.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type SMSReceiverParameters struct {

	// The country code of the SMS receiver.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the SMS receiver.
	// +kubebuilder:validation:Optional
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type VoiceReceiverInitParameters struct {

	// The country code of the voice receiver.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the voice receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the voice receiver.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type VoiceReceiverObservation struct {

	// The country code of the voice receiver.
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the voice receiver.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the voice receiver.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type VoiceReceiverParameters struct {

	// The country code of the voice receiver.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the voice receiver.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The phone number of the voice receiver.
	// +kubebuilder:validation:Optional
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type WebhookReceiverInitParameters struct {

	// The aad_auth block as defined below
	AADAuth []AADAuthInitParameters `json:"aadAuth,omitempty" tf:"aad_auth,omitempty"`

	// The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The URI where webhooks should be sent.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type WebhookReceiverObservation struct {

	// The aad_auth block as defined below
	AADAuth []AADAuthObservation `json:"aadAuth,omitempty" tf:"aad_auth,omitempty"`

	// The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The URI where webhooks should be sent.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type WebhookReceiverParameters struct {

	// The aad_auth block as defined below
	// +kubebuilder:validation:Optional
	AADAuth []AADAuthParameters `json:"aadAuth,omitempty" tf:"aad_auth,omitempty"`

	// The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The URI where webhooks should be sent.
	// +kubebuilder:validation:Optional
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

// MonitorActionGroupSpec defines the desired state of MonitorActionGroup
type MonitorActionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MonitorActionGroupInitParameters `json:"initProvider,omitempty"`
}

// MonitorActionGroupStatus defines the observed state of MonitorActionGroup.
type MonitorActionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroup is the Schema for the MonitorActionGroups API. Manages an Action Group within Azure Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shortName) || has(self.initProvider.shortName)",message="shortName is a required parameter"
	Spec   MonitorActionGroupSpec   `json:"spec"`
	Status MonitorActionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroupList contains a list of MonitorActionGroups
type MonitorActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionGroup `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionGroup_Kind             = "MonitorActionGroup"
	MonitorActionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorActionGroup_Kind}.String()
	MonitorActionGroup_KindAPIVersion   = MonitorActionGroup_Kind + "." + CRDGroupVersion.String()
	MonitorActionGroup_GroupVersionKind = CRDGroupVersion.WithKind(MonitorActionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionGroup{}, &MonitorActionGroupList{})
}
