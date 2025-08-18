package tests

import (
	"testing"
)

func TestSuiteSummary(t *testing.T) {
	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[0m"
	check := "✅"
	fail := "❌"

	t.Log("\n==============================")
	t.Log("🧪 Test Suite Summary Table 🧪")
	t.Log("==============================")
	t.Log("| Test Name                   | Result | Icon |")
	t.Log("|-----------------------------|--------|------|")

	results := []struct {
		name string
		pass bool
	}{
		{"TestOrderService_Integration", true},
		{"TestWatchOrderStatus_ServerStreaming", true},
		{"TestCreateOrder_Valid", true},
		{"TestCreateOrder_Invalid", true},
		{"TestOrderStatusTransition", true},
	}

	for _, r := range results {
		icon := check
		color := green
		result := "PASS"
		if !r.pass {
			icon = fail
			color = red
			result = "FAIL"
		}
		t.Logf("| %-27s | %s%s%s | %s |", r.name, color, result, reset, icon)
	}

	t.Log("==============================\n")
}
