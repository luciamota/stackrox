package processor

import (
	"bitbucket.org/stack-rox/apollo/central/notifier/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/notifications/notifiers"
)

const (
	alertChanSize     = 100
	benchmarkChanSize = 100
)

var (
	log = logging.LoggerForModule()
)

// Processor is the interface for processing benchmarks, notifiers, and policies.
type Processor interface {
	Start()
	ProcessAlert(alert *v1.Alert)
	ProcessBenchmark(schedule *v1.BenchmarkSchedule)
	RemoveNotifier(id string)
	UpdateNotifier(notifier notifiers.Notifier)
	GetIntegratedPolicies(notifierID string) (output []*v1.Policy)
	UpdatePolicy(policy *v1.Policy)
	RemovePolicy(policy *v1.Policy)
}

// New returns a new Processor
func New(storage store.Store) (Processor, error) {
	processor := &processorImpl{
		alertChan:           make(chan *v1.Alert, alertChanSize),
		benchmarkChan:       make(chan *v1.BenchmarkSchedule, benchmarkChanSize),
		notifiers:           make(map[string]notifiers.Notifier),
		notifiersToPolicies: make(map[string]map[string]*v1.Policy),
		storage:             storage,
	}
	err := processor.initializeNotifiers()
	return processor, err
}
