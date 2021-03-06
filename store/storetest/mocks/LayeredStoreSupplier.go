// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// LayeredStoreSupplier is an autogenerated mock type for the LayeredStoreSupplier type
type LayeredStoreSupplier struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *LayeredStoreSupplier) Next() store.LayeredStoreSupplier {
	ret := _m.Called()

	var r0 store.LayeredStoreSupplier
	if rf, ok := ret.Get(0).(func() store.LayeredStoreSupplier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LayeredStoreSupplier)
		}
	}

	return r0
}

// ReactionDelete provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreSupplier) ReactionDelete(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionDeleteAllWithEmojiName provides a mock function with given fields: ctx, emojiName, hints
func (_m *LayeredStoreSupplier) ReactionDeleteAllWithEmojiName(ctx context.Context, emojiName string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, emojiName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, emojiName, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionGetForPost provides a mock function with given fields: ctx, postId, hints
func (_m *LayeredStoreSupplier) ReactionGetForPost(ctx context.Context, postId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, postId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, postId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionPermanentDeleteBatch provides a mock function with given fields: ctx, endTime, limit, hints
func (_m *LayeredStoreSupplier) ReactionPermanentDeleteBatch(ctx context.Context, endTime int64, limit int64, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, endTime, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, endTime, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionSave provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreSupplier) ReactionSave(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGet provides a mock function with given fields: ctx, roleId, hints
func (_m *LayeredStoreSupplier) RoleGet(ctx context.Context, roleId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, roleId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, roleId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByName provides a mock function with given fields: ctx, name, hints
func (_m *LayeredStoreSupplier) RoleGetByName(ctx context.Context, name string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, name, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByNames provides a mock function with given fields: ctx, names, hints
func (_m *LayeredStoreSupplier) RoleGetByNames(ctx context.Context, names []string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, names)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, names, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RolePermanentDeleteAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) RolePermanentDeleteAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleSave provides a mock function with given fields: ctx, role, hints
func (_m *LayeredStoreSupplier) RoleSave(ctx context.Context, role *model.Role, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Role, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, role, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SetChainNext provides a mock function with given fields: _a0
func (_m *LayeredStoreSupplier) SetChainNext(_a0 store.LayeredStoreSupplier) {
	_m.Called(_a0)
}
