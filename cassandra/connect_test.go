package cassandra

import "testing"

func TestCassandraInitConnectionPool(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			CassandraInitConnectionPool()
		})
	}
}

func TestCassandraGetConnection(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			CassandraGetConnection()
		})
	}
}
