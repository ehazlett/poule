package mocks

import "github.com/stretchr/testify/mock"

import "github.com/google/go-github/github"

// PullRequestsService is an autogenerated mock type for the PullRequestsService type
type PullRequestsService struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: owner, repo, number, comment
func (_m *PullRequestsService) CreateComment(owner string, repo string, number int, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error) {
	ret := _m.Called(owner, repo, number, comment)

	var r0 *github.PullRequestComment
	if rf, ok := ret.Get(0).(func(string, string, int, *github.PullRequestComment) *github.PullRequestComment); ok {
		r0 = rf(owner, repo, number, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.PullRequestComment)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(string, string, int, *github.PullRequestComment) *github.Response); ok {
		r1 = rf(owner, repo, number, comment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, int, *github.PullRequestComment) error); ok {
		r2 = rf(owner, repo, number, comment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteComment provides a mock function with given fields: owner, repo, number
func (_m *PullRequestsService) DeleteComment(owner string, repo string, number int) (*github.Response, error) {
	ret := _m.Called(owner, repo, number)

	var r0 *github.Response
	if rf, ok := ret.Get(0).(func(string, string, int) *github.Response); ok {
		r0 = rf(owner, repo, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(owner, repo, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: owner, repo, opt
func (_m *PullRequestsService) List(owner string, repo string, opt *github.PullRequestListOptions) ([]github.PullRequest, *github.Response, error) {
	ret := _m.Called(owner, repo, opt)

	var r0 []github.PullRequest
	if rf, ok := ret.Get(0).(func(string, string, *github.PullRequestListOptions) []github.PullRequest); ok {
		r0 = rf(owner, repo, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]github.PullRequest)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(string, string, *github.PullRequestListOptions) *github.Response); ok {
		r1 = rf(owner, repo, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, *github.PullRequestListOptions) error); ok {
		r2 = rf(owner, repo, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListComments provides a mock function with given fields: owner, repo, number, opt
func (_m *PullRequestsService) ListComments(owner string, repo string, number int, opt *github.PullRequestListCommentsOptions) ([]github.PullRequestComment, *github.Response, error) {
	ret := _m.Called(owner, repo, number, opt)

	var r0 []github.PullRequestComment
	if rf, ok := ret.Get(0).(func(string, string, int, *github.PullRequestListCommentsOptions) []github.PullRequestComment); ok {
		r0 = rf(owner, repo, number, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]github.PullRequestComment)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(string, string, int, *github.PullRequestListCommentsOptions) *github.Response); ok {
		r1 = rf(owner, repo, number, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, int, *github.PullRequestListCommentsOptions) error); ok {
		r2 = rf(owner, repo, number, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListCommits provides a mock function with given fields: owner, repo, number, opt
func (_m *PullRequestsService) ListCommits(owner string, repo string, number int, opt *github.ListOptions) ([]github.RepositoryCommit, *github.Response, error) {
	ret := _m.Called(owner, repo, number, opt)

	var r0 []github.RepositoryCommit
	if rf, ok := ret.Get(0).(func(string, string, int, *github.ListOptions) []github.RepositoryCommit); ok {
		r0 = rf(owner, repo, number, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]github.RepositoryCommit)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(string, string, int, *github.ListOptions) *github.Response); ok {
		r1 = rf(owner, repo, number, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, int, *github.ListOptions) error); ok {
		r2 = rf(owner, repo, number, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
