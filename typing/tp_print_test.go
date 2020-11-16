package typing

import (
	"github.com/prometheus/alertmanager/types"
	"github.com/prometheus/common/model"
	"testing"
	"time"
)

func TestOutAlert(t *testing.T) {
	label := map[model.LabelName]model.LabelValue{}
	label["a"] = "a"

	alert := &types.Alert{
		Alert:     model.Alert{Labels: label},
		UpdatedAt: time.Time{},
	}
	OutAlert([]*types.Alert{alert})
}

func TestPrintAlert(t *testing.T) {
	label := map[model.LabelName]model.LabelValue{}
	anno := map[model.LabelName]model.LabelValue{}
	label["a"] = "a"
	anno["id"] = "adfasf"
	alert := &types.Alert{
		Alert:     model.Alert{Labels: label, Annotations: anno},
		UpdatedAt: time.Time{},
	}
	InAlert(alert)
}
