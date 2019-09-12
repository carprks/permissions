package service_test

import (
	"fmt"
	"github.com/carprks/permissions/service"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPermission_CreateEntry(t *testing.T) {
	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "tester",
		Permissions: []service.Permission{
			{
				Name:   "account",
				Action: "create",
			},
		},
	}
	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo create",
			request: perm,
			expect: service.Permissions{
				Identifier: "tester",
				Permissions: []service.Permission{
					{
						Name:       "account",
						Action:     "create",
						Identifier: "tester",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := perm.CreateEntry()
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("create test type err: %w", err)
			}
			passed = assert.Equal(t, test.expect, response)
			if !passed {
				t.Errorf("create test equal err: %v, %v", test.expect, response)
			}
		})
	}
}

func BenchmarkPermissions_CreateEntry(b *testing.B) {
	b.ReportAllocs()

	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "benchmark",
		Permissions: []service.Permission{
			{
				Name:   "account",
				Action: "create",
			},
		},
	}
	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo create",
			request: perm,
			expect: service.Permissions{
				Identifier: "benchmark",
				Permissions: []service.Permission{
					{
						Name:       "account",
						Action:     "create",
						Identifier: "tester",
					},
				},
			},
		},
	}

	b.ResetTimer()
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := perm.CreateEntry()
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("create test type err: %w", err)
			}

			b.StartTimer()
		})
	}
}

func TestPermission_UpdateEntry(t *testing.T) {
	if os.Getenv("DB_TABLE") != "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	orig := service.Permissions{
		Identifier: "tester",
		Permissions: []service.Permission{
			{
				Name:   "account",
				Action: "create",
			},
		},
	}
	n := service.Permissions{
		Identifier: "tester",
		Permissions: []service.Permission{
			{
				Name:       "account",
				Action:     "create",
				Identifier: "tester",
			},
			{
				Name:       "*",
				Action:     "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct {
		name    string
		request service.Permissions
		update  service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo update",
			request: orig,
			update:  n,
			expect:  n,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := test.request.UpdateEntry(test.update)
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("update test type err: %w", err)
			}
			passed = assert.Equal(t, test.expect, response)
			if !passed {
				t.Errorf("update test equal err: %v, %v", test.expect, response)
			}
		})
	}
}

func BenchmarkPermissions_UpdateEntry(b *testing.B) {
	b.ReportAllocs()

	if os.Getenv("DB_TABLE") != "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	orig := service.Permissions{
		Identifier: "benchmark",
		Permissions: []service.Permission{
			{
				Name:   "account",
				Action: "create",
			},
		},
	}
	n := service.Permissions{
		Identifier: "benchmark",
		Permissions: []service.Permission{
			{
				Name:       "account",
				Action:     "create",
				Identifier: "tester",
			},
			{
				Name:       "*",
				Action:     "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct {
		name    string
		request service.Permissions
		update  service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo update",
			request: orig,
			update:  n,
			expect:  n,
		},
	}

	b.ResetTimer()
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := test.request.UpdateEntry(test.update)
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("update test type err: %w", err)
			}

			b.StartTimer()
		})
	}
}

func TestScanEntries(t *testing.T) {
	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	tests := []struct {
		name   string
		expect int
		err    error
	}{
		{
			name:   "dynamo scan",
			expect: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := service.ScanEntries()
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("scan type err: %w", err)
			}
			passed = assert.GreaterOrEqual(t, len(response), test.expect)
			if !passed {
				t.Errorf("scan equal err: %v, %v", len(response), test.expect)
			}
		})
	}
}

func BenchmarkScanEntries(b *testing.B) {
	b.ReportAllocs()

	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	tests := []struct {
		name   string
		expect int
		err    error
	}{
		{
			name:   "dynamo scan",
			expect: 1,
		},
	}

	b.ResetTimer()
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := service.ScanEntries()
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("scan type err: %w", err)
			}

			b.StartTimer()
		})
	}
}

func TestPermission_RetrieveEntry(t *testing.T) {
	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "tester",
		Permissions: []service.Permission{
			{
				Name:       "account",
				Action:     "create",
				Identifier: "tester",
			},
			{
				Name:       "*",
				Action:     "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo retrieve",
			request: perm,
			expect:  perm,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := test.request.RetrieveEntry()
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("retrieve test type err: %w", err)
			}
			passed = assert.Equal(t, test.expect, resp)
			if !passed {
				t.Errorf("retrieve test equal err: %v, %v", test.expect, resp)
			}
		})
	}
}

func BenchmarkPermissions_RetrieveEntry(b *testing.B) {
	b.ReportAllocs()

	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "benchmark",
		Permissions: []service.Permission{
			{
				Name:       "account",
				Action:     "create",
				Identifier: "tester",
			},
			{
				Name:       "*",
				Action:     "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo retrieve",
			request: perm,
			expect:  perm,
			err:     nil,
		},
	}

	b.ResetTimer()
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := test.request.RetrieveEntry()
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("retrieve test type err: %w", err)
			}

			b.StartTimer()
		})
	}
}

func TestPermission_DeleteEntry(t *testing.T) {
	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "tester",
	}

	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo delete",
			request: perm,
			expect: service.Permissions{
				Identifier: "tester",
				Status:     "deleted",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := test.request.DeleteEntry()
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("delete test type err: %w", err)
			}
			passed = assert.Equal(t, test.expect, response)
			if !passed {
				t.Errorf("delete test equal err: %v, %v", test.expect, response)
			}
		})
	}
}

func BenchmarkPermissions_DeleteEntry(b *testing.B) {
	b.ReportAllocs()

	if os.Getenv("DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := service.Permissions{
		Identifier: "tester",
	}

	tests := []struct {
		name    string
		request service.Permissions
		expect  service.Permissions
		err     error
	}{
		{
			name:    "dynamo delete",
			request: perm,
			expect: service.Permissions{
				Identifier: "benchmark",
				Status:     "deleted",
			},
		},
	}

	b.ResetTimer()
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := test.request.DeleteEntry()
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("delete test type err: %w", err)
			}

			b.StartTimer()
		})
	}
}
