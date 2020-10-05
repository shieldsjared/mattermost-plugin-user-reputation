import {combineReducers} from 'redux';

import ActionTypes from '../action_types';

function userReputations(state = {}, action) {
    switch (action.type) {
    case ActionTypes.RECEIVED_USER_REPUTATIONS: {
        const nextState = {...state};
        nextState[action.userId] = action.data;
        return nextState;
    }
    default:
        return state;
    }
}

export default combineReducers({
    userReputations,
});