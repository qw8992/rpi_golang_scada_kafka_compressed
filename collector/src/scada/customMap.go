package main

import "sync"

// SyncMap 구조
type SyncMap struct {
	v   map[string]interface{}
	mux sync.RWMutex
}

// Get 함수는 key로 데이터 조회
func (sm *SyncMap) Get(key string) interface{} {
	var value interface{}
	sm.mux.RLock()
	value = sm.v[key]
	sm.mux.RUnlock()
	return value
}

// Set 함수는 key로 데이터 저장
func (sm *SyncMap) Set(key string, value interface{}) {
	sm.mux.Lock()
	sm.v[key] = value
	sm.mux.Unlock()
}

// Delete 함수는 key로 데이터 삭제
func (sm *SyncMap) Delete(key string) {
	sm.mux.Lock()
	delete(sm.v, key)
	sm.mux.Unlock()
}

// GetMap 함수는 저장되어있는 Map을 카피해서 반환
func (sm *SyncMap) GetMap() map[string]interface{} {
	sm.mux.RLock()
	value := CopyMap(sm.v)
	sm.mux.RUnlock()
	return value
}

// CopyMap 함수는 맵을 카피하는 기능을 수행
func CopyMap(originalMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

//syncMap 사이즈
func (sm *SyncMap) Size() int {
	value := len(sm.v)
	return value
}
