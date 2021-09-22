// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"

	v1 "github.com/spotahome/redis-operator/api/redisfailover/v1"
)

// RedisFailoverCheck is an autogenerated mock type for the RedisFailoverCheck type
type RedisFailoverCheck struct {
	mock.Mock
}

// CheckAllSlavesFromMaster provides a mock function with given fields: master, rFailover
func (_m *RedisFailoverCheck) CheckAllSlavesFromMaster(master string, rFailover *v1.RedisFailover) error {
	ret := _m.Called(master, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.RedisFailover) error); ok {
		r0 = rf(master, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckRedisNumber provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) CheckRedisNumber(rFailover *v1.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckRedisSlavesReady provides a mock function with given fields: slaveIP, rFailover
func (_m *RedisFailoverCheck) CheckRedisSlavesReady(slaveIP string, rFailover *v1.RedisFailover) (bool, error) {
	ret := _m.Called(slaveIP, rFailover)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *v1.RedisFailover) bool); ok {
		r0 = rf(slaveIP, rFailover)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *v1.RedisFailover) error); ok {
		r1 = rf(slaveIP, rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckSentinelMonitor provides a mock function with given fields: sentinel, monitor
func (_m *RedisFailoverCheck) CheckSentinelMonitor(sentinel string, monitor ...string) error {
	_va := make([]interface{}, len(monitor))
	for _i := range monitor {
		_va[_i] = monitor[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, sentinel)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(sentinel, monitor...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckSentinelNumber provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) CheckSentinelNumber(rFailover *v1.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckSentinelNumberInMemory provides a mock function with given fields: sentinel, rFailover
func (_m *RedisFailoverCheck) CheckSentinelNumberInMemory(sentinel string, rFailover *v1.RedisFailover) error {
	ret := _m.Called(sentinel, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.RedisFailover) error); ok {
		r0 = rf(sentinel, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckSentinelSlavesNumberInMemory provides a mock function with given fields: sentinel, rFailover
func (_m *RedisFailoverCheck) CheckSentinelSlavesNumberInMemory(sentinel string, rFailover *v1.RedisFailover) error {
	ret := _m.Called(sentinel, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.RedisFailover) error); ok {
		r0 = rf(sentinel, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMasterIP provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetMasterIP(rFailover *v1.RedisFailover) (string, error) {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMinimumRedisPodTime provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetMinimumRedisPodTime(rFailover *v1.RedisFailover) (time.Duration, error) {
	ret := _m.Called(rFailover)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) time.Duration); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberMasters provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetNumberMasters(rFailover *v1.RedisFailover) (int, error) {
	ret := _m.Called(rFailover)

	var r0 int
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) int); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisRevisionHash provides a mock function with given fields: podName, rFailover
func (_m *RedisFailoverCheck) GetRedisRevisionHash(podName string, rFailover *v1.RedisFailover) (string, error) {
	ret := _m.Called(podName, rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, *v1.RedisFailover) string); ok {
		r0 = rf(podName, rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *v1.RedisFailover) error); ok {
		r1 = rf(podName, rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisesIPs provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetRedisesIPs(rFailover *v1.RedisFailover) ([]string, error) {
	ret := _m.Called(rFailover)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) []string); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisesMasterPod provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetRedisesMasterPod(rFailover *v1.RedisFailover) (string, error) {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedisesSlavesPods provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetRedisesSlavesPods(rFailover *v1.RedisFailover) ([]string, error) {
	ret := _m.Called(rFailover)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) []string); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelsIPs provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetSentinelsIPs(rFailover *v1.RedisFailover) ([]string, error) {
	ret := _m.Called(rFailover)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) []string); ok {
		r0 = rf(rFailover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatefulSetUpdateRevision provides a mock function with given fields: rFailover
func (_m *RedisFailoverCheck) GetStatefulSetUpdateRevision(rFailover *v1.RedisFailover) (string, error) {
	ret := _m.Called(rFailover)

	var r0 string
	if rf, ok := ret.Get(0).(func(*v1.RedisFailover) string); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.RedisFailover) error); ok {
		r1 = rf(rFailover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
