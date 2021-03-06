package types

import (
	. "github.com/flant/shell-operator/pkg/schedule_manager/types"

	"github.com/flant/shell-operator/pkg/kube_events_manager"
	"github.com/flant/shell-operator/pkg/validating_webhook"
)

type BindingType string

const (
	Schedule             BindingType = "schedule"
	OnStartup            BindingType = "onStartup"
	OnKubernetesEvent    BindingType = "kubernetes"
	KubernetesValidating BindingType = "kubernetesValidating"
)

// Types for effective binding configs
type CommonBindingConfig struct {
	BindingName  string
	AllowFailure bool
}

type OnStartupConfig struct {
	CommonBindingConfig
	Order float64
}

type ScheduleConfig struct {
	CommonBindingConfig
	ScheduleEntry        ScheduleEntry
	IncludeSnapshotsFrom []string
	Queue                string
	Group                string
}

type OnKubernetesEventConfig struct {
	CommonBindingConfig
	Monitor                      *kube_events_manager.MonitorConfig
	IncludeSnapshotsFrom         []string
	Queue                        string
	Group                        string
	ExecuteHookOnSynchronization bool
	WaitForSynchronization       bool
	KeepFullObjectsInMemory      bool
}

type ValidatingConfig struct {
	CommonBindingConfig

	Webhook *validating_webhook.ValidatingWebhookConfig

	IncludeSnapshotsFrom []string
	Group                string
}
