package cache

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewMemCache(t *testing.T) {
	tests := []struct {
		name string
		want Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Del(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.Del(tt.args.key); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Exists(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Flush(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.Flush(); got != tt.want {
				t.Errorf("Flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Get(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			got, got1 := m.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_memCache_GetValSize(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.GetValSize(tt.args.val); got != tt.want {
				t.Errorf("GetValSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Keys(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.Keys(); got != tt.want {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Set(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key    string
		val    interface{}
		expire time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.Set(tt.args.key, tt.args.val, tt.args.expire); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_SetMaxMemory(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		size string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.SetMaxMemory(tt.args.size); got != tt.want {
				t.Errorf("SetMaxMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_add(t *testing.T) {

	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key   string
		value *memCacheValue
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Add new key-value pair to an empty cache",
			fields: fields{
				maxMemorySizeStr:     "100MB",
				maxMemorySize:        100 * 1024 * 1024, // 100MB
				currMemorySize:       0,
				values:               make(map[string]*memCacheValue),
				locker:               sync.RWMutex{},
				clearExpiredItemTime: time.Duration(0),
			},
			args: args{
				key:   "testKey",
				value: &memCacheValue{val: "testValue", expire: time.Now().Add(time.Hour)},
			},
			want: true, // Expecting true since we're adding a new key-value pair to an empty cache
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			if got := m.add(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_clearExpiredItem(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			m.clearExpiredItem()
		})
	}
}

func Test_memCache_del(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			m.del(tt.args.key)
		})
	}
}

func Test_memCache_get(t *testing.T) {
	type fields struct {
		maxMemorySizeStr     string
		maxMemorySize        int64
		currMemorySize       int64
		values               map[string]*memCacheValue
		locker               sync.RWMutex
		clearExpiredItemTime time.Duration
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *memCacheValue
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				maxMemorySizeStr:     tt.fields.maxMemorySizeStr,
				maxMemorySize:        tt.fields.maxMemorySize,
				currMemorySize:       tt.fields.currMemorySize,
				values:               tt.fields.values,
				locker:               tt.fields.locker,
				clearExpiredItemTime: tt.fields.clearExpiredItemTime,
			}
			got, got1 := m.get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
