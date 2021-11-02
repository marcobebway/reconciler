// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	git "github.com/go-git/go-git/v5"
	mock "github.com/stretchr/testify/mock"

	plumbing "github.com/go-git/go-git/v5/plumbing"
)

// RepoClient is an autogenerated mock type for the RepoClient type
type RepoClient struct {
	mock.Mock
}

// Clone provides a mock function with given fields: ctx, path, isBare, o
func (_m *RepoClient) Clone(ctx context.Context, path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	ret := _m.Called(ctx, path, isBare, o)

	var r0 *git.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, *git.CloneOptions) *git.Repository); ok {
		r0 = rf(ctx, path, isBare, o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, *git.CloneOptions) error); ok {
		r1 = rf(ctx, path, isBare, o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveRevision provides a mock function with given fields: rev
func (_m *RepoClient) ResolveRevision(rev plumbing.Revision) (*plumbing.Hash, error) {
	ret := _m.Called(rev)

	var r0 *plumbing.Hash
	if rf, ok := ret.Get(0).(func(plumbing.Revision) *plumbing.Hash); ok {
		r0 = rf(rev)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*plumbing.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(plumbing.Revision) error); ok {
		r1 = rf(rev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Worktree provides a mock function with given fields:
func (_m *RepoClient) Worktree() (*git.Worktree, error) {
	ret := _m.Called()

	var r0 *git.Worktree
	if rf, ok := ret.Get(0).(func() *git.Worktree); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.Worktree)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
