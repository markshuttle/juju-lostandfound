// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package converter

import (
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/names"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/watcher"
)

var logger = loggo.GetLogger("juju.apiserver.converter")

func init() {
	common.RegisterStandardFacade("Converter", 1, NewConverterAPI)
}

type ConverterAPI struct {
	st         *state.State
	resources  *common.Resources
	authorizer common.Authorizer
}

func NewConverterAPI(
	st *state.State,
	resources *common.Resources,
	authorizer common.Authorizer,
) (*ConverterAPI, error) {
	if !authorizer.AuthMachineAgent() {
		return nil, common.ErrPerm
	}
	return &ConverterAPI{
		st:         st,
		resources:  resources,
		authorizer: authorizer,
	}, nil
}

func (c *ConverterAPI) getMachine(tag names.Tag) (*state.Machine, error) {
	entity, err := c.st.FindEntity(tag)
	if err != nil {
		return nil, err
	}
	return entity.(*state.Machine), nil
}

func (c *ConverterAPI) WatchForJobsChanges(args params.Entities) (params.NotifyWatchResults, error) {
	result := params.NotifyWatchResults{
		Results: make([]params.NotifyWatchResult, len(args.Entities)),
	}
	for i, agent := range args.Entities {
		logger.Infof("Watching on entity %#v", agent)
		tag, err := names.ParseMachineTag(agent.Tag)
		if err != nil {
			return params.NotifyWatchResults{}, errors.Trace(err)
		}
		err = common.ErrPerm
		if c.authorizer.AuthOwner(tag) {
			logger.Infof("Watching for jobs on %#v", tag)
			machine, err := c.getMachine(tag)
			if err != nil {
				return params.NotifyWatchResults{}, errors.Trace(err)
			}

			watch := machine.Watch()
			// Consume the initial event. Technically, API
			// calls to Watch 'transmit' the initial event
			// in the Watch response. But NotifyWatchers
			// have no state to transmit.
			if _, ok := <-watch.Changes(); ok {
				result.Results[i].NotifyWatcherId = c.resources.Register(watch)
				err = nil
			} else {
				err = watcher.EnsureErr(watch)
			}
		}
		result.Results[i].Error = common.ServerError(err)
	}
	return result, nil
}
